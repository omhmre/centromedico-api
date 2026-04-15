package models

import "time"

type NominaModel struct {
	Id             int       `json:"id"`
	PersonalId     int       `json:"personal_id"`
	NombrePersonal string    `json:"nombre_personal"` // Extra field for UI join
	FechaInicio    time.Time `json:"fecha_inicio"`
	FechaFin       time.Time `json:"fecha_fin"`
	TipoPeriodo    string    `json:"tipo_periodo"`
	MontoBase      float64   `json:"monto_base"`
	Bonificaciones float64   `json:"bonificaciones"`
	Deducciones    float64   `json:"deducciones"`
	MontoTotal     float64   `json:"monto_total"`
	Status         string    `json:"status"`
	FechaPago      *time.Time `json:"fecha_pago"`
	EgresoId       int       `json:"egreso_id"`
	Notas          string    `json:"notas"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type RespNomina struct {
	Status  int           `json:"status"`
	Mensaje string        `json:"mensaje"`
	Data    []NominaModel `json:"data"`
}
