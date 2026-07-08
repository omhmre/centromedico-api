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

	// Check if 619 or 576 still exist
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM medi001.pacientes WHERE cedula IN ('619', '576')").Scan(&count)
	if err != nil { log.Fatal(err) }
	fmt.Printf("Number of patients with cedula 619 or 576: %d\n", count)

	// Check citations for 140
	err = db.QueryRow("SELECT COUNT(*) FROM medi001.citas WHERE cedula = '140'").Scan(&count)
	if err != nil { log.Fatal(err) }
	fmt.Printf("Number of citas for cedula 140: %d\n", count)

	// Check evaluations for 140
	err = db.QueryRow("SELECT COUNT(*) FROM medi001.evaluaciones_sociales WHERE cedula_paciente = '140'").Scan(&count)
	if err != nil { log.Fatal(err) }
	fmt.Printf("Number of evaluaciones_sociales for cedula 140: %d\n", count)
}
