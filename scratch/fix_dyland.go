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

	// Find the evaluation I inserted for Dyland Gonzalez, which has "Jose Joaquin de Leon" in grado_escolar
	var id int
	err = db.QueryRow("SELECT id FROM medi001.evaluaciones_sociales WHERE grado_escolar ILIKE '%Jose Joaquin%'").Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	
	// Update it to cedula 619
	_, err = db.Exec("UPDATE medi001.evaluaciones_sociales SET cedula_paciente = '619' WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Updated evaluation %v to cedula 619\n", id)
}
