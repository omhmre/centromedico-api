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

	fmt.Println("--- Citas con Porcentaje de Comisión ---")
	rows, err := db.Query("SELECT id, iddoctor, especialidad, montoref, pagado, porcentaje_comision FROM medi001.citas WHERE porcentaje_comision > 0 LIMIT 15")
	if err != nil {
		fmt.Println("Error en query con filtro:", err)
		rows, err = db.Query("SELECT id, iddoctor, especialidad, montoref, pagado, porcentaje_comision FROM medi001.citas LIMIT 15")
		if err != nil {
			log.Fatal(err)
		}
	}
	defer rows.Close()

	found := false
	for rows.Next() {
		found = true
		var id, iddoctor int
		var especialidad string
		var montoref, pagado, porcentaje_comision float64
		err := rows.Scan(&id, &iddoctor, &especialidad, &montoref, &pagado, &porcentaje_comision)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, DocID: %d, Serv: %s, Ref: %.2f, Pag: %.2f, Com: %.2f%%\n", id, iddoctor, especialidad, montoref, pagado, porcentaje_comision)
	}
	if !found {
		fmt.Println("No se encontraron citas con comisión mayor a 0.")
	}
}
