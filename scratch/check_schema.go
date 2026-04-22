//go:build ignore
package main

import (
	"fmt"
	"log"
	"omhmre.com/centromedico/app/domain/database"
)

func main() {
	database.FetchVars()
	db := &database.DB{}
	if err := db.Open(); err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}
	defer db.Close()

	rows, err := db.Conn.Query("SELECT column_name FROM information_schema.columns WHERE table_schema = 'medi001' AND table_name = 'informe_medico'")
	if err != nil {
		log.Fatalf("Error querying schema: %v", err)
	}
	defer rows.Close()

	fmt.Println("Columns in medi001.informe_medico:")
	for rows.Next() {
		var columnName string
		if err := rows.Scan(&columnName); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("- %s\n", columnName)
	}
}
