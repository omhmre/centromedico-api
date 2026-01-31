package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"omhmre.com/centromedico/app/domain/models"
	"omhmre.com/centromedico/app/domain/utils"
)

func (a *App) AddCita() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var citas []models.CitaModel
		err := json.NewDecoder(r.Body).Decode(&citas)
		if err != nil {
			respuesta := models.Respuesta{
				Status:  http.StatusBadRequest,
				Mensaje: "Cuerpo de la solicitud inválido. Se esperaba un array de citas: " + err.Error(),
			}
			sendResponse(w, r, respuesta, http.StatusBadRequest)
			return
		}

		if len(citas) == 0 {
			respuesta := models.Respuesta{
				Status:  http.StatusBadRequest,
				Mensaje: "El array de citas no puede estar vacío.",
			}
			sendResponse(w, r, respuesta, http.StatusBadRequest)
			return
		}

		rp := a.DB.AddCita(citas)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			// Notificar a todos los clientes sobre el cambio
			message, _ := json.Marshal(map[string]string{"event": "CITAS_UPDATED"})
			a.Hub.Broadcast <- message
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) UpdateCita() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cita models.CitaModel
		err := json.NewDecoder(r.Body).Decode(&cita)
		if err != nil {
			respuesta := models.Respuesta{
				Status:  http.StatusBadRequest,
				Mensaje: "Cuerpo de la solicitud inválido: " + err.Error(),
			}
			sendResponse(w, r, respuesta, http.StatusBadRequest)
			return
		}
		rp := a.DB.UpdateCita(cita)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			message, _ := json.Marshal(map[string]string{"event": "CITAS_UPDATED"})
			a.Hub.Broadcast <- message
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) AddDiagnosis() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cita models.CitaModel
		err := json.NewDecoder(r.Body).Decode(&cita)
		if err != nil {
			respuesta := models.Respuesta{
				Status:  http.StatusBadRequest,
				Mensaje: "Cuerpo de la solicitud inválido: " + err.Error(),
			}
			sendResponse(w, r, respuesta, http.StatusBadRequest)
			return
		}
		rp := a.DB.UpdateDiagnosticoCita(cita)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			message, _ := json.Marshal(map[string]string{"event": "CITAS_UPDATED"})
			a.Hub.Broadcast <- message
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) GetCitas() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.Respuesta
		datos, err := a.DB.GetCitas()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Citas"
			utils.CreateLog("Error al obteniendo las citas: " + err.Mensaje)
		}
		rp.Status = 10
		rp.Mensaje = "Citas listadas correctamente!"

		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) GetCitasPaciente() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.Respuesta
		var paciente models.PacientesModel
		json.NewDecoder(r.Body).Decode(&paciente)
		datos, err := a.DB.GetCitasPaciente(paciente)
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Citas"
			utils.CreateLog("Error al obteniendo las citas: " + err.Mensaje)
		}
		rp.Status = 10
		rp.Mensaje = "Citas listadas correctamente!"

		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) DelCita() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var lacita models.IdCitas
		json.NewDecoder(r.Body).Decode(&lacita)

		rp := a.DB.DelCita(lacita)
		// La lógica original aquí era un poco confusa. Asumimos que un status >= 400 es un error.
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
			return
		}
		message, _ := json.Marshal(map[string]string{"event": "CITAS_UPDATED"})
		a.Hub.Broadcast <- message
		sendResponse(w, r, rp, http.StatusOK)
	}
}

func (a *App) GetDoctores() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.Respuesta
		datos, err := a.DB.GetDoctores()
		if err.Status > 200 {
			rp.Status = err.Status
			rp.Mensaje = "Error Cargando Doctores"
			utils.CreateLog("Error al obteniendo los doctores: " + err.Mensaje)
		}
		rp.Status = 10
		rp.Mensaje = "Doctores listados correctamente!"

		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) UpdateDoctores() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var doctor models.DoctoresModel
		err := json.NewDecoder(r.Body).Decode(&doctor)
		if err != nil {
			respuesta := models.Respuesta{
				Status:  http.StatusBadRequest,
				Mensaje: "Cuerpo de la solicitud inválido: " + err.Error(),
			}
			sendResponse(w, r, respuesta, http.StatusBadRequest)
			return
		}
		rp := a.DB.UpdDoctores(doctor)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) GetPacientes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Llamar a la función de la base de datos sin paginación.
		pacientes, rp := a.DB.GetPacientes()
		if rp.Status >= 400 {
			sendResponse(w, r, rp, rp.Status)
			return
		}

		// Enviar la lista de pacientes en el cuerpo de la respuesta.
		sendResponse(w, r, pacientes, http.StatusOK)
	}
}

func (a *App) PostPaciente() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var paciente models.PacientesModel
		err := json.NewDecoder(r.Body).Decode(&paciente)
		if err != nil {
			respuesta := models.Respuesta{
				Status:  http.StatusBadRequest,
				Mensaje: "Cuerpo de la solicitud inválido: " + err.Error(),
			}
			sendResponse(w, r, respuesta, http.StatusBadRequest)
			return
		}
		rp := a.DB.PostPaciente(paciente)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) UpdPaciente() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var paciente models.PacientesModel
		err := json.NewDecoder(r.Body).Decode(&paciente)
		if err != nil {
			respuesta := models.Respuesta{
				Status:  http.StatusBadRequest,
				Mensaje: "Cuerpo de la solicitud inválido: " + err.Error(),
			}
			sendResponse(w, r, respuesta, http.StatusBadRequest)
			return
		}
		rp := a.DB.UpdPaciente(paciente)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) DelPaciente() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var paciente models.PacientesModel
		var resp models.Respuesta
		err := json.NewDecoder(r.Body).Decode(&paciente)
		if err != nil {
			// log.Printf("Unable to decode the request body.  %v", err)
		}
		rp := a.DB.DelPaciente(paciente)
		if rp.Status >= 200 {
			resp.Status = 503
			resp.Mensaje = "No se pudo eliminar el Paciente. Error=%v \n"
			sendResponse(w, r, rp, http.StatusInternalServerError)
			return
		}
		resp.Status = 201
		resp.Mensaje = rp.Mensaje
		sendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) PostDoctor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var doctor models.DoctoresModel
		err := json.NewDecoder(r.Body).Decode(&doctor)
		if err != nil {
			respuesta := models.Respuesta{
				Status:  http.StatusBadRequest,
				Mensaje: "Cuerpo de la solicitud inválido: " + err.Error(),
			}
			sendResponse(w, r, respuesta, http.StatusBadRequest)
			return
		}
		rp := a.DB.PostDoctor(doctor)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) DelDoctor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var doctor models.DoctoresModel
		var resp models.Respuesta
		err := json.NewDecoder(r.Body).Decode(&doctor)
		if err != nil {
			// log.Printf("Unable to decode the request body.  %v", err)
			utils.CreateLog("Unable to decode the doctor request body.  " + err.Error())
		}
		rp := a.DB.DelDoctor(doctor)
		if rp.Status >= 200 {
			resp.Status = 503
			resp.Mensaje = "No se pudo eliminar el Especialista. Error=%v \n"
			sendResponse(w, r, rp, http.StatusInternalServerError)
			return
		}
		resp.Status = 201
		resp.Mensaje = rp.Mensaje
		sendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) GetPayments() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.Respuesta
		var id models.Id
		err := json.NewDecoder(r.Body).Decode(&id)
		if err != nil {
			utils.CreateLog("Unable to decode the doctor request body.  " + err.Error())
		}

		datos, err2 := a.DB.GetPayments(id)
		if err2.Status > 200 {
			rp.Status = err2.Status
			rp.Mensaje = "Error Cargando Pagos"
			utils.CreateLog("Error al obteniendo los pagos: " + err2.Mensaje)
		}
		rp.Status = 10
		rp.Mensaje = "Pagos listados correctamente!"

		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) PostPayments() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pagos []models.Payments
		err := json.NewDecoder(r.Body).Decode(&pagos)
		if err != nil {
			utils.CreateLog("Unable to decode the request body.  " + err.Error())
			respuesta := models.Respuesta{
				Status:  http.StatusBadRequest,
				Mensaje: "Cuerpo de la solicitud inválido: " + err.Error(),
			}
			sendResponse(w, r, respuesta, http.StatusBadRequest)
			return
		}
		rp := a.DB.PostPayments(pagos)
		if rp.Status >= 400 {
			sendResponse(w, r, rp, http.StatusInternalServerError)
		} else {
			sendResponse(w, r, rp, http.StatusOK)
		}
	}
}

func (a *App) GetRelPagos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var rp models.Respuesta
		var id models.Fechas
		err := json.NewDecoder(r.Body).Decode(&id)
		if err != nil {
			utils.CreateLog("Unable to decode the payments request body.  " + err.Error())
		}

		datos, err2 := a.DB.GetRelPagos(id)
		if err2.Status > 200 {
			rp.Status = err2.Status
			rp.Mensaje = "Error Cargando Pagos"
			utils.CreateLog("Error al obteniendo los pagos: " + err2.Mensaje)
		}
		rp.Status = 10
		rp.Mensaje = "Pagos listados correctamente!"

		data := json.NewEncoder(w).Encode(datos)
		if data != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", data)
			return
		}
	}
}

func (a *App) DelPayment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payment models.Id
		var resp models.Respuesta
		err := json.NewDecoder(r.Body).Decode(&payment)
		if err != nil {
			// log.Printf("Unable to decode the request body.  %v", err)
			utils.CreateLog("Unable to decode the doctor request body.  " + err.Error())
		}
		rp := a.DB.DelPayment(payment)
		if rp.Status >= 200 {
			resp.Status = 503
			resp.Mensaje = "No se pudo eliminar el pago. Error=%v \n"
			sendResponse(w, r, rp, http.StatusInternalServerError)
			return
		}
		resp.Status = 201
		resp.Mensaje = rp.Mensaje
		sendResponse(w, r, resp, http.StatusOK)
	}
}

// UpsertPrecioEspecialidad maneja la creación/actualización de un precio especial para un paciente.
func (a *App) UpsertPrecioEspecialidad() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var precio models.PrecioEspecialidad
		if err := json.NewDecoder(r.Body).Decode(&precio); err != nil {
			respuesta := models.Respuesta{
				Status:  http.StatusBadRequest,
				Mensaje: "Cuerpo de la solicitud inválido: " + err.Error(),
			}
			sendResponse(w, r, respuesta, http.StatusBadRequest)
			return
		}

		respuesta := a.DB.UpsertPrecioEspecialidad(precio)

		httpStatus := http.StatusOK
		if respuesta.Status >= 400 {
			// Mapear el código de estado de la lógica de negocio a un código HTTP.
			switch respuesta.Status {
			case 400:
				httpStatus = http.StatusBadRequest
			default:
				httpStatus = http.StatusInternalServerError
			}
		}
		sendResponse(w, r, respuesta, httpStatus)
	}
}

// DelPrecioEspecialidad maneja la eliminación de un precio especial.
func (a *App) DelPrecioEspecialidad() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var precio models.PrecioEspecialidad
		if err := json.NewDecoder(r.Body).Decode(&precio); err != nil {
			respuesta := models.Respuesta{
				Status:  http.StatusBadRequest,
				Mensaje: "Cuerpo de la solicitud inválido: " + err.Error(),
			}
			sendResponse(w, r, respuesta, http.StatusBadRequest)
			return
		}

		respuesta := a.DB.DelPrecioEspecialidad(precio)

		httpStatus := http.StatusOK
		if respuesta.Status >= 400 {
			// Mapear el código de estado de la lógica de negocio a un código HTTP.
			switch respuesta.Status {
			case 400:
				httpStatus = http.StatusBadRequest
			case 404:
				httpStatus = http.StatusNotFound
			default:
				httpStatus = http.StatusInternalServerError
			}
		}
		sendResponse(w, r, respuesta, httpStatus)
	}
}

func (a *App) GetCitasFecha() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var fechas models.Fechas
		var rp models.Respuesta

		err := json.NewDecoder(r.Body).Decode(&fechas)
		if err != nil {
			rp.Status = http.StatusBadRequest
			rp.Mensaje = "Cuerpo de la solicitud inválido: " + err.Error()
			sendResponse(w, r, rp, http.StatusBadRequest)
			return
		}

		datos, resp := a.DB.GetCitasFecha(fechas)
		if resp.Status >= 400 {
			sendResponse(w, r, resp, resp.Status)
			return
		}

		errEncode := json.NewEncoder(w).Encode(datos)
		if errEncode != nil {
			fmt.Fprintf(w, "Cannot format json, err=%v/n", errEncode)
		}
	}
}

// HandlePreciosEspecialidad dispatches requests for /pacientes/precios based on the HTTP method.
func (a *App) HandlePreciosEspecialidad() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost, http.MethodPut:
			// The UpsertPrecioEspecialidad handler will be called for POST or PUT requests.
			a.UpsertPrecioEspecialidad().ServeHTTP(w, r)
		case http.MethodDelete:
			// The DelPrecioEspecialidad handler will be called for DELETE requests.
			a.DelPrecioEspecialidad().ServeHTTP(w, r)
		default:
			// For any other method, respond with a 405 Method Not Allowed error.
			w.Header().Set("Allow", "POST, PUT, DELETE")
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	}
}

// ExchangeRateUpdateResponse defines the structure for the response of the exchange rate update endpoint.
type ExchangeRateUpdateResponse struct {
	Status  int     `json:"status"`
	Mensaje string  `json:"mensaje"`
	Rate    float64 `json:"rate,omitempty"` // Use omitempty if rate might not always be present (e.g., on error)
}

// UpdateExchangeRateAndAppointments fetches the latest exchange rate and updates relevant appointments.
func (a *App) UpdateExchangeRateAndAppointments() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Fetch the latest exchange rate
		newRate, rpFetch := a.DB.FetchExchangeRate()

		if rpFetch.Status != http.StatusOK {
			// If fetching rate fails, return the error from FetchExchangeRate
			response := ExchangeRateUpdateResponse{
				Status:  rpFetch.Status,
				Mensaje: rpFetch.Mensaje,
				Rate:    0.0, // Rate is not available on fetch error
			}
			sendResponse(w, r, response, rpFetch.Status)
			return
		}

		// 2. Update appointments in the database
		rpUpdate := a.DB.UpdateUnpaidAppointmentsVESRate(newRate)
		if rpUpdate.Status != http.StatusOK {
			sendResponse(w, r, rpUpdate, rpUpdate.Status)
			return
		}

		// If both operations are successful
		response := ExchangeRateUpdateResponse{
			Status:  http.StatusOK,
			Mensaje: rpUpdate.Mensaje, // Use the message from the update operation
			Rate:    newRate,
		}
		sendResponse(w, r, response, http.StatusOK)
	}
}
