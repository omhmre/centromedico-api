package models

type Parametros struct {
	Id          int64  `json:"id"`
	Parametro   string `json:"parametro"`
	Descripcion string `json:"descripcion"`
	Valor       int64  `json:"valor"`
	Valores     string `json:"valores"`
	Descvalor   string `json:"descvalor"`
}

type RespParametros struct {
	Status  int          `json:"status"`
	Mensaje string       `json:"mensaje"`
	Data    []Parametros `json:"data"`
}
