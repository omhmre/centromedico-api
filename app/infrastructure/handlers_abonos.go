package infrastructure

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"omhmre.com/centromedico/app/domain/models"
)

// GetAbonos maneja las peticiones GET para consultar abonos de un paciente o patrocinante.
func (a *App) GetAbonos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cedula := r.URL.Query().Get("cedula")
		patrocinante := r.URL.Query().Get("patrocinante")

		datos, resp := a.DB.GetAbonos(cedula, patrocinante)
		if resp.Status >= 400 {
			sendResponse(w, r, resp, resp.Status)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(datos)
	}
}

// PostAbono maneja las peticiones POST para registrar un nuevo abono.
func (a *App) PostAbono() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto struct {
			ID             int         `json:"id"`
			CedulaPaciente string      `json:"cedula_paciente"`
			NombrePaciente string      `json:"nombre_paciente"`
			Patrocinante   string      `json:"patrocinante"`
			FechaAbono     interface{} `json:"fecha_abono"`
			Monto          float64     `json:"monto"`
			Tasa           float64     `json:"tasa"`
			MetodoPago     string      `json:"metodo_pago"`
			Referencia     string      `json:"referencia"`
			Observaciones  string      `json:"observaciones"`
			CreadoPor      string      `json:"creado_por"`
		}

		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "Cuerpo de solicitud inválido: " + err.Error()}, 400)
			return
		}

		abono := models.PacienteAbono{
			ID:             dto.ID,
			CedulaPaciente: dto.CedulaPaciente,
			NombrePaciente: dto.NombrePaciente,
			Patrocinante:   dto.Patrocinante,
			FechaAbono:     parseFlexibleTime(dto.FechaAbono),
			Monto:          dto.Monto,
			Tasa:           dto.Tasa,
			MetodoPago:     dto.MetodoPago,
			Referencia:     dto.Referencia,
			Observaciones:  dto.Observaciones,
			CreadoPor:      dto.CreadoPor,
		}

		if abono.CedulaPaciente == "" || abono.Monto <= 0 {
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "Cédula del paciente y monto válido son requeridos"}, 400)
			return
		}

		resp := a.DB.PostAbono(abono)
		sendResponse(w, r, resp, resp.Status)
	}
}

// DeleteAbono maneja las peticiones DELETE para eliminar un abono por ID.
func (a *App) DeleteAbono() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "ID de abono inválido"}, 400)
			return
		}

		resp := a.DB.DeleteAbono(id)
		sendResponse(w, r, resp, resp.Status)
	}
}

// GetConsumos maneja las peticiones GET para consultar consumos descontados de un paciente.
func (a *App) GetConsumos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cedula := r.URL.Query().Get("cedula")

		datos, resp := a.DB.GetConsumos(cedula)
		if resp.Status >= 400 {
			sendResponse(w, r, resp, resp.Status)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(datos)
	}
}

// PostConsumo maneja las peticiones POST para registrar un nuevo consumo descontado.
func (a *App) PostConsumo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto struct {
			ID             int         `json:"id"`
			IDAbono        *int        `json:"id_abono"`
			CedulaPaciente string      `json:"cedula_paciente"`
			NombrePaciente string      `json:"nombre_paciente"`
			IDCita         *int        `json:"id_cita"`
			FechaConsumo   interface{} `json:"fecha_consumo"`
			Especialidad   string      `json:"especialidad"`
			Servicio       string      `json:"servicio"`
			Monto          float64     `json:"monto"`
			Observaciones  string      `json:"observaciones"`
			CreadoPor      string      `json:"creado_por"`
		}

		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "Cuerpo de solicitud inválido: " + err.Error()}, 400)
			return
		}

		consumo := models.PacienteConsumo{
			ID:             dto.ID,
			IDAbono:        dto.IDAbono,
			CedulaPaciente: dto.CedulaPaciente,
			NombrePaciente: dto.NombrePaciente,
			IDCita:         dto.IDCita,
			FechaConsumo:   parseFlexibleTime(dto.FechaConsumo),
			Especialidad:   dto.Especialidad,
			Servicio:       dto.Servicio,
			Monto:          dto.Monto,
			Observaciones:  dto.Observaciones,
			CreadoPor:      dto.CreadoPor,
		}

		if consumo.CedulaPaciente == "" || consumo.Monto <= 0 || consumo.Especialidad == "" {
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "Cédula, especialidad y monto de consumo requeridos"}, 400)
			return
		}

		resp := a.DB.PostConsumo(consumo)
		sendResponse(w, r, resp, resp.Status)
	}
}

// DeleteConsumo maneja las peticiones DELETE para eliminar un consumo por ID.
func (a *App) DeleteConsumo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			sendResponse(w, r, models.Respuesta{Status: 400, Mensaje: "ID de consumo inválido"}, 400)
			return
		}

		resp := a.DB.DeleteConsumo(id)
		sendResponse(w, r, resp, resp.Status)
	}
}

// GetEstadoCuentaAbonos genera el reporte consolidado de rendición de cuentas para el patrocinante o paciente.
func (a *App) GetEstadoCuentaAbonos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cedula := r.URL.Query().Get("cedula")
		patrocinante := r.URL.Query().Get("patrocinante")
		desde := r.URL.Query().Get("desde")
		hasta := r.URL.Query().Get("hasta")

		datos, resp := a.DB.GetEstadoCuentaAbonos(cedula, patrocinante, desde, hasta)
		if resp.Status >= 400 {
			sendResponse(w, r, resp, resp.Status)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(datos)
	}
}

func parseFlexibleTime(val interface{}) time.Time {
	if val == nil {
		return time.Now()
	}
	str, ok := val.(string)
	if !ok || str == "" {
		return time.Now()
	}

	formats := []string{
		time.RFC3339,
		"2006-01-02T15:04:05.000Z07:00",
		"2006-01-02T15:04:05.000",
		"2006-01-02T15:04:05",
		"2006-01-02",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, str); err == nil {
			return t
		}
	}

	return time.Now()
}
