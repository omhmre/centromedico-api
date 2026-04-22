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

	fmt.Println("--- Todos los Doctores ---")
	rows, err := db.Query("SELECT id, nombres FROM medi001.doctores")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nombres string
		err := rows.Scan(&id, &nombres)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Nombre: %s\n", id, nombres)
	}
}
