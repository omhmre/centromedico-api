package models

type Factura struct {
	Id         int            `json:"id"`
	Idcliente  string         `json:"idcliente"`
	Cliente    string         `json:"cliente"`
	Dirfiscal  string         `json:"dirfiscal"`
	Rif        string         `json:"rif"`
	Persconta  string         `json:"persconta"`
	Tlfconta   string         `json:"tlfconta"`
	Fecha      string         `json:"fecha"`
	Subtotal   float32        `json:"subtotal"`
	Dscto      float32        `json:"dscto"`
	Mototal    float32        `json:"mototal"`
	Deimp      string         `json:"deimp"`
	Tasaimp    float32        `json:"tasaimp"`
	Moimp      float32        `json:"moimp"`
	Moneto     float32        `json:"moneto"`
	Tasadiv    float32        `json:"tasadiv"`
	Monetodiv  float32        `json:"monetodiv"`
	Idvendedor int            `json:"idvendedor"`
	Vendedor   string         `json:"vendedor"`
	Idsesion   int            `json:"idsesion"`
	Idmesa     int            `json:"idmesa"`
	Idmesonero string         `json:"idmesonero"`
	Pagado     float32        `json:"pagado"`
	Porpagar   float32        `json:"porpagar"`
	Cambio     float32        `json:"cambio"`
	Cxcbs      float32        `json:"cxcbs"`
	Cxcdiv     float32        `json:"cxcdiv"`
	Valido     bool           `json:"valido"`
	Esdivisa   bool           `json:"esdivisa"`
	Items      []ItemsFactura `json:"items"`
	Pagos      []DetPago      `json:"pagos"`
}

type Presupuestos struct {
	Id          int                 `json:"id"`
	Idcliente   string              `json:"idcliente"`
	Cliente     string              `json:"cliente"`
	Dirfiscal   string              `json:"dirfiscal"`
	Rif         string              `json:"rif"`
	Persconta   string              `json:"persconta"`
	Tlfconta    string              `json:"tlfconta"`
	Fecha       string              `json:"fecha"`
	Diasvence   int                 `json:"diasvence"`
	Vence       string              `json:"vence"`
	Subtotal    float32             `json:"subtotal"`
	Dscto       float32             `json:"dscto"`
	Mototal     float32             `json:"mototal"`
	Deimp       string              `json:"deimp"`
	Tasaimp     float32             `json:"tasaimp"`
	Moimp       float32             `json:"moimp"`
	Moneto      float32             `json:"moneto"`
	Tasadiv     float32             `json:"tasadiv"`
	Monetodiv   float32             `json:"monetodiv"`
	Idvendedor  int                 `json:"idvendedor"`
	Vendedor    string              `json:"vendedor"`
	Idsesion    int                 `json:"idsesion"`
	Idmesa      int                 `json:"idmesa"`
	Idmesonero  string              `json:"idmesonero"`
	Condiciones string              `json:"condiciones"`
	Items       []ItemsPresupuestos `json:"items"`
}

type RespFactura struct {
	Status  int     `json:"status"`
	Mensaje string  `json:"mensaje"`
	Data    Factura `json:"data"`
}

type RespPresupuestos struct {
	Status  int            `json:"status"`
	Mensaje string         `json:"mensaje"`
	Data    []Presupuestos `json:"data"`
}

type ItemsFactura struct {
	Idfact      int     `json:"idfact"`
	Codprod     string  `json:"codprod"`
	Producto    string  `json:"producto"`
	Cant        float32 `json:"cant"`
	Precio      float64 `json:"precio"`
	Subtotal    float64 `json:"subtotal"`
	Descuento   float32 `json:"descuento"`
	Neto        float32 `json:"neto"`
	Descripcion string  `json:"descripcion"`
	Cantmp      float32 `json:"cantmp"`
	Cantpres    float32 `json:"cantpres"`
	Iditempres  float32 `json:"iditempres"`
}

type ItemsPresupuestos struct {
	Idpre       int     `json:"idpre"`
	Codprod     string  `json:"codprod"`
	Producto    string  `json:"producto"`
	Cant        float32 `json:"cant"`
	Precio      float64 `json:"precio"`
	Subtotal    float64 `json:"subtotal"`
	Descuento   float32 `json:"descuento"`
	Neto        float32 `json:"neto"`
	Descripcion string  `json:"descripcion"`
	Cantmp      float32 `json:"cantmp"`
	Iditempres  float32 `json:"iditempres"`
}

type VentasResumen struct {
	Id         int     `json:"id"`
	Idcliente  string  `json:"idcliente"`
	Cliente    string  `json:"cliente"`
	Dirfiscal  string  `json:"dirfiscal"`
	Rif        string  `json:"rif"`
	Persconta  string  `json:"persconta"`
	Tlfconta   string  `json:"tlfconta"`
	Fecha      string  `json:"fecha"`
	Subtotal   float32 `json:"subtotal"`
	Dscto      float32 `json:"dscto"`
	Mototal    float32 `json:"mototal"`
	Deimp      string  `json:"deimp"`
	Tasaimp    float32 `json:"tasaimp"`
	Moimp      float32 `json:"moimp"`
	Moneto     float32 `json:"moneto"`
	Tasadiv    float32 `json:"tasadiv"`
	Monetodiv  float32 `json:"monetodiv"`
	Idvendedor int     `json:"idvendedor"`
	Vendedor   string  `json:"vendedor"`
	Idsesion   int     `json:"idsesion"`
	Idmesa     int     `json:"idmesa"`
	Idmesonero string  `json:"idmesonero"`
	Pagado     float32 `json:"pagado"`
	Porpagar   float32 `json:"porpagar"`
	Cambio     float32 `json:"cambio"`
	Cxcbs      float32 `json:"cxcbs"`
	Cxcdiv     float32 `json:"cxcdiv"`
	Valido     bool    `json:"valido"`
	Esdivisa   bool    `json:"esdivisa"`
}

type ResVentasDia struct {
	Cant           float32 `json:"cant"`
	Monto          float32 `json:"monto"`
	Descuento      float32 `json:"descuento"`
	Subtotal       float32 `json:"subtotal"`
	Subtotaldivisa float32 `json:"subtotaldivisa"`
}

type IdFactura struct {
	Id     string  `json:"id"`
	Pagado float64 `json:"pagado"`
	Pagos  Pagos   `json:"pagos"`
}

type ResVentasProductos struct {
	CodProd        string  `json:"codprod"`
	Producto       string  `json:"deprod"`
	Cantidad       float32 `json:"cantidad"`
	Subtotal       float32 `json:"subtotal"`
	Subtotaldivisa float32 `json:"subtotaldivisa"`
	Cantmp         float32 `json:"cantmp"`
	Cantpres       float32 `json:"cantpres"`
}

type RespProductos struct {
	Status  int                  `json:"status"`
	Mensaje string               `json:"mensaje"`
	Data    []ResVentasProductos `json:"data"`
}

type TopVentas struct {
	Codprod  string  `json:"codprod"`
	Producto string  `json:"producto"`
	Cantidad float32 `json:"cantidad"`
	Venta    float32 `json:"venta"`
}

type RespTopVentas struct {
	Status  int         `json:"status"`
	Mensaje string      `json:"mensaje"`
	Data    []TopVentas `json:"data"`
}

type VentasMensual struct {
	Nro     int     `json:"nro"`
	Mes     string  `json:"mes"`
	Bs      float32 `json:"bs"`
	Divisas float32 `json:"divisas"`
}

type RespVentasMensual struct {
	Status  int             `json:"status"`
	Mensaje string          `json:"mensaje"`
	Data    []VentasMensual `json:"data"`
}
