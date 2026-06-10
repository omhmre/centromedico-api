package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DoctorService struct {
	Nombre   string  `json:"nombre"`
	Precio   float64 `json:"precio"`
	Comision float64 `json:"comision"`
}

type DoctoresModel struct {
	Id             int             `json:"id"`
	Nombres        string          `json:"nombres"`
	Servicios      []DoctorService `json:"servicios"`
	Dir            string          `json:"dir" default:""`
	Correo         string          `json:"correo" default:""`
	Whatsapp       string          `json:"whatsapp" default:""`
	Instagram      string          `json:"instagram" default:""`
	DaysOfWeek     []int           `json:"days_of_week"`
	StartTime      string          `json:"start_time" default:"08:00"`
	EndTime        string          `json:"end_time" default:"18:00"`
	SlotDuration   int             `json:"slot_duration" default:"45"`
	MontoCita      float64         `json:"monto_cita" default:"0.0"`
	EsMedico       bool            `json:"es_medico"`
	Titulo         string          `json:"titulo" default:""`
	TituloAcademico string         `json:"titulo_academico" default:""`
	NumMPPS        string          `json:"num_mpps" default:""`
	NumCM          string          `json:"num_cm" default:""`
	Rif            string          `json:"rif" default:""`
	Cedula         string          `json:"cedula" default:""`
	FechaNacimiento string         `json:"fecha_nacimiento" default:""`
	FechaIngreso    string         `json:"fecha_ingreso" default:""`
	Sueldo          float64        `json:"sueldo" default:"0.0"`
	FrecuenciaPago  string         `json:"frecuencia_pago" default:"Mensual"`
	Espec           string         `json:"espec" default:""`
	Tlf             string         `json:"tlf" default:""`
}

func main() {
	connStr := "postgres://postgres.sczbdxihtitkuesatpng:Omhmre2025*@aws-1-us-east-2.pooler.supabase.com:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Let's try to update ID 23 (Josmely Rojas) to be a doctor
	// We'll use the exact same logic as in db.go
	sqlUpdDoctores := `UPDATE medi001.doctores SET nombres = $2, servicios = $3, dir = $4, correo = $5, whatsapp = $6, instagram = $7, days_of_week = $8, start_time = $9, end_time = $10, slot_duration = $11, monto_cita = $12, es_medico = $13, titulo = $14, titulo_academico = $15, num_mpps = $16, num_cm = $17, rif = $18, cedula = $19, fecha_nacimiento = $20, fecha_ingreso = $21, sueldo = $22, frecuencia_pago = $23, espec = $24, tlf = $25 WHERE id = $1`
	
	id := 23
	nombres := "Josmely Rojas"
	serviciosJson := "[]"
	dir := ""
	correo := ""
	whatsapp := ""
	instagram := ""
	daysOfWeekJson := "[1,2,3,4,5]"
	startTime := "08:00"
	endTime := "18:00"
	slotDuration := 45
	montoCita := 0.0
	esMedico := true
	titulo := "LICENCIADA"
	tituloAcademico := "PSICOPEDAGOGA"
	numMPPS := "123"
	numCM := "456"
	rif := "V123456789"
	cedula := "12345678"
	fechaNacimiento := "1990-01-01"
	fechaIngreso := "2024-01-01"
	sueldo := 500.0
	frecuenciaPago := "Quincenal"
	espec := "Psicopedagogia EDITADA"
	tlf := "04120000000"

	resp, err := db.Exec(sqlUpdDoctores, 
		id,
		nombres,
		serviciosJson,
		dir,
		correo,
		whatsapp,
		instagram,
		daysOfWeekJson,
		startTime,
		endTime,
		slotDuration,
		montoCita,
		esMedico,
		titulo,
		tituloAcademico,
		numMPPS,
		numCM,
		rif,
		cedula,
		fechaNacimiento,
		fechaIngreso,
		sueldo,
		frecuenciaPago,
		espec,
		tlf,
	)

	if err != nil {
		log.Fatal(err)
	}

	rows, _ := resp.RowsAffected()
	fmt.Printf("Rows affected: %d\n", rows)
}
