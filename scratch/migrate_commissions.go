package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DoctorService struct {
	Nombre   string  `json:"nombre"`
	Precio   float64 `json:"precio"`
	Comision float64 `json:"comision"`
}

func main() {
	connStr := "postgres://postgres.sczbdxihtitkuesatpng:Omhmre2025*@aws-1-us-east-2.pooler.supabase.com:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("--- Iniciando Migración de Honorarios (con progreso) ---")

	rows, err := db.Query("SELECT id, servicios FROM medi001.doctores")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	doctorServices := make(map[int][]DoctorService)
	for rows.Next() {
		var id int
		var serviciosJSON string
		if err := rows.Scan(&id, &serviciosJSON); err != nil {
			log.Fatal(err)
		}
		var services []DoctorService
		json.Unmarshal([]byte(serviciosJSON), &services)
		doctorServices[id] = services
	}

	citaRows, err := db.Query("SELECT id, iddoctor, especialidad, porcentaje_comision FROM medi001.citas WHERE porcentaje_comision = 0")
	if err != nil {
		log.Fatal(err)
	}
	defer citaRows.Close()

	count := 0
	processed := 0
	for citaRows.Next() {
		processed++
		var id, idDoctor int
		var especialidad string
		var comision float64
		if err := citaRows.Scan(&id, &idDoctor, &especialidad, &comision); err != nil {
			log.Fatal(err)
		}

		services, ok := doctorServices[idDoctor]
		if !ok || len(services) == 0 {
			continue
		}

		newComision := 0.0
		newEspecialidad := especialidad
		found := false

		for _, s := range services {
			if s.Nombre == especialidad {
				newComision = s.Comision
				found = true
				break
			}
		}

		if !found || especialidad == "Consulta" {
			newComision = services[0].Comision
			newEspecialidad = services[0].Nombre
		}

		if newComision > 0 {
			_, err := db.Exec("UPDATE medi001.citas SET porcentaje_comision = $1, especialidad = $2 WHERE id = $3", newComision, newEspecialidad, id)
			if err != nil {
				fmt.Printf("Error actualizando cita %d: %v\n", id, err)
			} else {
				count++
			}
		}

		if processed % 100 == 0 {
			fmt.Printf("Procesadas: %d, Actualizadas: %d\n", processed, count)
		}
	}

	fmt.Printf("--- Migración completada. %d citas actualizadas de %d procesadas ---\n", count, processed)
}
