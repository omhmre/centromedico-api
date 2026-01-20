package models

type Compra struct {
	Id        int          `json:"id"`
	Idprov    int          `json:"idprov"`
	Proveedor string       `json:"proveedor"`
	Fecha     string       `json:"fecha"`
	Subtotal  float64      `json:"subtotal"`
	Dscto     float64      `json:"dscto"`
	Mototal   float64      `json:"mototal"`
	Deimp     string       `json:"deimp"`
	Tasaimp   float64      `json:"tasaimp"`
	Moimp     float64      `json:"moimp"`
	Moneto    float64      `json:"moneto"`
	Idsesion  int          `json:"idsesion"`
	Items     []ItemCompra `json:"items"`
}

type ItemCompra struct {
	Idcompra      int64   `json:"idcompra"`
	Codprod       string  `json:"codprod"`
	Deprod        string  `json:"deprod"`
	Cant          float64 `json:"cant"`
	Precio        float64 `json:"precio"`
	Subtotal      float64 `json:"subtotal"`
	Costo         float64 `json:"costo"`
	Subtotalcosto float64 `json:"subtotalcosto"`
}
