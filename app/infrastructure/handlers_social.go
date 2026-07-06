package infrastructure

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"omhmre.com/centromedico/app/domain/models"
)

// GetSocialEvaluationHandler maneja la petición GET para la evaluación social
func (a *App) GetSocialEvaluationHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cedula := r.URL.Query().Get("cedula")
		if cedula == "" {
			http.Error(w, "Cédula es requerida", http.StatusBadRequest)
			return
		}

		evaluation, err := a.DB.GetSocialEvaluation(cedula)
		if err != nil {
			// Si el error contiene sql.ErrNoRows, devolvemos un objeto vacío con la cédula
			if errors.Is(err, sql.ErrNoRows) || (err != nil && errors.Is(errors.Unwrap(err), sql.ErrNoRows)) {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(&models.SocialEvaluation{CedulaPaciente: cedula})
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(evaluation)
	}
}

// UpsertSocialEvaluationHandler maneja la creación o actualización de la evaluación social
func (a *App) UpsertSocialEvaluationHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var eval models.SocialEvaluation
		if err := json.NewDecoder(r.Body).Decode(&eval); err != nil {
			http.Error(w, "Body inválido", http.StatusBadRequest)
			return
		}

		id, err := a.DB.UpsertSocialEvaluation(&eval)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		eval.ID = id

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(eval)
	}
}

// UnlockSocialEvaluationHandler maneja el desbloqueo de la evaluación social
func (a *App) UnlockSocialEvaluationHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Cedula string `json:"cedula"`
			UserId int    `json:"user_id"`
			Reason string `json:"reason"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Body inválido", http.StatusBadRequest)
			return
		}

		if req.Cedula == "" || req.UserId == 0 || req.Reason == "" {
			http.Error(w, "Faltan parámetros", http.StatusBadRequest)
			return
		}

		id, err := a.DB.UnlockSocialEvaluation(req.Cedula, req.UserId, req.Reason)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"id":      id,
			"message": "Evaluación desbloqueada correctamente",
		})
	}
}
