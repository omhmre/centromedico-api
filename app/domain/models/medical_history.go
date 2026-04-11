package models

import "time"

type SignosVitales struct {
	Id                    int       `json:"id"`
	IdPaciente            int       `json:"id_paciente"`
	IdCita                *int      `json:"id_cita"`
	Fecha                 time.Time `json:"fecha"`
	TensionArterial       string    `json:"tension_arterial"`
	FrecuenciaCardiaca    int       `json:"frecuencia_cardiaca"`
	FrecuenciaRespiratoria int       `json:"frecuencia_respiratoria"`
	Temperatura           float64   `json:"temperatura"`
	SaturacionOxigeno     int       `json:"saturacion_oxigeno"`
	Peso                  float64   `json:"peso"`
	Talla                 float64   `json:"talla"`
	Imc                   float64   `json:"imc"`
	Notas                 string    `json:"notas"`
	UsuarioOperacion      string    `json:"usuario_operacion"`
}

type Antecedentes struct {
	IdPaciente         int       `json:"id_paciente"`
	Medicos            string    `json:"medicos"`
	Quirurgicos        string    `json:"quirurgicos"`
	Alergicos          string    `json:"alergicos"`
	Familiares         string    `json:"familiares"`
	Habitos            string    `json:"habitos"`
	Otros              string    `json:"otros"`
	UltimaActualizacion time.Time `json:"ultima_actualizacion"`
}

type HistoryTimelineItem struct {
	Fecha       time.Time `json:"fecha"`
	Tipo        string    `json:"tipo"` // Consulta, Examen, Cirugía
	Detalle     string    `json:"detalle"`
	Doctor      string    `json:"doctor"`
	Especialidad string    `json:"especialidad"`
	IdReferencia int       `json:"id_referencia"`
}

type PatientMedicalHistory struct {
	Paciente     PacientesModel        `json:"paciente"`
	Antecedentes Antecedentes          `json:"antecedentes"`
	Timeline     []HistoryTimelineItem `json:"timeline"`
	LatestVitals *SignosVitales        `json:"latest_vitals"`
}
