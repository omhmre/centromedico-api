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

	// Get foreign keys referencing pacientes table
	q := `
	SELECT
		tc.table_schema, 
		tc.table_name, 
		kcu.column_name,
		ccu.table_name AS foreign_table_name,
		ccu.column_name AS foreign_column_name 
	FROM 
		information_schema.table_constraints AS tc 
		JOIN information_schema.key_column_usage AS kcu
		  ON tc.constraint_name = kcu.constraint_name
		  AND tc.table_schema = kcu.table_schema
		JOIN information_schema.constraint_column_usage AS ccu
		  ON ccu.constraint_name = tc.constraint_name
		  AND ccu.table_schema = tc.table_schema
	WHERE tc.constraint_type = 'FOREIGN KEY' AND ccu.table_name='pacientes';
	`
	
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Tables referencing pacientes:")
	for rows.Next() {
		var schema, table, col, refTable, refCol string
		rows.Scan(&schema, &table, &col, &refTable, &refCol)
		fmt.Printf("- %s.%s (%s) -> %s(%s)\n", schema, table, col, refTable, refCol)
	}
	
	// Also look for cedula references since some tables might use cedula directly instead of a formal foreign key
	q2 := `
	SELECT table_schema, table_name, column_name 
	FROM information_schema.columns 
	WHERE column_name LIKE '%cedula%' OR column_name LIKE '%paciente%';
	`
	rows2, err := db.Query(q2)
	if err != nil {
		log.Fatal(err)
	}
	defer rows2.Close()
	fmt.Println("\nColumns that might reference cedula or paciente:")
	for rows2.Next() {
		var schema, table, col string
		rows2.Scan(&schema, &table, &col)
		fmt.Printf("- %s.%s (%s)\n", schema, table, col)
	}
}
