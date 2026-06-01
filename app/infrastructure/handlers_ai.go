package infrastructure

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"omhmre.com/centromedico/app/domain/database"
	"omhmre.com/centromedico/app/domain/models"
	"omhmre.com/centromedico/app/domain/utils"
	"omhmre.com/centromedico/app/infrastructure/services"
)

// AIChat handles incoming chat prompts and delegates to Gemini service
func (a *App) AIChat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := database.GEMINI_API_KEY
		if apiKey == "" {
			apiKey = a.DB.GetParametroValor("GEMINI_API_KEY")
		}

		if apiKey == "" {
			respuesta := models.Respuesta{
				Status:  http.StatusInternalServerError,
				Mensaje: "La API Key de Gemini no está configurada en el servidor backend (revisa las variables de entorno o la configuración de parámetros).",
			}
			sendResponse(w, r, respuesta, http.StatusInternalServerError)
			return
		}

		var chatReq models.AIChatRequest
		err := json.NewDecoder(r.Body).Decode(&chatReq)
		if err != nil {
			respuesta := models.Respuesta{
				Status:  http.StatusBadRequest,
				Mensaje: "Cuerpo de solicitud inválido: " + err.Error(),
			}
			sendResponse(w, r, respuesta, http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 45*time.Second)
		defer cancel()

		response, err := services.ProcessAIChat(ctx, apiKey, a.DB, chatReq)
		if err != nil {
			utils.CreateLog("Error procesando chat de IA: " + err.Error())
			respuesta := models.Respuesta{
				Status:  http.StatusInternalServerError,
				Mensaje: "Error procesando tu solicitud con Jarvis: " + err.Error(),
			}
			sendResponse(w, r, respuesta, http.StatusInternalServerError)
			return
		}

		sendResponse(w, r, response, http.StatusOK)
	}
}
