package models

import "encoding/json"

type Clientes []Cliente

func UnmarshalClientes(data []byte) (Clientes, error) {
	var r Clientes
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Clientes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Cliente struct {
	Id        string             `json:"id"`
	Tipo      string             `json:"tipo"`
	Nombre    string             `json:"nombre"`
	Rif       string             `json:"rif"`
	Dirfiscal string             `json:"dirfiscal"`
	Ciudad    string             `json:"ciudad"`
	Estado    string             `json:"estado"`
	Telf      string             `json:"telf"`
	Correo    string             `json:"correo"`
	Twitter   string             `json:"twitter"`
	Facebook  string             `json:"facebook"`
	Whatsapp  string             `json:"whatsapp"`
	Instagram string             `json:"instagram"`
	Status    string             `json:"status"`
	Clasif    string             `json:"clasif"`
	Dscto     float64            `json:"dscto"`
	Cred      float64            `json:"cred"`
	Diascr    int64              `json:"diascr"`
	Cxcbs     float64            `json:"cxcbs"`
	Persconta string             `json:"persconta"`
	Tlfconta  string             `json:"tlfconta"`
	Codvend   string             `json:"codvend"`
	Cxcdiv    float64            `json:"cxcdiv"`
	Items     []ItemsCxcClientes `json:"items"`
}

type RespClientes struct {
	Status  int       `json:"status"`
	Mensaje string    `json:"mensaje"`
	Data    []Cliente `json:"data"`
}

type ItemsCxcClientes struct {
	Id         int     `json:"id"`
	Idfact     int     `json:"idfact"`
	Fecha      string  `json:"fecha"`
	Montobs    float64 `json:"montobs"`
	Cobradobs  float64 `json:"cobradobs"`
	Saldobs    float64 `json:"saldobs"`
	Montodiv   float64 `json:"montodiv"`
	Cobradodiv float64 `json:"cobradodiv"`
	Saldodiv   float64 `json:"saldodiv"`
}
