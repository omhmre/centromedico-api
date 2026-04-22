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

	var nombre string
	err = db.QueryRow("SELECT nombre FROM medi001.personal WHERE UPPER(nombre) LIKE '%INDHIRA NAVARRO%'").Scan(&nombre)
	if err == sql.ErrNoRows {
		fmt.Println("No está en personal")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Encontrado en personal: %s\n", nombre)
	}
}
