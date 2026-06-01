package services

import (
	"context"
	"fmt"
	"time"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	"omhmre.com/centromedico/app/domain/database"
	"omhmre.com/centromedico/app/domain/models"
	"omhmre.com/centromedico/app/domain/utils"
)

// ProcessAIChat runs the Gemini model with tool calling logic and returns a final response and optional widget
func ProcessAIChat(ctx context.Context, apiKey string, db database.PostDB, req models.AIChatRequest) (models.AIChatResponse, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return models.AIChatResponse{}, fmt.Errorf("failed to create GenAI client: %w", err)
	}
	defer client.Close()

	// Use gemini-1.5-flash for faster responsiveness suitable for chats
	model := client.GenerativeModel("gemini-1.5-flash")

	// Set temperature and system instructions
	model.SetTemperature(0.2)
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{
			genai.Text("Eres Jarvis, el Asistente Virtual Inteligente del Centro Médico. Tu objetivo es ayudar al usuario con el agendamiento de citas médicas.\n\n" +
				"Sigue estas reglas estrictas:\n" +
				"1. Escribe de forma concisa, profesional y amable en Español.\n" +
				"2. Para buscar médicos, utiliza siempre la herramienta 'buscarMedicos' indicando la especialidad requerida.\n" +
				"3. Para buscar disponibilidad, utiliza siempre 'buscarDisponibilidadCitas' con el ID del médico y la fecha (YYYY-MM-DD).\n" +
				"4. Cuando el paciente decida agendar, pídele su cédula de identidad y el motivo de la consulta. Cuando tengas médico, fecha, hora, cédula y motivo, invoca la herramienta 'preConfirmarCita'.\n" +
				"5. Explica al usuario que la cita queda pre-agendada y que debe presionar el botón de 'Confirmar' en la pantalla para guardarla definitivamente.\n" +
				"6. Si el usuario te da fechas relativas como 'mañana' o 'el próximo martes', calcula la fecha correspondiente basándote en la fecha actual que es: " + time.Now().Format("2006-01-02") + " (Día de la semana: " + time.Now().Weekday().String() + ").\n"),
		},
	}

	// Define tools
	model.Tools = []*genai.Tool{
		{
			FunctionDeclarations: []*genai.FunctionDeclaration{
				{
					Name:        "buscarMedicos",
					Description: "Busca médicos o especialistas activos según una especialidad (por ejemplo, 'pediatria', 'cardiologia', 'ginecologia'). Retorna sus nombres, ID, especialidad, costo de cita y contacto.",
					Parameters: &genai.Schema{
						Type: genai.TypeObject,
						Properties: map[string]*genai.Schema{
							"especialidad": {
								Type:        genai.TypeString,
								Description: "La especialidad médica a buscar.",
							},
						},
						Required: []string{"especialidad"},
					},
				},
				{
					Name:        "buscarDisponibilidadCitas",
					Description: "Busca horarios de citas disponibles para un doctor específico en una fecha. Retorna las horas libres.",
					Parameters: &genai.Schema{
						Type: genai.TypeObject,
						Properties: map[string]*genai.Schema{
							"medicoId": {
								Type:        genai.TypeInteger,
								Description: "El ID único del médico del cual buscar disponibilidad.",
							},
							"fecha": {
								Type:        genai.TypeString,
								Description: "La fecha a consultar en formato YYYY-MM-DD (ej: '2026-06-02').",
							},
						},
						Required: []string{"medicoId", "fecha"},
					},
				},
				{
					Name:        "preConfirmarCita",
					Description: "Genera una propuesta o pre-confirmación de cita. NO guarda nada en la base de datos, solo prepara los datos para que el usuario confirme físicamente.",
					Parameters: &genai.Schema{
						Type: genai.TypeObject,
						Properties: map[string]*genai.Schema{
							"medicoId": {
								Type:        genai.TypeInteger,
								Description: "El ID del médico especialista.",
							},
							"fecha": {
								Type:        genai.TypeString,
								Description: "La fecha deseada de la cita en formato YYYY-MM-DD.",
							},
							"hora": {
								Type:        genai.TypeString,
								Description: "La hora deseada en formato HH:MM (ej: '09:00', '14:30').",
							},
							"cedula": {
								Type:        genai.TypeString,
								Description: "La cédula del paciente.",
							},
							"motivo": {
								Type:        genai.TypeString,
								Description: "El motivo de la consulta médica.",
							},
						},
						Required: []string{"medicoId", "fecha", "hora", "cedula", "motivo"},
					},
				},
			},
		},
	}

	// Initialize chat session
	cs := model.StartChat()

	// Reconstruct conversation history
	var genaiHistory []*genai.Content
	for _, msg := range req.History {
		if msg.Content == "" {
			continue
		}
		role := "user"
		if msg.Role == "model" || msg.Role == "assistant" {
			role = "model"
		}
		
		// The history must start with a user turn
		if len(genaiHistory) == 0 && role == "model" {
			continue
		}

		// Ensure alternating roles
		if len(genaiHistory) > 0 && genaiHistory[len(genaiHistory)-1].Role == role {
			continue
		}

		genaiHistory = append(genaiHistory, &genai.Content{
			Role: role,
			Parts: []genai.Part{
				genai.Text(msg.Content),
			},
		})
	}
	cs.History = genaiHistory

	// Send user's new message
	resp, err := cs.SendMessage(ctx, genai.Text(req.Message))
	if err != nil {
		return models.AIChatResponse{}, fmt.Errorf("failed to send message: %w", err)
	}

	var lastPreConfirmation *models.PreConfirmationCard
	var lastDoctorCards []models.DoctorCard

	// Loop to handle tool requests (max 5 iterations)
	for iter := 0; iter < 5; iter++ {
		if resp == nil || len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil {
			break
		}

		hasFunctionCall := false
		var funcResponses []genai.Part

		for _, part := range resp.Candidates[0].Content.Parts {
			if fnCall, ok := part.(genai.FunctionCall); ok {
				hasFunctionCall = true
				utils.CreateLog(fmt.Sprintf("[AI Agent] Executing tool call: %s with args: %+v", fnCall.Name, fnCall.Args))

				resultMap, toolErr := executeToolCall(ctx, db, fnCall, &lastPreConfirmation, &lastDoctorCards)
				var respMap map[string]interface{}
				if toolErr != nil {
					respMap = map[string]interface{}{"error": toolErr.Error()}
				} else {
					respMap = resultMap
				}

				funcResponses = append(funcResponses, genai.FunctionResponse{
					Name:     fnCall.Name,
					Response: respMap,
				})
			}
		}

		if !hasFunctionCall {
			break
		}

		// Send tool responses back to the model
		resp, err = cs.SendMessage(ctx, funcResponses...)
		if err != nil {
			return models.AIChatResponse{}, fmt.Errorf("failed to send function responses: %w", err)
		}
	}

	// Prepare final response
	chatResp := models.AIChatResponse{}
	if resp != nil && len(resp.Candidates) > 0 && resp.Candidates[0].Content != nil {
		// Concatenate all text parts
		var finalMessage string
		for _, part := range resp.Candidates[0].Content.Parts {
			if textPart, ok := part.(genai.Text); ok {
				finalMessage += string(textPart)
			}
		}
		chatResp.Response = finalMessage
	}

	// Attach widgets/payloads if generated during the tool loop
	if lastPreConfirmation != nil {
		chatResp.Widget = &models.AIWidget{
			Type: "pre_confirmation",
			Data: lastPreConfirmation,
		}
	} else if len(lastDoctorCards) > 0 {
		chatResp.Widget = &models.AIWidget{
			Type: "doctor_list",
			Data: lastDoctorCards,
		}
	}

	return chatResp, nil
}

func executeToolCall(ctx context.Context, db database.PostDB, fnCall genai.FunctionCall, preConf **models.PreConfirmationCard, doctorCards *[]models.DoctorCard) (map[string]interface{}, error) {
	switch fnCall.Name {
	case "buscarMedicos":
		especialidad, ok := fnCall.Args["especialidad"].(string)
		if !ok {
			return nil, fmt.Errorf("argument 'especialidad' is missing or not a string")
		}

		doctores, dbResp := db.GetDoctoresPorEspecialidad(especialidad)
		if dbResp.Status >= 400 {
			return nil, fmt.Errorf("database query error: %s", dbResp.Mensaje)
		}

		var cards []models.DoctorCard
		var responseList []map[string]interface{}
		for _, doc := range doctores {
			card := models.DoctorCard{
				ID:           doc.Id,
				Nombres:      doc.Nombres,
				Especialidad: doc.Espec,
				MontoCita:    doc.MontoCita,
				Dir:          doc.Dir,
				Whatsapp:     doc.Whatsapp,
			}
			cards = append(cards, card)
			responseList = append(responseList, map[string]interface{}{
				"id":        doc.Id,
				"nombres":   doc.Nombres,
				"espec":     doc.Espec,
				"monto":     doc.MontoCita,
				"dias_trab": doc.DaysOfWeek,
				"inicio":    doc.StartTime,
				"fin":       doc.EndTime,
				"slot_min":  doc.SlotDuration,
			})
		}
		*doctorCards = cards
		return map[string]interface{}{"medicos": responseList}, nil

	case "buscarDisponibilidadCitas":
		medicoIDFloat, ok := fnCall.Args["medicoId"].(float64)
		if !ok {
			return nil, fmt.Errorf("argument 'medicoId' is missing or not a number")
		}
		medicoID := int(medicoIDFloat)

		fecha, ok := fnCall.Args["fecha"].(string)
		if !ok {
			return nil, fmt.Errorf("argument 'fecha' is missing or not a string")
		}

		doc, dbResp := db.GetDoctor(medicoID)
		if dbResp.Status >= 400 {
			return nil, fmt.Errorf("doctor with ID %d not found: %s", medicoID, dbResp.Mensaje)
		}

		citas, dbResp := db.GetCitasDoctorFecha(medicoID, fecha)
		if dbResp.Status >= 400 {
			return nil, fmt.Errorf("error reading appointments for date %s: %s", fecha, dbResp.Mensaje)
		}

		slotDuration := doc.SlotDuration
		if slotDuration <= 0 {
			slotDuration = 30
		}

		startTimeStr := doc.StartTime
		if startTimeStr == "" {
			startTimeStr = "08:00"
		}
		endTimeStr := doc.EndTime
		if endTimeStr == "" {
			endTimeStr = "17:00"
		}

		tLayout := "15:04"
		if len(startTimeStr) > 5 {
			startTimeStr = startTimeStr[:5]
		}
		if len(endTimeStr) > 5 {
			endTimeStr = endTimeStr[:5]
		}

		startT, err := time.Parse(tLayout, startTimeStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse doctor start time %s: %w", startTimeStr, err)
		}
		endT, err := time.Parse(tLayout, endTimeStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse doctor end time %s: %w", endTimeStr, err)
		}

		parsedDate, err := time.Parse("2006-01-02", fecha)
		if err != nil {
			return nil, fmt.Errorf("invalid date format %s, use YYYY-MM-DD: %w", fecha, err)
		}

		// Go Weekday: Sunday = 0, Monday = 1, ..., Saturday = 6
		weekday := int(parsedDate.Weekday())

		isWorkingDay := false
		for _, d := range doc.DaysOfWeek {
			if d == weekday || (weekday == 0 && d == 7) {
				isWorkingDay = true
				break
			}
		}

		var availableSlots []string
		if isWorkingDay {
			bookedMap := make(map[string]bool)
			for _, cita := range citas {
				bookedMap[cita.Inicio.Format("15:04")] = true
			}

			currentT := startT
			for currentT.Before(endT) {
				slotStr := currentT.Format("15:04")
				if !bookedMap[slotStr] {
					availableSlots = append(availableSlots, slotStr)
				}
				currentT = currentT.Add(time.Duration(slotDuration) * time.Minute)
			}
		}

		return map[string]interface{}{
			"fecha":       fecha,
			"medico":      doc.Nombres,
			"disponibles": availableSlots,
		}, nil

	case "preConfirmarCita":
		medicoIDFloat, ok := fnCall.Args["medicoId"].(float64)
		if !ok {
			return nil, fmt.Errorf("argument 'medicoId' is missing or not a number")
		}
		medicoID := int(medicoIDFloat)

		fecha, ok := fnCall.Args["fecha"].(string)
		if !ok {
			return nil, fmt.Errorf("argument 'fecha' is missing or not a string")
		}
		hora, ok := fnCall.Args["hora"].(string)
		if !ok {
			return nil, fmt.Errorf("argument 'hora' is missing or not a string")
		}
		cedula, ok := fnCall.Args["cedula"].(string)
		if !ok {
			return nil, fmt.Errorf("argument 'cedula' is missing or not a string")
		}
		motivo, ok := fnCall.Args["motivo"].(string)
		if !ok {
			return nil, fmt.Errorf("argument 'motivo' is missing or not a string")
		}

		doc, dbResp := db.GetDoctor(medicoID)
		if dbResp.Status >= 400 {
			return nil, fmt.Errorf("doctor with ID %d not found: %s", medicoID, dbResp.Mensaje)
		}

		paciente, dbResp := db.GetPacientePorCedula(cedula)
		pacienteName := "Paciente Nuevo / Externo"
		if dbResp.Status == 200 {
			pacienteName = paciente.Nombres
		}

		preCard := &models.PreConfirmationCard{
			DoctorID:     doc.Id,
			DoctorName:   doc.Nombres,
			Especialidad: doc.Espec,
			Cedula:       cedula,
			PacienteName: pacienteName,
			Fecha:        fecha,
			Hora:         hora,
			Motivo:       motivo,
			MontoCita:    doc.MontoCita,
		}
		*preConf = preCard

		return map[string]interface{}{
			"status":           "proposed",
			"pre_confirmacion": preCard,
		}, nil
	}

	return nil, fmt.Errorf("unknown tool call function: %s", fnCall.Name)
}
