package models

type InstrumentosPagos struct {
	Id          int     `json:"id"`
	Descripcion string  `json:"descripcion"`
	Tasa        float64 `json:"tasa"`
	Simbolo     string  `json:"simbolo"`
}

type Divisas struct {
	Id      int    `json:"id"`
	Divisa  string `json:"divisa"`
	Simbolo string `json:"simbolo"`
	Tasabs  string `json:"tasabs"`
	Fecha   string `json:"fechatasa"`
}

type RespDivisas struct {
	Status  int       `json:"status"`
	Mensaje string    `json:"mensaje"`
	Data    []Divisas `json:"divisas"`
}
