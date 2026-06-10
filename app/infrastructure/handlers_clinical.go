package app

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"omhmre.com/centromedico/app/domain/models"
)

// GetClinicalHistoryHandler maneja la petición GET para el historial clínico general
func (a *App) GetClinicalHistoryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cedula := r.URL.Query().Get("cedula")
		if cedula == "" {
			http.Error(w, "Cédula es requerida", http.StatusBadRequest)
			return
		}

		history, err := a.DB.GetClinicalHistory(cedula)
		if err != nil {
			if err == sql.ErrNoRows {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(&models.ClinicalHistory{CedulaPaciente: cedula})
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(history)
	}
}

// UpsertClinicalHistoryHandler maneja la actualización del historial base
func (a *App) UpsertClinicalHistoryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var history models.ClinicalHistory
		if err := json.NewDecoder(r.Body).Decode(&history); err != nil {
			http.Error(w, "Body inválido", http.StatusBadRequest)
			return
		}

		id, err := a.DB.UpsertClinicalHistory(&history)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		history.ID = id

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(history)
	}
}

// GetClinicalRecordsHandler maneja la petición GET para los registros médicos (notas)
func (a *App) GetClinicalRecordsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cedula := r.URL.Query().Get("cedula")
		if cedula == "" {
			http.Error(w, "Cédula es requerida", http.StatusBadRequest)
			return
		}

		records, err := a.DB.GetClinicalRecords(cedula)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if records == nil {
			records = []models.ClinicalRecord{}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(records)
	}
}

// PostClinicalRecordHandler maneja la inserción de una nueva nota clínica
func (a *App) PostClinicalRecordHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var record models.ClinicalRecord
		if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
			http.Error(w, "Body inválido", http.StatusBadRequest)
			return
		}

		id, err := a.DB.InsertClinicalRecord(&record)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		record.ID = id

		for i := range record.Attachments {
			record.Attachments[i].IDRegistro = id
			attId, _ := a.DB.InsertClinicalAttachment(&record.Attachments[i])
			record.Attachments[i].ID = attId
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(record)
	}
}
