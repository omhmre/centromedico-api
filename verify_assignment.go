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

	// 1. Buscar a Peggy
	fmt.Println("=== BUSCANDO USUARIO PEGGY ===")
	var peggyID int
	var peggyNombre string
	var peggyTipo int
	err = db.QueryRow("SELECT id, nombre, idtipouser FROM seguridad.usuarios WHERE nombre ILIKE '%peggy%'").Scan(&peggyID, &peggyNombre, &peggyTipo)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Usuario Peggy no encontrado.")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Encontrado: ID=%d, Nombre=%s, TipoUsuario=%d\n", peggyID, peggyNombre, peggyTipo)
	}

	// 2. Buscar especialistas disponibles
	fmt.Println("\n=== ALGUNOS ESPECIALISTAS EN LA BD ===")
	rows, err := db.Query("SELECT id, nombres, espec, activo FROM medi001.doctores")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var docID int
		var name, espec string
		var activo bool
		if err := rows.Scan(&docID, &name, &espec, &activo); err != nil {
			log.Fatal(err)
		}
		if name == "Michelle" || name == "Borregales" || name == "Michelle Borregales" || docID == 4 || docID == 5 || docID == 6 {
			fmt.Printf("Doctor ID: %d | Nombres: %s | Especialidad: %s | Activo: %v\n", docID, name, espec, activo)
		}
	}

	// 3. Buscar asignaciones de Peggy
	if peggyID > 0 {
		fmt.Printf("\n=== ASIGNACIONES ACTUALES PARA PEGGY (ID: %d) ===\n", peggyID)
		rowsAssigned, err := db.Query(`
			SELECT ud.id_doctor, d.nombres, d.espec, d.activo
			FROM seguridad.usuario_doctor ud
			INNER JOIN medi001.doctores d ON ud.id_doctor = d.id
			WHERE ud.id_usuario = $1`, peggyID)
		if err != nil {
			log.Fatal(err)
		}
		defer rowsAssigned.Close()
		count := 0
		for rowsAssigned.Next() {
			var dID int
			var name, espec string
			var activo bool
			if err := rowsAssigned.Scan(&dID, &name, &espec, &activo); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Asignado -> Doctor ID: %d | Nombres: %s | Especialidad: %s | Activo: %v\n", dID, name, espec, activo)
			count++
		}
		if count == 0 {
			fmt.Println("Peggy no tiene ningún especialista asignado en este momento.")
		}
	}
}
