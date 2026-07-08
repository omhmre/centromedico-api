package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var tables1To1_id = []string{
	"medi001.paciente_antecedentes",
}

var tables1ToN_id = []string{
	"medi001.informes_medicos",
	"medi001.paciente_precios_especialidad",
	"medi001.paciente_signos_vitales",
	"medi001.informe_medico",
}

var tables1ToN_cedulaPaciente = []string{
	"medi001.historial_clinico",
	"medi001.registros_clinicos",
	"medi001.evaluaciones_sociales",
}

var tables1ToN_cedula = []string{
	"medi001.citas",
}

func main() {
	connStr := "postgresql://postgres.sczbdxihtitkuesatpng:Omhmre2025*@aws-1-us-east-2.pooler.supabase.com:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = mergePatients(db, "140", "619")
	if err != nil {
		log.Printf("Failed to merge 619 into 140: %v\n", err)
	} else {
		log.Println("Successfully merged 619 into 140")
	}

	err = mergePatients(db, "186", "576")
	if err != nil {
		log.Printf("Failed to merge 576 into 186: %v\n", err)
	} else {
		log.Println("Successfully merged 576 into 186")
	}
}

func mergePatients(db *sql.DB, originalCedula, duplicateCedula string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var origId, dupId int
	log.Println("QueryRow..."); err = tx.QueryRow("SELECT id FROM medi001.pacientes WHERE cedula = $1 LIMIT 1", originalCedula).Scan(&origId)
	if err != nil {
		return fmt.Errorf("could not find original patient %s: %v", originalCedula, err)
	}
	log.Println("QueryRow..."); err = tx.QueryRow("SELECT id FROM medi001.pacientes WHERE cedula = $1 LIMIT 1", duplicateCedula).Scan(&dupId)
	if err != nil {
		return fmt.Errorf("could not find duplicate patient %s: %v", duplicateCedula, err)
	}

	var oWhatsapp, oCorreo, oDireccion, oFenac, oRepresentante sql.NullString
	log.Println("QueryRow..."); err = tx.QueryRow(`SELECT whatsapp, correo, direccion, fenac, representante FROM medi001.pacientes WHERE id = $1`, origId).Scan(
		&oWhatsapp, &oCorreo, &oDireccion, &oFenac, &oRepresentante,
	)
	if err != nil {
		return err
	}

	var dWhatsapp, dCorreo, dDireccion, dFenac, dRepresentante sql.NullString
	log.Println("QueryRow..."); err = tx.QueryRow(`SELECT whatsapp, correo, direccion, fenac, representante FROM medi001.pacientes WHERE id = $1`, dupId).Scan(
		&dWhatsapp, &dCorreo, &dDireccion, &dFenac, &dRepresentante,
	)
	if err != nil {
		return err
	}

	takeDup := func(orig, dup sql.NullString) *string {
		if (!orig.Valid || orig.String == "") && (dup.Valid && dup.String != "") {
			return &dup.String
		}
		return nil
	}

	updates := map[string]*string{
		"whatsapp": takeDup(oWhatsapp, dWhatsapp),
		"correo": takeDup(oCorreo, dCorreo),
		"direccion": takeDup(oDireccion, dDireccion),
		"fenac": takeDup(oFenac, dFenac),
		"representante": takeDup(oRepresentante, dRepresentante),
	}

	for col, val := range updates {
		if val != nil {
			log.Println("Exec..."); _, err = tx.Exec(fmt.Sprintf("UPDATE medi001.pacientes SET %s = $1 WHERE id = $2", col), *val, origId)
			if err != nil {
				return err
			}
		}
	}

	for _, table := range tables1To1_id {
		var exists bool
		log.Println("QueryRow..."); err = tx.QueryRow(fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE id_paciente = $1)", table), origId).Scan(&exists)
		if err != nil { return err }

		if exists {
			log.Println("Exec..."); _, err = tx.Exec(fmt.Sprintf("DELETE FROM %s WHERE id_paciente = $1", table), dupId)
			if err != nil { return err }
		} else {
			log.Println("Exec..."); _, err = tx.Exec(fmt.Sprintf("UPDATE %s SET id_paciente = $1 WHERE id_paciente = $2", table), origId, dupId)
			if err != nil { return err }
		}
	}

	for _, table := range tables1ToN_id {
		log.Println("Exec..."); _, err = tx.Exec(fmt.Sprintf("UPDATE %s SET id_paciente = $1 WHERE id_paciente = $2", table), origId, dupId)
		if err != nil { return err }
	}

	for _, table := range tables1ToN_cedulaPaciente {
		log.Println("Exec..."); _, err = tx.Exec(fmt.Sprintf("UPDATE %s SET cedula_paciente = $1 WHERE cedula_paciente = $2", table), originalCedula, duplicateCedula)
		if err != nil { return err }
	}

	for _, table := range tables1ToN_cedula {
		log.Println("Exec..."); _, err = tx.Exec(fmt.Sprintf("UPDATE %s SET cedula = $1 WHERE cedula = $2", table), originalCedula, duplicateCedula)
		if err != nil { return err }
	}

	log.Println("Exec..."); _, err = tx.Exec("DELETE FROM medi001.pacientes WHERE id = $1", dupId)
	if err != nil {
		return err
	}

	return tx.Commit()
}
