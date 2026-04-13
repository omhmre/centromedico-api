package infrastructure

import (
	"encoding/json"
	"net/http"

	"omhmre.com/centromedico/app/domain/models"
)

func (a *App) GetEgresos(w http.ResponseWriter, r *http.Request) {
	var f models.Fechas
	err := json.NewDecoder(r.Body).Decode(&f)
	if err != nil {
		Responder(w, http.StatusBadRequest, models.Respuesta{Status: 400, Mensaje: "datos de fecha invalidos"})
		return
	}

	res, status := a.DB.GetEgresos(f)
	if status.Status == 500 {
		Responder(w, http.StatusInternalServerError, status)
		return
	}
	Responder(w, http.StatusOK, res)
}

func (a *App) PostEgreso(w http.ResponseWriter, r *http.Request) {
	var e models.Egreso
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		Responder(w, http.StatusBadRequest, models.Respuesta{Status: 400, Mensaje: "datos de egreso invalidos"})
		return
	}

	status := a.DB.PostEgreso(e)
	if status.Status == 500 {
		Responder(w, http.StatusInternalServerError, status)
		return
	}
	Responder(w, http.StatusCreated, status)
}

func (a *App) PutEgreso(w http.ResponseWriter, r *http.Request) {
	var e models.Egreso
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		Responder(w, http.StatusBadRequest, models.Respuesta{Status: 400, Mensaje: "datos de egreso invalidos"})
		return
	}

	status := a.DB.PutEgreso(e)
	if status.Status == 500 {
		Responder(w, http.StatusInternalServerError, status)
		return
	}
	Responder(w, http.StatusOK, status)
}

func (a *App) DelEgreso(w http.ResponseWriter, r *http.Request) {
	var i models.Id
	err := json.NewDecoder(r.Body).Decode(&i)
	if err != nil {
		Responder(w, http.StatusBadRequest, models.Respuesta{Status: 400, Mensaje: "ID invalido"})
		return
	}

	status := a.DB.DelEgreso(i)
	if status.Status == 500 {
		Responder(w, http.StatusInternalServerError, status)
		return
	}
	Responder(w, http.StatusOK, status)
}
