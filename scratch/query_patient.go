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

	rows, err := db.Query("SELECT cedula, nombres FROM medi001.pacientes WHERE nombres ILIKE '%Placencia%'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var cedula, nombres string
		if err := rows.Scan(&cedula, &nombres); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Encontrado: %s - %s\n", cedula, nombres)
	}
}
