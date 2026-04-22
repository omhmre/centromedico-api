package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

type Staff struct {
	Nombre          string
	Cedula          string
	FechaNacimiento string
	Especialidad    string
	Correo          string
	Telefono        string
	AnioIngreso     string
	IsDoctor        bool
}

func main() {
	connStr := "postgres://postgres.sczbdxihtitkuesatpng:Omhmre2025*@aws-1-us-east-2.pooler.supabase.com:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	staffList := []Staff{
		{"Rosa Vergara", "14.220.587", "1978-07-28", "Presidente", "rosavergara36@gmail.com", "0426 5869865", "2017", false},
		{"Indhira Navarro", "11.197.486", "1971-12-18", "Psicopedagogia", "navarroindhira@gmail.com", "0416 8993309", "2020", true},
		{"Grisdelia Oropeza", "10.197.765", "1970-07-26", "Psicopedagogia", "grisdeliaoropeza70@gmail.com", "0424 8227007", "2017", true},
		{"Arizai Hernandez", "6.561.266", "1964-07-01", "Psicopedagogia", "arizai0107@gmail.com", "0424 8175150", "2023", true},
		{"Josmely Rojas", "18.939.209", "1989-03-03", "Psicopedagogia", "josmeli.vzla@gmail.com", "0414 8370748", "2023", true},
		{"Karla Medina", "11.924.862", "1975-06-20", "Psicopedagogia", "karlitamn75@gmail.com", "0412 0920058", "2025", true},
		{"Lisnelys Pineda", "26.127.375", "1996-09-30", "Terap Ocupacional", "lisnelyspineda@gmail.com", "0412 8543435", "2025", true},
		{"Maria Romero", "20.905.037", "1991-09-18", "Terap Ocupacional", "to.mariagabrielaromero@gmail.com", "0414 903 1245", "2023", true},
		{"Maria Noriega", "15.815.576", "1983-12-14", "Terap Lenguaje", "mariel241207@gmail.com", "0416 5796728", "2023", true},
		{"Jennifer Marcano", "14.803.767", "1981-06-19", "Pediatra", "jaleiferna@gmail.com", "0412 9461678", "2023", true},
		{"Michel Mendoza", "12.958.374", "1976-07-07", "Pediatra", "michelle.mendoza1516@gmail.com", "0416 6950941", "2022", true},
		{"Jesus Montaner", "14.055.710", "1978-10-27", "Neurologo", "castlevania1027@gmail.com", "0412 0969709", "2018", true},
		{"Yris Caraballo", "8.387.903", "1962-07-06", "Nutricionista", "yrismgcaraballo@hotmail.es", "0414 1886922", "2024", true},
		{"Alexandra Alcala", "6.365.682", "1963-03-01", "Trabajo Social", "alexandravictoria63@gmail.com", "0412 7160589", "2024", true},
		{"Maria Adrian", "13.404.935", "1977-12-30", "Psicologo", "madrianantunez@gmail.com", "0424 7450926", "2023", true},
		{"Miguel Jimenez", "4.910.280", "1956-04-01", "Psicologo", "miguelljif@hotmail.com", "0424 8725346", "2022", true},
		{"Ynes Lugo", "15.592.561", "1980-05-10", "Terapia de Conducta", "lvalerio.ynes@gmail.com", "0424 6606362", "2026", true},
		{"David Lovera", "13.126.206", "1976-10-27", "Arte Terapia", "dajaloap2@gmail.com", "0412 5826486", "2026", true},
		{"Francisco Lopez", "9.424.381", "1966-10-04", "Serv. Generales", "ismaellopez1510@gmail.com", "0412 5906539", "2017", false},
		{"Ismael Lopez", "29.655.021", "2002-10-15", "Asist. Legal", "ismaellopez1510@gmail.com", "0412 0251785", "2017", false},
		{"Karelys Coraspe", "22.629.265", "1990-12-02", "Asist. Presidencia", "karelyscoraspe2@gmail.com", "0412 586 5204", "2017", false},
		{"Patricia Urbano", "5.305.179", "1961-10-11", "Asist. Administrativo", "ttu6461@gmail.com", "0414 7896519", "2021", false},
		{"Dannymar Ramos", "13.848.052", "1979-01-14", "Protocolo", "dannymarramos13@gmail.com", "0414 188 8136", "2017", false},
		{"Trina Gonzales", "12165132", "1974-03-23", "Protocolo", "trinamgp.tg@gmail.com", "0412 3507364", "2017", false},
		{"Sofia Lopez", "32.475.272", "2008-10-26", "R.R.P.P", "spfilopillll2521@gmail.com", "0412 5096775", "2017", false},
		{"Peggy Colina", "14.024.232", "1978-02-23", "Asist. Medicos", "neumary22@gmail.com", "0412 6864016", "2026", false},
		{"Hiaura Brito", "32.729.075", "2007-01-10", "Redes", "ilaubc@gmail.com", "0424 899 1916", "2026", false},
	}

	for _, s := range staffList {
		if s.IsDoctor {
			// Update or Insert in doctores
			var id int
			err := db.QueryRow("SELECT id FROM medi001.doctores WHERE UPPER(nombres) LIKE $1", "%"+strings.ToUpper(s.Nombre)+"%").Scan(&id)
			if err == sql.ErrNoRows {
				// Insert
				fmt.Printf("Insertando Doctor: %s\n", s.Nombre)
				_, err = db.Exec(`INSERT INTO medi001.doctores (nombres, cedula, fecha_nacimiento, fecha_ingreso, correo, tlf, espec, servicios) 
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
					s.Nombre, s.Cedula, s.FechaNacimiento, s.AnioIngreso+"-01-01", s.Correo, s.Telefono, s.Especialidad, "[]")
			} else if err == nil {
				// Update
				fmt.Printf("Actualizando Doctor: %s (ID %d)\n", s.Nombre, id)
				_, err = db.Exec(`UPDATE medi001.doctores SET cedula = $1, fecha_nacimiento = $2, fecha_ingreso = $3, correo = $4, tlf = $5 WHERE id = $6`,
					s.Cedula, s.FechaNacimiento, s.AnioIngreso+"-01-01", s.Correo, s.Telefono, id)
			}
			if err != nil {
				log.Printf("Error con doctor %s: %v", s.Nombre, err)
			}
		} else {
			// Update or Insert in personal
			var id int
			err := db.QueryRow("SELECT id FROM medi001.personal WHERE UPPER(nombre) LIKE $1", "%"+strings.ToUpper(s.Nombre)+"%").Scan(&id)
			if err == sql.ErrNoRows {
				// Insert
				fmt.Printf("Insertando Personal: %s\n", s.Nombre)
				_, err = db.Exec(`INSERT INTO medi001.personal (nombre, cedula, fecha_nacimiento, fecha_ingreso, correo, telefono, cargo) 
					VALUES ($1, $2, $3, $4, $5, $6, $7)`,
					s.Nombre, s.Cedula, s.FechaNacimiento, s.AnioIngreso+"-01-01", s.Correo, s.Telefono, s.Especialidad)
			} else if err == nil {
				// Update
				fmt.Printf("Actualizando Personal: %s (ID %d)\n", s.Nombre, id)
				_, err = db.Exec(`UPDATE medi001.personal SET cedula = $1, fecha_nacimiento = $2, fecha_ingreso = $3, correo = $4, telefono = $5, cargo = $6 WHERE id = $7`,
					s.Cedula, s.FechaNacimiento, s.AnioIngreso+"-01-01", s.Correo, s.Telefono, s.Especialidad, id)
			}
			if err != nil {
				log.Printf("Error con personal %s: %v", s.Nombre, err)
			}
		}
	}
	fmt.Println("Migración completada.")
}
