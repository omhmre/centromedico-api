package models

type CxcResumen struct {
	Id       int64   `json:"id"`
	Idfact   int64   `json:"idfact"`
	Codclie  string  `json:"codclie"`
	Nombre   string  `json:"nombre"`
	Saldobs  float64 `json:"saldobs"`
	Saldodiv float64 `json:"saldodiv"`
}

type CxcVencida struct {
	Rango string  `json:"rango"`
	Saldo float64 `json:"saldo"`
}
