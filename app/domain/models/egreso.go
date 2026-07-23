package models

import "time"

type Egreso struct {
	ID               int       `json:"id"`
	Fecha            time.Time `json:"fecha"`
	Descripcion      string    `json:"descripcion"`
	Proveedor        string    `json:"proveedor"`
	Monto            float64   `json:"monto"`
	Categoria        string    `json:"categoria"`
	MetodoPago       string    `json:"metodo_pago"`
	Referencia       string    `json:"referencia"`
	UsuarioOperacion string    `json:"usuario_operacion"`
	FechaOperacion   time.Time `json:"fecha_operacion"`
}

type ConfigEgreso struct {
	ID    int    `json:"id"`
	Tipo  string `json:"tipo"` // 'frecuente' o 'metodo'
	Valor string `json:"valor"`
}

