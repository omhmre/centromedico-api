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

	fmt.Println("--- Personal Registrado ---")
	rows, err := db.Query("SELECT id, nombre, cargo FROM medi001.personal")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nombre, cargo string
		err := rows.Scan(&id, &nombre, &cargo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Nombre: %s, Cargo: %s\n", id, nombre, cargo)
	}
}
