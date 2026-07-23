package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"omhmre.com/centromedico/app/domain/models"
)

func (a *App) GetPersonal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		datos, rp := a.DB.GetPersonal()
		if rp.Status != 200 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(datos)
	}
}

func (a *App) PostPersonal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p models.PersonalModel
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "Cuerpo de solicitud inválido"}, http.StatusBadRequest)
			return
		}
		rp := a.DB.PostPersonal(p)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) UpdPersonal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p models.PersonalModel
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "Cuerpo de solicitud inválido"}, http.StatusBadRequest)
			return
		}
		rp := a.DB.UpdPersonal(p)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) DelPersonal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			ID interface{} `json:"id"`
		}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil || payload.ID == nil {
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "Cuerpo de solicitud o ID inválido"}, http.StatusBadRequest)
			return
		}
		idStr := fmt.Sprintf("%v", payload.ID)
		rp := a.DB.DelPersonal(models.Id{Id: idStr})
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

