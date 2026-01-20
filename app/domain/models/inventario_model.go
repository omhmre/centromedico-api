package models

type Inventario struct {
	Codigo         string             `json:"codigo"`
	Nombre         string             `json:"nombre"`
	Marca          string             `json:"marca"`
	Unidad         string             `json:"unidad"`
	Costo          float64            `json:"costo"`
	Costoa         float64            `json:"costoa"`
	Costopr        float64            `json:"costopr"`
	Precio1        float64            `json:"precio1"`
	Precio2        float64            `json:"precio2"`
	Precio3        float64            `json:"precio3"`
	Cantidad       float32            `json:"cantidad"`
	Enser          int                `json:"enser"`
	Exento         int                `json:"exento"`
	Clasif         int                `json:"clasif"`
	Tipo           int                `json:"tipo"`
	Empaque        string             `json:"empaque"`
	Cantemp        int                `json:"cantemp"`
	Pedido         float64            `json:"pedido"`
	Disponible     float64            `json:"disponible"`
	Preciom1       float64            `json:"preciom1"`
	Preciom2       float64            `json:"preciom2"`
	Preciom3       float64            `json:"preciom3"`
	Costodolar     float64            `json:"costodolar"`
	Dirfoto        string             `json:"dirfoto"`
	Foto           []byte             `json:"foto"`
	Descripcion    string             `json:"descripcion"`
	Codservicio    string             `json:"codservicio"`
	Preciovar      bool               `json:"preciovar"`
	Compuesto      bool               `json:"compuesto"`
	Mateprima      bool               `json:"mateprima"`
	Global         bool               `json:"global"`
	Cantvar        bool               `json:"cantvar"`
	Espresent      bool               `json:"espresent"`
	Items          []ItemsInventario  `json:"items"`
	Presentaciones []PresenInventario `json:"presentaciones"`
}

type RespInventario struct {
	Status  int          `json:"status"`
	Mensaje string       `json:"mensaje"`
	Data    []Inventario `json:"data"`
}

type InventarioNombre struct {
	Codigo   string  `json:"codigo,omitempty"`
	Nombre   string  `json:"nombre,omitempty"`
	Preciom1 float64 `json:"preciom1,omitempty"`
}

type RespInventarioNombre struct {
	Status  int                `json:"status"`
	Mensaje string             `json:"mensaje"`
	Data    []InventarioNombre `json:"data"`
}

type ItemsInventario struct {
	Codinventario string  `json:"codinventario"`
	Coditem       string  `json:"coditem"`
	Nombre        string  `json:"nombre"`
	Cantidad      float32 `json:"cantidad"`
}

type PresenInventario struct {
	Id           int                 `json:"id"`
	Codinv       string              `json:"codinv"`
	Presentacion string              `json:"presentacion"`
	Cantidad     float32             `json:"cantidad"`
	Precio       float32             `json:"precio"`
	Items        []ItemsPresentacion `json:"itemspres"`
}

type ItemsPresentacion struct {
	Id       int     `json:"id"`
	Codpres  int     `json:"codpres"`
	Codinv   string  `json:"codinv"`
	Cantidad float32 `json:"cantidad"`
}

type InventarioFormal struct {
	Codigo   string  `json:"codigo"`
	Nombre   string  `json:"nombre"`
	Unidad   string  `json:"unidad"`
	Cantidad float32 `json:"cantidad"`
	Costo    float64 `json:"costo"`
	Precio1  float64 `json:"precio1"`
	Preciom1 float64 `json:"preciom1"`
}
