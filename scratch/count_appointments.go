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

	var total int
	err = db.QueryRow("SELECT count(*) FROM medi001.citas").Scan(&total)
	if err != nil {
		log.Fatal(err)
	}
	
	var needUpdate int
	err = db.QueryRow("SELECT count(*) FROM medi001.citas WHERE porcentaje_comision = 0").Scan(&needUpdate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total Citas: %d, Necesitan actualización: %d\n", total, needUpdate)
}
