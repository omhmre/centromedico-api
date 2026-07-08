package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgresql://postgres.sczbdxihtitkuesatpng:Omhmre2025*@aws-1-us-east-2.pooler.supabase.com:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var cedula sql.NullString
	var id int
	err = db.QueryRow("SELECT id, cedula FROM medi001.pacientes WHERE nombres ILIKE '%dyland%' LIMIT 1").Scan(&id, &cedula)
	fmt.Printf("Dyland in pacientes: ID: %d, Cedula: %v\n", id, cedula.String)

	var evalId int
	var evalCedula string
	err = db.QueryRow("SELECT id, cedula_paciente FROM medi001.evaluaciones_sociales WHERE cedula_paciente = $1", cedula.String).Scan(&evalId, &evalCedula)
	if err != nil {
		fmt.Println("Evaluation not found for this cedula:", err)
	} else {
		fmt.Printf("Evaluation found! ID: %d, Cedula: %s\n", evalId, evalCedula)
	}
}
