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

	fmt.Println("--- Últimos 5 especialistas actualizados ---")
	rows, err := db.Query("SELECT id, nombres, titulo, num_mpps, rif FROM medi001.doctores ORDER BY id DESC LIMIT 5")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var nombres, titulo, mpps, rif sql.NullString
		err := rows.Scan(&id, &nombres, &titulo, &mpps, &rif)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Nombre: %s, Titulo: %s, MPPS: %s, RIF: %s\n", 
			id, nombres.String, titulo.String, mpps.String, rif.String)
	}
}
