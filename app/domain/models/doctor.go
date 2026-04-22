package models

import "fmt"

type DoctorService struct {
	Nombre   string  `json:"nombre"`
	Precio   float64 `json:"precio"`
	Comision float64 `json:"comision"`
}

type DoctoresModel struct {
	Id             int             `json:"id"`
	Nombres        string          `json:"nombres"`
	Servicios      []DoctorService `json:"servicios"`
	Dir            string          `json:"dir"`
	Correo         string          `json:"correo"`
	Whatsapp       string          `json:"whatsapp"`
	Instagram      string          `json:"instagram"`
	DaysOfWeek     []int           `json:"days_of_week"`
	StartTime      string          `json:"start_time"`
	EndTime        string          `json:"end_time"`
	SlotDuration   int             `json:"slot_duration"`
	MontoCita      float64         `json:"monto_cita"`
	EsMedico       bool            `json:"es_medico"`
	Titulo         string          `json:"titulo"`
	TituloAcademico string         `json:"titulo_academico"`
	NumMPPS        string          `json:"num_mpps"`
	NumCM          string          `json:"num_cm"`
	Rif            string          `json:"rif"`
	Cedula         string          `json:"cedula"`
	FechaNacimiento string         `json:"fecha_nacimiento"`
	FechaIngreso    string         `json:"fecha_ingreso"`
}

// Add String method for DoctoresModel
func (d DoctoresModel) String() string {
	return fmt.Sprintf("ID: %d, Nombres: %s, WhatsApp: %s, Correo: %s",
		d.Id, d.Nombres, d.Whatsapp, d.Correo)
}
