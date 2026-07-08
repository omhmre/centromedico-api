package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgresql://postgres.sczbdxihtitkuesatpng:Omhmre2025*@aws-1-us-east-2.pooler.supabase.com:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 1. Search for Dyland
	var cedula string
	var patientId int
	err = db.QueryRow("SELECT id, cedula FROM medi001.pacientes WHERE nombres ILIKE '%dyland%' LIMIT 1").Scan(&patientId, &cedula)
	
	if err == sql.ErrNoRows {
		fmt.Println("Dyland not found in pacientes. Inserting...")
		// Insert patient
		fenac, _ := time.Parse("2006-01-02", "2016-09-09") // 09-09-16
		err = db.QueryRow(`
			INSERT INTO medi001.pacientes (nombres, fenac, representante, whatsapp, direccion, correo, cedula)
			VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
		`, "Dyland Alexander González Rivas", fenac, "Anyelis Rivas", "0424-8591995", "Las Guevaras, Calle San Rafael, Casa 1", "anyelysrivas@gmail.com", "99999").Scan(&patientId)
		if err != nil {
			log.Fatal("Failed to insert patient:", err)
		}
		// update cedula to be string of patientId if we don't have real cedula
		cedula = fmt.Sprintf("%d", patientId)
		db.Exec("UPDATE medi001.pacientes SET cedula = $1 WHERE id = $2", cedula, patientId)
		fmt.Printf("Inserted Dyland with ID %d and Cedula %s\n", patientId, cedula)
	} else if err != nil {
		log.Fatal("Error searching patient:", err)
	} else {
		fmt.Printf("Found Dyland! ID: %d, Cedula: %s\n", patientId, cedula)
	}

	// 2. Insert Social Evaluation
	query := `
	INSERT INTO medi001.evaluaciones_sociales (
		cedula_paciente, id_especialista, lugar_nacimiento, escolaridad, grado_escolar, referido_por,
		madre_nombre, madre_edad, madre_ci, madre_telefono, madre_ocupacion, madre_correo, madre_direccion,
		padre_nombre, padre_edad, padre_ci, padre_telefono, padre_ocupacion, padre_direccion,
		antecedentes_desarrollo, aspecto_salud, grupo_familiar, vivienda_entorno, situacion_economica,
		diagnostico_social, conclusion, plan_accion, entregado
	) VALUES (
		$1, 2, 'Porlamar, 09-09-16', '4a', '2do E.B. Jose Joaquin de Leon', 'representante de FUNDAFI',
		'Anyelis Rivas', '26a', 'V-26.243.387', '0424-8591995', 'Analista de Inventario en Rattan', 'anyelysrivas@gmail.com', 'Las Guevaras, Calle San Rafael, Casa 1',
		'Richard Gonzalez', '46a', '', '', 'Taxista', '',
		'De acuerdo a informacion aportada por la madre, se trata de escolar masculino de 7a, 2m de EC, cursante de 2do de E.B., con 5a de escolaridad, sin Dx. previo y referido por recomendacion de una amiga de la familia, quien es representante en FUNDAFI. La madre informa agresividad y atencion dispersa lo que ha incidido negativamente en su rendimiento escolar. Es producto de IIIG (2 perdidas previas), con duracion de 9m, controlado y sin complicaciones. Nace por parto vaginal, con peso de 3,400kg y talla de 54cm. Proviene de madre sana y padre diabetico. En su desarrollo evolutivo se conoce de: sosten cefalico a los 3m, sedestacion 6m, gateo 9m, bipedestacion 11m, marcha 11m, lenguaje 11m y control de esfinteres 1a.',
		'Actualmente su alimentacion es en porciones pequenas y variadas con mayor incidencia de carbohidratos. Sus habitos de sueno no estan acorde a su e.c., no hace siesta y en el horario nocturno solo duerme 8hras. Presenta sonambulismo e inquietud. Emplea su tiempo libre en juegos variados y asistencia a tareas dirigidas 2hras diarias por 3dias a la semana.',
		'Convive en familia extendida conformada por la madre y su pareja (desde hace 4a), abuelos maternos y el evaluado, con relaciones armonicas entre ellos.',
		'Habitan en vivienda propiedad de los abuelos m, espaciosa y con todos los servicios basicos.',
		'El ingreso lo aportan entre todo el grupo familiar y es destinado a cubrir las necesidades de alimentacion, servicios y educacion.',
		'n/t',
		'',
		'Neurologia y Psicopedagogia.',
		true
	) ON CONFLICT (cedula_paciente) DO UPDATE SET
		id_especialista = EXCLUDED.id_especialista,
		lugar_nacimiento = EXCLUDED.lugar_nacimiento,
		escolaridad = EXCLUDED.escolaridad,
		grado_escolar = EXCLUDED.grado_escolar,
		referido_por = EXCLUDED.referido_por,
		madre_nombre = EXCLUDED.madre_nombre,
		madre_edad = EXCLUDED.madre_edad,
		madre_ci = EXCLUDED.madre_ci,
		madre_telefono = EXCLUDED.madre_telefono,
		madre_ocupacion = EXCLUDED.madre_ocupacion,
		madre_correo = EXCLUDED.madre_correo,
		madre_direccion = EXCLUDED.madre_direccion,
		padre_nombre = EXCLUDED.padre_nombre,
		padre_edad = EXCLUDED.padre_edad,
		padre_ci = EXCLUDED.padre_ci,
		padre_telefono = EXCLUDED.padre_telefono,
		padre_ocupacion = EXCLUDED.padre_ocupacion,
		padre_direccion = EXCLUDED.padre_direccion,
		antecedentes_desarrollo = EXCLUDED.antecedentes_desarrollo,
		aspecto_salud = EXCLUDED.aspecto_salud,
		grupo_familiar = EXCLUDED.grupo_familiar,
		vivienda_entorno = EXCLUDED.vivienda_entorno,
		situacion_economica = EXCLUDED.situacion_economica,
		diagnostico_social = EXCLUDED.diagnostico_social,
		conclusion = EXCLUDED.conclusion,
		plan_accion = EXCLUDED.plan_accion,
		entregado = EXCLUDED.entregado;
	`

	_, err = db.Exec(query, cedula)
	if err != nil {
		log.Fatal("Error inserting social evaluation:", err)
	}

	fmt.Println("Successfully inserted/updated social evaluation for Dyland!")
}
