package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=aws-1-us-east-2.pooler.supabase.com user=postgres.sczbdxihtitkuesatpng password=Omhmre2025* dbname=postgres port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}
	fmt.Println("Conectado exitosamente a PostgreSQL!")

	// Check seguridad.parametros
	fmt.Println("Listando parametros de seguridad...")
	query := `SELECT id, parametro, descripcion, valor FROM seguridad.parametros`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error in Query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var parametro, descripcion, valor sql.NullString
		err = rows.Scan(&id, &parametro, &descripcion, &valor)
		if err != nil {
			log.Fatalf("Scan error: %v", err)
		}
		fmt.Printf("PARAM: id=%d, name=%s, desc=%s, value=%s\n", id, parametro.String, descripcion.String, valor.String)
	}
}
