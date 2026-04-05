package models

import "fmt"

type DoctoresModel struct {
	Id        int     `json:"id"`
	Nombres   string  `json:"nombres"`
	Espec     string  `json:"espec"`
	Dir       string  `json:"dir"`
	Tlf       string  `json:"tlf"`
	Correo    string  `json:"correo"`
	Whatsapp  string  `json:"whatsapp"`
	Instagram    string  `json:"instagram"`
	Tasapago    float32 `json:"tasapago"`
	DaysOfWeek  []int   `json:"daysOfWeek"`
	StartTime   string  `json:"startTime"`
	EndTime     string  `json:"endTime"`
	SlotDuration int    `json:"slotDuration"`
}

// Add String method for DoctoresModel
func (d DoctoresModel) String() string {
	return fmt.Sprintf("ID: %d, Nombres: %s, Especialidad: %s, Teléfono: %s, Correo: %s",
		d.Id, d.Nombres, d.Espec, d.Tlf, d.Correo)
}
