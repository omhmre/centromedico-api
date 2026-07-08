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

	// Find tables with names containing pagos, cxc, facturas, etc
	q := `
	SELECT table_schema, table_name, column_name 
	FROM information_schema.columns 
	WHERE (table_name LIKE '%pago%' OR table_name LIKE '%cxc%' OR table_name LIKE '%factura%' OR table_name LIKE '%venta%');
	`
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Payment/CXC related tables:")
	for rows.Next() {
		var schema, table, col string
		rows.Scan(&schema, &table, &col)
		if col == "cedula" || col == "id_paciente" || col == "cedula_cliente" || col == "id_cliente" || col == "cliente" || col == "paciente_id" {
			fmt.Printf("- %s.%s (%s)\n", schema, table, col)
		}
	}
}
