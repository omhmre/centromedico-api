package models

import (
	"fmt"
	"time"
)

type CitaModel struct {
	Id           int       `json:"id"`
	IdDoctor     int       `json:"iddoctor"`
	Especialista string    `json:"especialista"`
	Especialidad string    `json:"especialidad"`
	Cedula       string    `json:"cedula"`
	Paciente     string    `json:"paciente"`
	Motivo       string    `json:"motivo"`
	Inicio       time.Time `json:"inicio"`
	Fin          time.Time `json:"fin"`
	Diagnostico  *string   `json:"diagnostico,omitempty"`
	Status       string    `json:"status"`
	Color        string    `json:"color"`
	Montoref     float64   `json:"montoref"`
	Tasa         float64   `json:"tasa"`
	Montobs      float64   `json:"montobs"`
	Pagado       float64   `json:"pagado"`
	Saldo        float64   `json:"saldo"`
	GroupID      *string   `json:"group_id,omitempty"`
	Weeks        int       `json:"weeks"`
	UpdateSeries bool      `json:"update_series"`
}

// Add String method for CitaModel
func (c CitaModel) String() string {
	return fmt.Sprintf("ID: %d, Doctor ID: %d, Paciente: %s, Motivo: %s, Inicio: %s, Fin: %s, Status: %s",
		c.Id, c.IdDoctor, c.Paciente, c.Motivo, c.Inicio.Format(time.RFC3339), c.Fin.Format(time.RFC3339), c.Status)
}
