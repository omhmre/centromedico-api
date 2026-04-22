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

	// Intentar actualizar el primer doctor
	var id int
	err = db.QueryRow("SELECT id FROM medi001.doctores LIMIT 1").Scan(&id)
	if err != nil {
		log.Fatal("No hay doctores para probar:", err)
	}

	fmt.Printf("Probando actualización para doctor ID: %d\n", id)

	res, err := db.Exec(`UPDATE medi001.doctores SET 
		titulo = 'Prueba', 
		titulo_academico = 'Prueba Académica', 
		num_mpps = '123', 
		num_cm = '456', 
		rif = 'V123' 
		WHERE id = $1`, id)

	if err != nil {
		log.Fatal("Error en UPDATE:", err)
	}

	affected, _ := res.RowsAffected()
	fmt.Printf("Filas afectadas: %d\n", affected)

	// Verificar si se guardó
	var t, ta, mpps, cm, rif string
	err = db.QueryRow("SELECT titulo, titulo_academico, num_mpps, num_cm, rif FROM medi001.doctores WHERE id = $1", id).Scan(&t, &ta, &mpps, &cm, &rif)
	if err != nil {
		log.Fatal("Error en SELECT:", err)
	}

	fmt.Printf("Valores guardados: Titulo=%s, Academico=%s, MPPS=%s, CM=%s, RIF=%s\n", t, ta, mpps, cm, rif)
}
