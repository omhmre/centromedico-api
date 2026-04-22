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

	fmt.Println("--- Últimos 5 especialistas actualizados (con servicios) ---")
	rows, err := db.Query("SELECT id, nombres, es_medico, servicios FROM medi001.doctores ORDER BY id DESC LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nombres string
		var esMedico bool
		var servicios sql.NullString
		err := rows.Scan(&id, &nombres, &esMedico, &servicios)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Nombre: %s, EsMedico: %v, Servicios: %s\n", 
			id, nombres, esMedico, servicios.String)
	}
}
