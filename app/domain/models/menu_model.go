package models

type Menu struct {
	Codigo      string  `json:"codigo"`
	Nombre      string  `json:"nombre"`
	Descripcion string  `json:"descripcion"`
	Idclase     int64   `json:"idclase"`
	Precio1     float64 `json:"precio1"`
	Precio2     float64 `json:"precio2"`
	Precio3     float64 `json:"precio3"`
	Cantidad    float64 `json:"cantidad"`
	Preciom1    float64 `json:"preciom1"`
	Preciom2    float64 `json:"preciom2"`
	Preciom3    float64 `json:"preciom3"`
	Dirfoto     string  `json:"dirfoto"`
	Foto        string  `json:"foto"`
}

type RespMenuLista struct {
	Status  int    `json:"status"`
	Mensaje string `json:"mensaje"`
	Items   []Menu `json:"data"`
}

type ClaseMenu struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
	Menu   []Menu `json:"data"`
}

type MenuCompleto struct {
	MiEmpresa   Empresa     `json:"empresa"`
	MiClaseMenu []ClaseMenu `json:"clases"`
}

type RespMenuCompleto struct {
	Status  int          `json:"status"`
	Mensaje string       `json:"mensaje"`
	Data    MenuCompleto `json:"data"`
}

type RespMenu struct {
	Status  int         `json:"status"`
	Mensaje string      `json:"mensaje"`
	Items   []ClaseMenu `json:"data"`
}

type Clase struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
}

type RespClase struct {
	Status  int     `json:"status"`
	Mensaje string  `json:"mensaje"`
	Data    []Clase `json:"data"`
}
