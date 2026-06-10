package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres.sczbdxihtitkuesatpng:Omhmre2025*@aws-1-us-east-2.pooler.supabase.com:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	migration := `
	-- Migration: Add retirement fields to specialists
	ALTER TABLE medi001.doctores ADD COLUMN IF NOT EXISTS activo BOOLEAN DEFAULT TRUE;
	ALTER TABLE medi001.doctores ADD COLUMN IF NOT EXISTS fecha_retiro VARCHAR(20) DEFAULT '';
	ALTER TABLE medi001.doctores ADD COLUMN IF NOT EXISTS motivo_retiro TEXT DEFAULT '';

	-- Update existing records to be active
	UPDATE medi001.doctores SET activo = TRUE WHERE activo IS NULL;
	`

	fmt.Println("Applying migration to medi001.doctores...")
	_, err = db.Exec(migration)
	if err != nil {
		log.Fatalf("Error applying migration: %v", err)
	}

	fmt.Println("Migration applied successfully!")
}
