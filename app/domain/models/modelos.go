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
