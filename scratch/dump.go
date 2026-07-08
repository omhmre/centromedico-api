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

	var id int
	var cedula, nombres string
	err = db.QueryRow("SELECT id, cedula, nombres FROM medi001.pacientes WHERE nombres ILIKE '%dyland%' LIMIT 1").Scan(&id, &cedula, &nombres)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, cedula, nombres)
}
