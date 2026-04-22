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

	rows, err := db.Query("SELECT column_name, character_maximum_length FROM information_schema.columns WHERE table_schema = 'medi001' AND table_name = 'doctores' AND data_type = 'character varying'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var length sql.NullInt64
		err := rows.Scan(&name, &length)
		if err != nil {
			log.Fatal(err)
		}
		lenStr := "unlimited"
		if length.Valid {
			lenStr = fmt.Sprintf("%d", length.Int64)
		}
		fmt.Printf("%s: %s\n", name, lenStr)
	}
}
