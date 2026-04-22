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

	fmt.Println("--- Verificando nulos en es_medico ---")
	rows, err := db.Query("SELECT id, nombres, es_medico FROM medi001.doctores WHERE es_medico IS NULL")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nombres string
		var esMedico sql.NullBool
		err := rows.Scan(&id, &nombres, &esMedico)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Nombre: %s, EsMedico IS NULL\n", id, nombres)
	}
}
