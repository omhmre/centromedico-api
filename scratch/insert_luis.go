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

	query := `
	INSERT INTO medi001.evaluaciones_sociales (
		cedula_paciente, id_especialista, lugar_nacimiento, escolaridad, grado_escolar, referido_por,
		madre_nombre, madre_edad, madre_ci, madre_telefono, madre_ocupacion, madre_correo, madre_direccion,
		padre_nombre, padre_edad, padre_ci, padre_telefono, padre_ocupacion, padre_direccion,
		antecedentes_desarrollo, aspecto_salud, grupo_familiar, vivienda_entorno, situacion_economica,
		diagnostico_social, conclusion, plan_accion, entregado
	) VALUES (
		'85', 1, 'El Valle, 01-11-22', 'E.I', 'Nivel maternal', 'Iniciativa familiar',
		'Paloma Reyes', '32a', '20.535.151', '0412-116-66-48', 'docente E.I', 'reyesdeplacencia@gmail.com', 'San Juan Bautista, calle La Fe, sector Las Huertas, casa s/#, Mun. Díaz.',
		'Luis Placencia', '36a', '18.939.222', '0424-848-49-90', 'Representante de ventas - Grupo Amanecer.', 'la misma.',
		'De acuerdo a información aportada por los padres, se trata de lactante mayor de 2a, 8m de E.C, cursante de nivel maternal de E.I, evaluado por neurología y diagnosticado con TEA nivel 1-2. Es producto de III g, controlada, a tiempo y sin complicaciones. Nace mediante cesárea (cp), con peso de 3,200 kg y talla de 53 cm. De su desarrollo evolutivo se menciona: sostén cefálico a los 3m, sedestación 6m, gateo ausente, marcha 1a, lenguaje 1a, 3m muy pocas palabras y control de esfínteres aún en proceso de consolidación.',
		'En la actualidad ha sido medicado con melatonina para corregir trastorno del sueño. Su alimentación está guiada por dieta neuroprotectora. Los padres afirman que ha mejorado su conducta desde que inició control neurológico. Refieren que le agrada asistir a la escuela y compartir con sus pares, sin embargo su lenguaje oral aún está basado en señas y gritos, mantiene contacto tipo flash, estereotipias motoras, autoagresión y pataletas.',
		'Pertenece a familia extendida, conformada por los padres, el evaluado, una hermana (14a), un hermano (7a) y la abuela p. Se mencionan adecuadas relaciones interpersonales en este grupo y de aceptación.',
		'Habitan en vivienda propia, tipo casa, ubicada en zona residencial, acceso a servicios básicos, elaborada con materiales aptos para la habitabilidad y con espacio acorde al número de personas.',
		'El ingreso económico proviene del padre, quien se desempeña como vendedor para una empresa de productos de consumo masivo y de la madre, quien labora como docente en una institución educativa.',
		'TEA 1-2. Dra. Martha Hernández.',
		'De la presente entrevista se concluye, familia con nivel socioeconómico medio-bajo.',
		'Psicopedagogía, T.L y T.O.',
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

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success!")
}
