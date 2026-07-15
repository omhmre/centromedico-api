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

	// See if there's any row with entregado = true
	rows, err := db.Query("SELECT id, cedula_paciente, entregado FROM medi001.evaluaciones_sociales ORDER BY updated_at DESC LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Últimas 5 evaluaciones:")
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
