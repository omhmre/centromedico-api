package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Connection string from env or hardcoded from local.env
	connStr := "postgresql://postgres.sczbdxihtitkuesatpng:Omhmre2025*@aws-1-us-east-2.pooler.supabase.com:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 1. Query for GEMINI_API_KEY in parametros
	fmt.Println("=== SEARCHING PARAMETROS ===")
	rows, err := db.Query("SELECT parametro, valores FROM seguridad.parametros WHERE parametro = 'GEMINI_API_KEY'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var param string
		var valor string
		if err := rows.Scan(&param, &valor); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Parametro: %s | Valor: %s\n", param, valor)
	}

	// 2. Search all doctores matching "jesus montaner"
	fmt.Println("\n=== SEARCHING DOCTORES ===")
	rowsDoc, err := db.Query("SELECT id, Nombres, espec, days_of_week, start_time, end_time, slot_duration FROM medi001.doctores WHERE id = 11")
	if err != nil {
		log.Fatal(err)
	}
	defer rowsDoc.Close()
	for rowsDoc.Next() {
		var id int
		var nombres string
		var espec string
		// Let's use sql.NullString/interface{} to avoid scan errors if null
		var daysRaw interface{}
		var startRaw sql.NullString
		var endRaw sql.NullString
		var slotRaw sql.NullInt32
		if err := rowsDoc.Scan(&id, &nombres, &espec, &daysRaw, &startRaw, &endRaw, &slotRaw); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d | Nombres: %s | Especialidad: %s | Days: %v | Start: %s | End: %s | Slot: %d\n", id, nombres, espec, daysRaw, startRaw.String, endRaw.String, slotRaw.Int32)
	}
}
