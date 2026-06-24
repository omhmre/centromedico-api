package infrastructure

import (
	"encoding/json"
	"net/http"
	"strconv"

	"omhmre.com/centromedico/app/domain/models"
	"omhmre.com/centromedico/app/domain/utils"
)

// GetDoctoresPorUsuario retorna los doctores asignados a un usuario.
// GET /getdoctoresporusuario?id_usuario=X
func (a *App) GetDoctoresPorUsuario() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		userIDStr := r.URL.Query().Get("id_usuario")
		if userIDStr == "" {
			var rp models.Respuesta
			rp.Status = http.StatusBadRequest
			rp.Mensaje = "El parámetro id_usuario es requerido"
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(rp)
			return
		}

		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			var rp models.Respuesta
			rp.Status = http.StatusBadRequest
			rp.Mensaje = "El parámetro id_usuario debe ser un número entero"
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(rp)
			return
		}

		datos, dbResp := a.DB.GetDoctoresPorUsuario(userID)
		if dbResp.Status >= 400 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(dbResp)
			return
		}

		json.NewEncoder(w).Encode(datos)
	}
}

// GuardarAsignacionRequest define el cuerpo para guardar las asignaciones.
type GuardarAsignacionRequest struct {
	IdUsuario int64 `json:"id_usuario"`
	Doctores  []int `json:"doctores"`
}

// SaveDoctoresUsuario guarda los doctores asignados a un usuario.
// POST /savedoctoresusuario
func (a *App) SaveDoctoresUsuario() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		var req GuardarAsignacionRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			var rp models.Respuesta
			rp.Status = http.StatusBadRequest
			rp.Mensaje = "Cuerpo de solicitud inválido: " + err.Error()
			utils.CreateLog(rp.Mensaje)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(rp)
			return
		}

		if req.IdUsuario == 0 {
			var rp models.Respuesta
			rp.Status = http.StatusBadRequest
			rp.Mensaje = "El campo id_usuario es requerido y no puede ser 0"
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(rp)
			return
		}

		rp := a.DB.SaveDoctoresUsuario(req.IdUsuario, req.Doctores)
		if rp.Status >= 400 {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		json.NewEncoder(w).Encode(rp)
	}
}
