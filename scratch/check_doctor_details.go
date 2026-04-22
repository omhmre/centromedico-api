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

	fmt.Println("--- Detalles de columnas medi001.doctores ---")
	rows, err := db.Query("SELECT column_name, data_type, is_nullable, column_default FROM information_schema.columns WHERE table_schema = 'medi001' AND table_name = 'doctores'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name, dataType, isNullable string
		var columnDefault sql.NullString
		err := rows.Scan(&name, &dataType, &isNullable, &columnDefault)
		if err != nil {
			log.Fatal(err)
		}
		def := "NULL"
		if columnDefault.Valid {
			def = columnDefault.String
		}
		fmt.Printf("%s (%s) - Nullable: %s, Default: %s\n", name, dataType, isNullable, def)
	}
}
