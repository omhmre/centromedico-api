package models

type Respuesta struct {
	Status  int    `json:"status"`
	Mensaje string `json:"mensaje"`
}

type Id struct {
	Id string `json:"id"`
}

type IdCitas struct {
	Id   int `json:"id"`
	Tipo int `json:"tipo"`
}

type Fechas struct {
	Desde string `json:"desde"`
	Hasta string `json:"hasta"`
}
type MailSend struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Archivo string `json:"archivo"`
}

type EmailConfig struct {
	Id      int    `json:"id"`
	Smtp    string `json:"smtp"`
	Puerto  int    `json:"puerto"`
	Usuario string `json:"usuario"`
	Clave   string `json:"clave"`
	Tls     bool   `json:"tls"`
}

// ─── Modelos de Inteligencia de Negocio ───────────────────────────────────────

type BIResumen struct {
	TotalCitas       int     `json:"totalCitas"`
	Completadas      int     `json:"completadas"`
	Canceladas       int     `json:"canceladas"`
	Pendientes       int     `json:"pendientes"`
	TotalIngresosUSD float64 `json:"totalIngresosUSD"`
	IngresoPorCita   float64 `json:"ingresoPorCita"`
	PacientesUnicos  int     `json:"pacientesUnicos"`
	TasaCompletadas  float64 `json:"tasaCompletadas"`
}

type BITimeSeries struct {
	Fecha    string  `json:"fecha"`
	Citas    int     `json:"citas"`
	Ingresos float64 `json:"ingresos"`
}

type BIEspecialidad struct {
	Especialidad   string  `json:"especialidad"`
	TotalCitas     int     `json:"totalCitas"`
	Completadas    int     `json:"completadas"`
	TotalIngresos  float64 `json:"totalIngresos"`
	TasaEficiencia float64 `json:"tasaEficiencia"`
}

type BIDoctor struct {
	Nombre      string  `json:"nombre"`
	TotalCitas  int     `json:"totalCitas"`
	Completadas int     `json:"completadas"`
	Ingresos    float64 `json:"ingresos"`
	Eficiencia  float64 `json:"eficiencia"`
}

type BIMetodoPago struct {
	Metodo     string  `json:"metodo"`
	Total      float64 `json:"total"`
	Porcentaje float64 `json:"porcentaje"`
}

type BIHeatmapCell struct {
	DiaSemana int `json:"diaSemana"`
	Hora      int `json:"hora"`
	Cantidad  int `json:"cantidad"`
}
