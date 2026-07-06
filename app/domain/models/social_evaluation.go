package models

import "time"

type SocialEvaluation struct {
	ID                     int       `json:"id"`
	CedulaPaciente         string    `json:"cedula_paciente"`
	IDEspecialista         int       `json:"id_especialista"`
	NombreEspecialista     string    `json:"nombre_especialista,omitempty"`
	
	// Datos Específicos del Menor
	LugarNacimiento        string    `json:"lugar_nacimiento"`
	GradoEscolar           string    `json:"grado_escolar"`
	Escolaridad            string    `json:"escolaridad"`
	ReferidoPor            string    `json:"referido_por"`

	// Datos de la Madre
	MadreNombre            string    `json:"madre_nombre"`
	MadreEdad              string    `json:"madre_edad"`
	MadreCI                string    `json:"madre_ci"`
	MadreTelefono          string    `json:"madre_telefono"`
	MadreOcupacion         string    `json:"madre_ocupacion"`
	MadreCorreo            string    `json:"madre_correo"`
	MadreDireccion         string    `json:"madre_direccion"`

	// Datos del Padre
	PadreNombre            string    `json:"padre_nombre"`
	PadreEdad              string    `json:"padre_edad"`
	PadreCI                string    `json:"padre_ci"`
	PadreTelefono          string    `json:"padre_telefono"`
	PadreOcupacion         string    `json:"padre_ocupacion"`
	PadreDireccion         string    `json:"padre_direccion"`

	// Cuerpo del Informe
	AntecedentesDesarrollo string    `json:"antecedentes_desarrollo"`
	GrupoFamiliar          string    `json:"grupo_familiar"`
	SituacionEconomica     string    `json:"situacion_economica"`
	ViviendaEntorno        string    `json:"vivienda_entorno"`
	AspectoSalud           string    `json:"aspecto_salud"`
	DiagnosticoSocial      string    `json:"diagnostico_social"`
	Conclusion             string    `json:"conclusion"`
	PlanAccion             string    `json:"plan_accion"`
	
	Entregado              bool      `json:"entregado"`
	UnlockedBy             *int      `json:"unlocked_by"`
	UnlockedAt             *time.Time`json:"unlocked_at"`
	UnlockReason           *string   `json:"unlock_reason"`
	
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}
