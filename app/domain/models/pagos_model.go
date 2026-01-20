package models

import "time"

type Pagos struct {
	Id        int64       `json:"id"`
	Idcliente string      `json:"idcliente"`
	Fecha     string      `json:"fecha"`
	Monto     float64     `json:"monto"`
	Dscto     float64     `json:"dscto"`
	Total     float64     `json:"total"`
	Idsesion  int64       `json:"idsesion"`
	Items     []ItemsPago `json:"items"`
}

type DetPago struct {
	Id          int64   `json:"id"`
	Idpago      int64   `json:"idpago"`
	Idinstpago  int64   `json:"idinstpago"`
	Descripcion string  `json:"descripcion"`
	Comenta     string  `json:"comenta"`
	Monto       float64 `json:"monto"`
	Tasa        float64 `json:"tasa"`
	Total       float64 `json:"total"`
	Idfact      int64   `json:"idfact"`
}

type ItemsPago struct {
	Idfactura int     `json:"idfactura"`
	Codpago   int     `json:"codpago"`
	Depago    string  `json:"depago"`
	Cant      float64 `json:"cant"`
	Tasa      float64 `json:"tasa"`
	Subtotal  float64 `json:"subtotal"`
}

type ResumenDetPago struct {
	Idinstpago  int64   `json:"idinstpago"`
	Descripcion string  `json:"descripcion"`
	Cant        int64   `json:"cant"`
	Montos      float64 `json:"montos"`
	Tasa        float64 `json:"tasa"`
	Totales     float64 `json:"totales"`
}

type RespDetPagos struct {
	Status  int       `json:"status"`
	Mensaje string    `json:"mensaje"`
	Data    []DetPago `json:"data"`
}

type Payments struct {
	Id            int64     `json:"id"`
	Appointmentid int64     `json:"appointmentid"`
	Paymentmethod string    `json:"paymentmethod"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Reference     string    `json:"reference"`
	Date          time.Time `json:"date"`
	Status        string    `json:"status"`
	Notes         string    `json:"notes"`
}

type RelPagos struct {
	Doctor_id       int64   `json:"doctor_id"`
	Doctor_name     string  `json:"doctor_name"`
	Cita_id         int64   `json:"cita_id"`
	Paciente_nombre string  `json:"paciente_nombre"`
	Fecha_pago      string  `json:"fecha_pago"`
	Monto_cita      float64 `json:"monto_cita"`
	Pago_doctor     float64 `json:"pago_doctor"`
	Monto_doctor    float64 `json:"monto_doctor"`
	Forma_pago      string  `json:"formas_pago"`
	Saldo           float64 `json:"saldo"`
}
