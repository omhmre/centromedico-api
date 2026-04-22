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

	fmt.Println("--- Doctores con es_medico: true ---")
	rows, err := db.Query("SELECT id, nombres, titulo, rif FROM medi001.doctores WHERE es_medico = true")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nombres, titulo, rif string
		err := rows.Scan(&id, &nombres, &titulo, &rif)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Nombre: %s, Titulo: %s, RIF: %s\n", id, nombres, titulo, rif)
	}
}
