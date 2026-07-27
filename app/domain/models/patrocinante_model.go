package models

import "time"

// Patrocinante representa una empresa, fundación o particular patrocinador de becas médicas.
type Patrocinante struct {
	ID                int       `json:"id"`
	Nombre            string    `json:"nombre"`
	RIF               string    `json:"rif"`
	PersonaContacto   string    `json:"persona_contacto"`
	Telefono          string    `json:"telefono"`
	Email             string    `json:"email"`
	Tipo              string    `json:"tipo"` // 'Fundación', 'Corporativo', 'Particular', 'Donante Anónimo', 'ONG'
	SaldoTotalAbonado float64   `json:"saldo_total_abonado"`
	NroBecados        int       `json:"nro_becados"`
	Observaciones     string    `json:"observaciones"`
	Activo            bool      `json:"activo"`
	FechaRegistro     time.Time `json:"fecha_registro"`
}
