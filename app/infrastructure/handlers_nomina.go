package infrastructure

import (
	"encoding/json"
	"net/http"

	"omhmre.com/centromedico/app/domain/models"
	"omhmre.com/centromedico/app/domain/utils"
)

func (a *App) GetNominas() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		desde := r.URL.Query().Get("desde")
		hasta := r.URL.Query().Get("hasta")

		if desde == "" || hasta == "" {
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "Se requieren fechas desde y hasta"}, http.StatusBadRequest)
			return
		}

		datos, rp := a.DB.GetNominas(desde, hasta)
		if rp.Status != 200 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
			return
		}
		sendResponse(w, r, datos, http.StatusOK)
	}
}

func (a *App) PostNomina() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var n models.NominaModel
		if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "Cuerpo de solicitud inválido"}, http.StatusBadRequest)
			return
		}
		rp := a.DB.PostNomina(n)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) PutNomina() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var n models.NominaModel
		if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
			utils.CreateLog("Error decodificando PutNomina: " + err.Error())
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "Cuerpo de solicitud inválido"}, http.StatusBadRequest)
			return
		}
		rp := a.DB.UpdNomina(n)
		if rp.Status >= 400 {
			utils.CreateLog("Error en DB.UpdNomina: " + rp.Mensaje)
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) PayNomina() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			NominaID         int    `json:"nomina_id"`
			FechaPago        string `json:"fecha_pago"`
			MetodoPago       string `json:"metodo_pago"`
			UsuarioOperacion string `json:"usuario_operacion"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "Cuerpo de solicitud inválido"}, http.StatusBadRequest)
			return
		}
		rp := a.DB.PayNomina(req.NominaID, req.FechaPago, req.MetodoPago, req.UsuarioOperacion)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) DelNomina() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i models.Id
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "Cuerpo de solicitud inválido"}, http.StatusBadRequest)
			return
		}
		rp := a.DB.DelNomina(i)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}
