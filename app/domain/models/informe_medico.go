package models

import (
	"time"
)

// InformeMedico representa la estructura de un informe médico en la base de datos.
type InformeMedico struct {
	Id                    *int       `json:"id"`
	IdPaciente            *int       `json:"id_paciente"`
	Fecha                 time.Time  `json:"fecha"`
	IdDoctor              *int       `json:"id_doctor"`
	IdCita                *int       `json:"id_cita"`
	Diagnostico           string     `json:"diagnostico"`
	Evolucion             string     `json:"evolucion"`
	Plan                  string     `json:"plan"`
	Recomendaciones       string     `json:"recomendaciones"`
	Contenido             string     `json:"contenido"`
	Entregado             bool       `json:"entregado"`
	FechaEntrega          *time.Time `json:"fecha_entrega"`
	ModificadoPostEntrega bool       `json:"modificado_post_entrega"`
	UsuarioOperacion      string     `json:"usuario_operacion"`
	DoctorNombre   string `json:"doctor_nombre"`
	DoctorEspec    string `json:"doctor_espec"`
	EsMedico       bool   `json:"es_medico"`
	DoctorTitulo   string `json:"doctor_titulo"`
	DoctorTituloAcademico string `json:"doctor_titulo_academico"`
	DoctorMPPS     string `json:"doctor_mpps"`
	DoctorCM       string `json:"doctor_cm"`
	DoctorRIF      string `json:"doctor_rif"`
}



