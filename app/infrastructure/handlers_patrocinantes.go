package infrastructure

import (
	"encoding/json"
	"net/http"
	"strconv"

	"omhmre.com/centromedico/app/domain/models"
)

// GetPatrocinantes maneja las peticiones GET para consultar patrocinantes.
func (a *App) GetPatrocinantes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		datos, resp := a.DB.GetPatrocinantes()
		if resp.Status >= 400 {
			sendResponse(w, r, resp, resp.Status)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(datos)
	}
}

// PostPatrocinante maneja las peticiones POST para registrar un patrocinante.
func (a *App) PostPatrocinante() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p models.Patrocinante
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "Cuerpo de solicitud inválido: " + err.Error()}, 400)
			return
		}

		if p.Nombre == "" {
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "Nombre del patrocinante es requerido"}, 400)
			return
		}

		resp := a.DB.PostPatrocinante(p)
		sendResponse(w, r, resp, resp.Status)
	}
}

// DeletePatrocinante maneja las peticiones DELETE para eliminar un patrocinante por ID.
func (a *App) DeletePatrocinante() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "ID de patrocinante inválido"}, 400)
			return
		}

		resp := a.DB.DeletePatrocinante(id)
		sendResponse(w, r, resp, resp.Status)
	}
}
