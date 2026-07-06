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

	query := `
	ALTER TABLE medi001.evaluaciones_sociales 
	ADD COLUMN IF NOT EXISTS entregado BOOLEAN DEFAULT FALSE,
	ADD COLUMN IF NOT EXISTS unlocked_by INT,
	ADD COLUMN IF NOT EXISTS unlocked_at TIMESTAMP,
	ADD COLUMN IF NOT EXISTS unlock_reason TEXT;
	`
	fmt.Println("Running ALTER TABLE...")
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Error altering table: %v", err)
	}
	fmt.Println("Success!")
}
