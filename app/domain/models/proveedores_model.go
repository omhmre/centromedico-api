package models

type Proveedor struct {
	Id        int64   `json:"id"`
	Tipo      string  `json:"tipo"`
	Proveedor string  `json:"proveedor"`
	Rif       string  `json:"rif"`
	Dirfiscal string  `json:"dirfiscal"`
	Ciudad    string  `json:"ciudad"`
	Estado    string  `json:"estado"`
	Telf      string  `json:"telf"`
	Correo    string  `json:"correo"`
	Twitter   string  `json:"twitter"`
	Facebook  string  `json:"facebook"`
	Status    string  `json:"status"`
	Obs       string  `json:"obs"`
	Clasif    string  `json:"clasif"`
	Credito   int64   `json:"credito"`
	Diascred  int64   `json:"diascred"`
	Cxp       float32 `json:"cxp"`
}

type RespProveedores struct {
	Status  int         `json:"status"`
	Mensaje string      `json:"mensaje"`
	Data    []Proveedor `json:"data"`
}
