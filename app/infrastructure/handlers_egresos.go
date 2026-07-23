package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"omhmre.com/centromedico/app/domain/models"
)

func (a *App) GetEgresos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var f models.Fechas
		if r.Body != nil {
			_ = json.NewDecoder(r.Body).Decode(&f)
		}

		res, st := a.DB.GetEgresos(f)
		if st.Status >= 400 {
			sendResponse(w, r, st, http.StatusInternalServerError)
			return
		}
		// Enviar la lista de egresos
		sendResponse(w, r, res, http.StatusOK)
	}
}

func (a *App) PostEgreso() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Egreso
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			sendResponse(w, r, models.Respuesta{
				Status:  http.StatusBadRequest,
				Mensaje: "Datos de egreso inválidos: " + err.Error(),
			}, http.StatusBadRequest)
			return
		}

		st := a.DB.PostEgreso(e)
		if st.Status >= 400 {
			sendResponse(w, r, st, http.StatusInternalServerError)
		} else {
			a.broadcastEvent("EGRESOS_UPDATED", nil)
			sendResponse(w, r, st, http.StatusOK)
		}
	}
}

func (a *App) PutEgreso() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e models.Egreso
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			sendResponse(w, r, models.Respuesta{
				Status:  http.StatusBadRequest,
				Mensaje: "Datos de egreso inválidos: " + err.Error(),
			}, http.StatusBadRequest)
			return
		}

		st := a.DB.PutEgreso(e)
		if st.Status >= 400 {
			sendResponse(w, r, st, http.StatusInternalServerError)
		} else {
			a.broadcastEvent("EGRESOS_UPDATED", nil)
			sendResponse(w, r, st, http.StatusOK)
		}
	}
}

func (a *App) DelEgreso() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			ID interface{} `json:"id"`
		}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil || payload.ID == nil {
			sendResponse(w, r, models.Respuesta{
				Status:  http.StatusBadRequest,
				Mensaje: "ID de egreso inválido",
			}, http.StatusBadRequest)
			return
		}

		idStr := fmt.Sprintf("%v", payload.ID)
		st := a.DB.DelEgreso(models.Id{Id: idStr})
		if st.Status >= 400 {
			sendResponse(w, r, st, http.StatusInternalServerError)
		} else {
			a.broadcastEvent("EGRESOS_UPDATED", nil)
			sendResponse(w, r, st, http.StatusOK)
		}
	}
}

// ─── Handlers de Configuración de Egresos ─────────────────────────────────────

func (a *App) GetConfigEgresos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, st := a.DB.GetConfigEgresos()
		if st.Status >= 400 {
			sendResponse(w, r, st, http.StatusInternalServerError)
			return
		}
		sendResponse(w, r, res, http.StatusOK)
	}
}

func (a *App) PostConfigEgreso() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c models.ConfigEgreso
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			sendResponse(w, r, models.Respuesta{
				Status:  http.StatusBadRequest,
				Mensaje: "Datos de configuración inválidos: " + err.Error(),
			}, http.StatusBadRequest)
			return
		}

		st := a.DB.PostConfigEgreso(c)
		if st.Status >= 400 {
			sendResponse(w, r, st, http.StatusInternalServerError)
		} else {
			a.broadcastEvent("EGRESOS_CONFIG_UPDATED", nil)
			sendResponse(w, r, st, http.StatusOK)
		}
	}
}

func (a *App) DelConfigEgreso() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			ID interface{} `json:"id"`
		}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil || payload.ID == nil {
			sendResponse(w, r, models.Respuesta{
				Status:  http.StatusBadRequest,
				Mensaje: "ID de opción inválido",
			}, http.StatusBadRequest)
			return
		}

		idStr := fmt.Sprintf("%v", payload.ID)
		st := a.DB.DelConfigEgreso(models.Id{Id: idStr})
		if st.Status >= 400 {
			sendResponse(w, r, st, http.StatusInternalServerError)
		} else {
			a.broadcastEvent("EGRESOS_CONFIG_UPDATED", nil)
			sendResponse(w, r, st, http.StatusOK)
		}
	}
}

