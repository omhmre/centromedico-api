package infrastructure

import (
	"encoding/json"
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
		var i models.Id
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "Cuerpo de solicitud inválido"}, http.StatusBadRequest)
			return
		}
		rp := a.DB.DelPersonal(i)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}
