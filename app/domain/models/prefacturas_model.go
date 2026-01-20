package models

type Prefactura struct {
	Id          int     `json:"id"`
	Idcliente   string  `json:"idcliente"`
	Cliente     string  `json:"cliente"`
	Fecha       string  `json:"fecha"`
	Diasvence   int     `json:"diasvence"`
	Vence       string  `json:"vence"`
	Subtotal    float64 `json:"subtotal"`
	Dscto       float64 `json:"dscto"`
	Mototal     float64 `json:"mototal"`
	Deimp       string  `json:"deimp"`
	Tasaimp     float64 `json:"tasaimp"`
	Moimp       float64 `json:"moimp"`
	Moneto      float64 `json:"moneto"`
	Monetodiv   float64 `json:"monetodiv"`
	Idvendedor  int     `json:"idvendedor"`
	Idsesion    int     `json:"idsesion"`
	Idmesa      int     `json:"idmesa"`
	Idmesonero  string  `json:"idmesonero"`
	Pagado      float64 `json:"pagado"`
	Porpagar    float64 `json:"porpagar"`
	Cambio      float64 `json:"cambio"`
	Tasadiv     float32 `json:"tasadiv"`
	Cxcbs       float64 `json:"cxcbs"`
	Cxcdiv      float64 `json:"cxcdiv"`
	Condiciones string  `json:"condiciones"`
	Items       []Item  `json:"items"`
}

type RespPrefactura struct {
	Status  int          `json:"status"`
	Mensaje string       `json:"mensaje"`
	Data    []Prefactura `json:"data"`
}

type Item struct {
	Idprefact   int     `json:"idprefact"`
	Codprod     string  `json:"codprod"`
	Producto    string  `json:"producto"`
	Cant        float32 `json:"cant"`
	Precio      float64 `json:"precio"`
	Subtotal    float64 `json:"subtotal"`
	Descuento   float64 `json:"descuento"`
	Neto        float64 `json:"neto"`
	Descripcion string  `json:"descripcion"`
	Cantmp      float32 `json:"cantmp"`
	Cantpres    float32 `json:"cantpres"`
	Iditempres  int     `json:"iditempres"`
}
