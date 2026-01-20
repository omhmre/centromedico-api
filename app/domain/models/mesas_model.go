package models

type Mesas struct {
	Id           int     `json:"id"`
	Nombre       string  `json:"nombre"`
	Subtotal     float32 `json:"subtotal"`
	Abierta      int     `json:"abierta"`
	Idprefactura int     `json:"idprefactura"`
	Idcliente    string  `json:"idcliente"`
	Cliente      string  `json:"cliente"`
	Idmesonero   string  `json:"idmesonero"`
	Mesonero     string  `json:"mesonero"`
	Inicio       string  `json:"inicio"`
	Fin          string  `json:"fin"`
}

type RespMesas struct {
	Status  int     `json:"status"`
	Mensaje string  `json:"mensaje"`
	Data    []Mesas `json:"data"`
}
