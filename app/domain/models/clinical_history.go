package models

import "time"

type ClinicalHistory struct {
	ID                     int       `json:"id"`
	CedulaPaciente         string    `json:"cedula_paciente"`
	AntecedentesFamiliares string    `json:"antecedentes_familiares"`
	PatologiasPrevias      string    `json:"patologias_previas"`
	Alergias               string    `json:"alergias"`
	Habitos                string    `json:"habitos"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

type ClinicalRecord struct {
	ID                       int       `json:"id"`
	CedulaPaciente           string    `json:"cedula_paciente"`
	IDEspecialista           int       `json:"id_especialista"`
	NombreEspecialista       string    `json:"nombre_especialista,omitempty"`
	EspecialidadEspecialista string    `json:"especialidad_especialista,omitempty"`
	MotivoConsulta           string    `json:"motivo_consulta"`
	ExamenFisico             string    `json:"examen_fisico"`
	Diagnostico              string    `json:"diagnostico"`
	Tratamiento              string    `json:"tratamiento"`
	Observaciones            string    `json:"observaciones"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`

	Attachments []ClinicalAttachment `json:"attachments,omitempty"`
}

type ClinicalAttachment struct {
	ID            int       `json:"id"`
	IDRegistro    int       `json:"id_registro"`
	NombreArchivo string    `json:"nombre_archivo"`
	TipoArchivo   string    `json:"tipo_archivo"`
	UrlArchivo    string    `json:"url_archivo"`
	CreatedAt     time.Time `json:"created_at"`
}
