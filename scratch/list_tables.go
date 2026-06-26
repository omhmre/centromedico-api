package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "postgres://postgres.sczbdxihtitkuesatpng:Omhmre2025*@aws-1-us-east-2.pooler.supabase.com:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error al conectar: %v", err)
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT id_usuario, id_doctor 
		FROM seguridad.usuario_doctor;
	`)
	if err != nil {
		log.Fatalf("Error al consultar asignaciones: %v", err)
	}
	defer rows.Close()

	fmt.Println("Asignaciones en seguridad.usuario_doctor:")
	count := 0
	for rows.Next() {
		var uID, dID int
		if err := rows.Scan(&uID, &dID); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("- Usuario ID: %d, Doctor ID: %d\n", uID, dID)
		count++
	}
	fmt.Printf("Total registros: %d\n", count)
}
