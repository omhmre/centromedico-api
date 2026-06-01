package models

// AIChatMessage represents a single message in the chat history
type AIChatMessage struct {
	Role    string `json:"role"`    // "user" or "model"
	Content string `json:"content"`
}

// AIChatRequest represents the incoming chat request from Flutter
type AIChatRequest struct {
	Message string          `json:"message"`
	History []AIChatMessage `json:"history"`
	Cedula  string          `json:"cedula,omitempty"` // Identificación del paciente activo en el cliente
}

// AIChatResponse represents the response containing the assistant's message and optional interactive widgets/payloads
type AIChatResponse struct {
	Response string      `json:"response"`
	Widget   *AIWidget   `json:"widget,omitempty"` // Tarjeta adaptativa o acción interactiva para Flutter
}

// AIWidget defines standard structured actions that the frontend can render interactively
type AIWidget struct {
	Type string      `json:"type"` // "pre_confirmation", "doctor_list", "slot_list"
	Data interface{} `json:"data"`
}

// DoctorCard represents doctor info returned in cards
type DoctorCard struct {
	ID           int     `json:"id"`
	Nombres      string  `json:"nombres"`
	Especialidad string  `json:"especialidad"`
	MontoCita    float64 `json:"monto_cita"`
	Dir          string  `json:"dir"`
	Whatsapp     string  `json:"whatsapp"`
}

// SlotCard represents a proposed time slot for booking
type SlotCard struct {
	Hora string `json:"hora"` // HH:MM
}

// PreConfirmationCard represents the payload for human confirmation before writing to DB
type PreConfirmationCard struct {
	DoctorID     int     `json:"doctor_id"`
	DoctorName   string  `json:"doctor_name"`
	Especialidad string  `json:"especialidad"`
	Cedula       string  `json:"cedula"`
	PacienteName string  `json:"paciente_name"`
	Fecha        string  `json:"fecha"` // YYYY-MM-DD
	Hora         string  `json:"hora"`  // HH:MM
	Motivo       string  `json:"motivo"`
	MontoCita    float64 `json:"monto_cita"`
}
