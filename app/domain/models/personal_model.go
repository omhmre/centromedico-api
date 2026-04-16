package models

import "time"

type PersonalModel struct {
	Id              *int       `json:"id"`
	Nombre          string     `json:"nombre"`
	Cedula          string     `json:"cedula"`
	Telefono        *string    `json:"telefono"`
	Correo          *string    `json:"correo"`
	Direccion       *string    `json:"direccion"`
	Titulo          *string    `json:"titulo"`
	Universidad     *string    `json:"universidad"`
	FechaIngreso    *time.Time `json:"fecha_ingreso"`
	FechaNacimiento *time.Time `json:"fecha_nacimiento"`
	Cargo           *string    `json:"cargo"`
	Sueldo          float64    `json:"sueldo"`
	FrecuenciaPago  string     `json:"frecuencia_pago"`
	Status          string     `json:"status"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

type RespPersonal struct {
	Status  int             `json:"status"`
	Mensaje string          `json:"mensaje"`
	Data    []PersonalModel `json:"data"`
}
