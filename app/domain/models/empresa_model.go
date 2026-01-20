package models

type Empresa struct {
	ID        int64   `json:"id"`
	Rif       string  `json:"rif"`
	Rasocial  string  `json:"rasocial"`
	Dirfisc   string  `json:"dirfisc"`
	Ciudad    string  `json:"ciudad"`
	Estado    string  `json:"estado"`
	Telf      string  `json:"telf"`
	Logo      []byte  `json:"logo"`
	Comercial string  `json:"comercial"`
	Slogan    string  `json:"slogan"`
	Iva       float64 `json:"iva"`
	Correo    string  `json:"correo"`
	Instagram string  `json:"instagram"`
	Whatsapp  string  `json:"whatsapp"`
}

type RespEmpre struct {
	Status  int       `json:"status"`
	Mensaje string    `json:"mensaje"`
	Data    []Empresa `json:"data"`
}
