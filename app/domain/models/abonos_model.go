package models

import "time"

// PacienteAbono representa un pago adelantado realizado por un paciente o empresa patrocinadora (ej. Digitel).
type PacienteAbono struct {
	ID             int       `json:"id"`
	CedulaPaciente string    `json:"cedula_paciente"`
	NombrePaciente string    `json:"nombre_paciente,omitempty"`
	Patrocinante   string    `json:"patrocinante"` // Ej: 'Digitel', 'Particular', 'Fundafi'
	FechaAbono     time.Time `json:"fecha_abono"`
	Monto          float64   `json:"monto"`
	Tasa           float64   `json:"tasa"`
	MetodoPago     string    `json:"metodo_pago"`
	Referencia     string    `json:"referencia"`
	Observaciones  string    `json:"observaciones"`
	CreadoPor      string    `json:"creado_por,omitempty"`
	FechaCreacion  time.Time `json:"fecha_creacion,omitempty"`
}

// PacienteConsumo representa la deducción/gasto de un abono por concepto de terapia o especialidad.
type PacienteConsumo struct {
	ID             int       `json:"id"`
	IDAbono        *int      `json:"id_abono,omitempty"`
	CedulaPaciente string    `json:"cedula_paciente"`
	NombrePaciente string    `json:"nombre_paciente,omitempty"`
	IDCita         *int      `json:"id_cita,omitempty"`
	FechaConsumo   time.Time `json:"fecha_consumo"`
	Especialidad   string    `json:"especialidad"`
	Servicio       string    `json:"servicio"`
	Monto          float64   `json:"monto"`
	Observaciones  string    `json:"observaciones"`
	CreadoPor      string    `json:"creado_por,omitempty"`
	FechaCreacion  time.Time `json:"fecha_creacion,omitempty"`
}

// EstadoCuentaAbonos representa la rendición de cuentas (Semestral / Período) para la empresa o paciente.
type EstadoCuentaAbonos struct {
	CedulaPaciente  string            `json:"cedula_paciente"`
	NombrePaciente  string            `json:"nombre_paciente"`
	Patrocinante    string            `json:"patrocinante"`
	TotalAbonado    float64           `json:"total_abonado"`
	TotalConsumido  float64           `json:"total_consumido"`
	SaldoDisponible float64           `json:"saldo_disponible"`
	Abonos          []PacienteAbono   `json:"abonos"`
	Consumos        []PacienteConsumo `json:"consumos"`
}
