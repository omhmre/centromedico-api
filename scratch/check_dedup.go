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

	rows, err := db.Query("SELECT id, cedula, nombres FROM medi001.pacientes WHERE id IN (140, 619, 186, 576) OR cedula IN ('140', '619', '186', '576')")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var cedula, nombres string
		rows.Scan(&id, &cedula, &nombres)
		fmt.Printf("ID: %v, Cedula: %v, Nombres: %v\n", id, cedula, nombres)
	}
}
