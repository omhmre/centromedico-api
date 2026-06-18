package models

import "time"

type SocialEvaluation struct {
	ID                 int       `json:"id"`
	CedulaPaciente     string    `json:"cedula_paciente"`
	IDEspecialista     int       `json:"id_especialista"`
	NombreEspecialista string    `json:"nombre_especialista,omitempty"`
	GrupoFamiliar      string    `json:"grupo_familiar"`
	SituacionEconomica string    `json:"situacion_economica"`
	ViviendaEntorno    string    `json:"vivienda_entorno"`
	AspectoSalud       string    `json:"aspecto_salud"`
	DiagnosticoSocial  string    `json:"diagnostico_social"`
	PlanAccion         string    `json:"plan_accion"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}
