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

	rows, err := db.Query("SELECT id, cedula_paciente, entregado FROM medi001.evaluaciones_sociales")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var cedula string
		var entregado sql.NullBool
		if err := rows.Scan(&id, &cedula, &entregado); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id=%d cedula=%s entregado=%v\n", id, cedula, entregado.Bool)
	}
}
