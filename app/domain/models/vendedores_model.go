package models

type Vendedores struct {
	Success bool       `json:"success"`
	Mensaje string     `json:"mensaje"`
	Datos   []Vendedor `json:"datos"`
}

type Vendedor struct {
	Id        int    `json:"id"`
	Cedula    string `json:"cedula"`
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
	Correo    string `json:"correo"`
	Codvend   string `json:"codvend"`
}

type VendedoresCxc struct {
	Id        int            `json:"id"`
	Cedula    string         `json:"cedula"`
	Nombre    string         `json:"nombre"`
	Direccion string         `json:"direccion"`
	Telefono  string         `json:"telefono"`
	Correo    string         `json:"correo"`
	Codvend   string         `json:"codvend"`
	Totcxc    float64        `json:"totcxc"`
	Totcxcbs  float64        `json:"totcxcbs"`
	Totcxcdiv float64        `json:"totcxcdiv"`
	Items     []ItemsCxcVend `json:"items"`
}

type ItemsCxcVend struct {
	Nombre string  `json:"nombre"`
	Cxcbs  float64 `json:"cxcbs"`
	Cxcdiv float64 `json:"cxcdiv"`
}
