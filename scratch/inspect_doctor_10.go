package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres.sczbdxihtitkuesatpng:Omhmre2025*@aws-1-us-east-2.pooler.supabase.com:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("--- Detalle Doctor ID 10 ---")
	var id int
	var nombres, titulo, academico, mpps, cm, rif string
	var esMedico bool
	err = db.QueryRow("SELECT id, nombres, es_medico, titulo, titulo_academico, num_mpps, num_cm, rif FROM medi001.doctores WHERE id = 10").Scan(&id, &nombres, &esMedico, &titulo, &academico, &mpps, &cm, &rif)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %d\nNombre: %s\nEsMedico: %v\nTitulo: %s\nAcademico: %s\nMPPS: %s\nCM: %s\nRIF: %s\n", 
		id, nombres, esMedico, titulo, academico, mpps, cm, rif)
}
