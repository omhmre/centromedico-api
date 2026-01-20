package models

import (
	"fmt"
	"time"
)

type PacientesModel struct {
	Id            int                    `json:"id"`
	Cedula        *string                `json:"cedula,omitempty"`
	Nombres       string                 `json:"nombres"`
	Fenac         time.Time              `json:"fenac"`
	Representante *string                `json:"representante,omitempty"`
	Whatsapp      *string                `json:"whatsapp,omitempty"`
	Direccion     *string                `json:"direccion,omitempty"`
	Correo        *string                `json:"correo,omitempty"`
	Diagnostico   *string                `json:"diagnostico,omitempty"`
	CXC           float64                `json:"cxc"`
	CreatedAt     time.Time              `json:"createdAt"`
	Citas         []CitaModel            `json:"citas,omitempty"`
	Especialistas []EspecialistaAtencion `json:"especialistas,omitempty"`
	Precios       []PrecioEspecialidad   `json:"precios,omitempty"`
	Pagos         []Payments             `json:"pagos,omitempty"`
}

// PrecioEspecialidad representa un precio personalizado para un paciente en una especialidad específica.
type PrecioEspecialidad struct {
	IDPaciente   int     `json:"id_paciente,omitempty"` // Usado para el CRUD de precios
	Especialidad string  `json:"especialidad"`
	Precio       float64 `json:"precio"`
}

// EspecialistaAtencion representa a un doctor que ha atendido a un paciente.
type EspecialistaAtencion struct {
	ID           int    `json:"id_doctor"`
	Nombres      string `json:"nombres"`
	Especialidad string `json:"especialidad"`
}

// Add String method for PacientesModel
func (p PacientesModel) String() string {
	cedula := "N/A"
	if p.Cedula != nil {
		cedula = *p.Cedula
	}
	return fmt.Sprintf("ID: %d, Cédula: %s, Nombres: %s, Fecha de Nacimiento: %s",
		p.Id, cedula, p.Nombres, p.Fenac.Format("2006-01-02"))
}
