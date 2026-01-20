package models

import (
	"time"
)

// InformeMedico representa la estructura de un informe médico en la base de datos.
type InformeMedico struct {
	Id              int       `json:"id"`
	IdPaciente      int       `json:"id_paciente"`
	Fecha           time.Time `json:"fecha"`
	IdDoctor        int       `json:"id_doctor"`
	IdCita          int       `json:"id_cita,omitempty"` // Opcional
	Diagnostico     string    `json:"diagnostico"`
	Evolucion       string    `json:"evolucion"`
	Plan            string    `json:"plan"`
	Recomendaciones string    `json:"recomendaciones"`
}
