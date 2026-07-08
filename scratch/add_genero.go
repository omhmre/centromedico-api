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

	_, err = db.Exec(`ALTER TABLE medi001.pacientes ADD COLUMN genero VARCHAR(50);`)
	if err != nil {
		fmt.Println("Error modifying table or column already exists:", err)
	} else {
		fmt.Println("Successfully added 'genero' column to medi001.pacientes")
	}
}
