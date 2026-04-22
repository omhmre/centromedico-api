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

	fmt.Println("--- Todas las Citas (Primeras 20) ---")
	rows, err := db.Query("SELECT id, iddoctor, especialidad, montoref, pagado, porcentaje_comision, status FROM medi001.citas ORDER BY inicio DESC LIMIT 20")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, iddoctor int
		var especialidad, status string
		var montoref, pagado, porcentaje_comision float64
		err := rows.Scan(&id, &iddoctor, &especialidad, &montoref, &pagado, &porcentaje_comision, &status)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, DocID: %d, Serv: %s, Ref: %.2f, Pag: %.2f, Com: %.2f%%, Status: %s\n", id, iddoctor, especialidad, montoref, pagado, porcentaje_comision, status)
	}
}
