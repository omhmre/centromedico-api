package database

import (
	"fmt"
	"time"

	"omhmre.com/centromedico/app/domain/models"
	"omhmre.com/centromedico/app/domain/utils"
)

// InitPatrocinantesTable crea la tabla medi001.patrocinantes y siembra datos iniciales.
func (d *DB) InitPatrocinantesTable() {
	if _, err := d.db.Exec(`CREATE SCHEMA IF NOT EXISTS medi001;`); err != nil {
		utils.CreateLog("Error creando esquema medi001: " + err.Error())
	}

	query := `
	CREATE TABLE IF NOT EXISTS medi001.patrocinantes (
		id SERIAL PRIMARY KEY,
		nombre VARCHAR(150) NOT NULL UNIQUE,
		rif VARCHAR(50) DEFAULT '',
		persona_contacto VARCHAR(150) DEFAULT '',
		telefono VARCHAR(50) DEFAULT '',
		email VARCHAR(100) DEFAULT '',
		tipo VARCHAR(50) DEFAULT 'Fundación',
		saldo_total_abonado NUMERIC(12, 2) DEFAULT 0.00,
		nro_becados INT DEFAULT 0,
		observaciones TEXT DEFAULT '',
		activo BOOLEAN DEFAULT true,
		fecha_registro TIMESTAMP WITH TIME ZONE DEFAULT NOW()
	);`

	if _, err := d.db.Exec(query); err != nil {
		utils.CreateLog("Error creando tabla patrocinantes: " + err.Error())
		return
	}

	// Sembrar patrocinantes por defecto si la tabla está vacía
	var count int
	_ = d.db.QueryRow("SELECT COUNT(*) FROM medi001.patrocinantes").Scan(&count)
	if count == 0 {
		seedQuery := `
		INSERT INTO medi001.patrocinantes (nombre, rif, persona_contacto, tipo, observaciones) VALUES
		('Digitel', 'J-30468971-3', 'Fundación Digitel', 'Fundación', 'Programa de Becas Semestrales para Terapias Infantiles'),
		('Fundafi', 'J-40123987-0', 'Gerencia Fundafi', 'Fundación', 'Fondo Social de Asistencia Médica'),
		('Empresas Polar', 'J-00006372-9', 'Coordinación de Gestión Social', 'Corporativo', 'Patrocinio corporativo directo a pacientes'),
		('Particular', '', 'Representante / Familiar', 'Particular', 'Abonos directos realizados por padres o tutores'),
		('Otros', '', 'Donante Anónimo', 'Donante Anónimo', 'Aportes eventuales y donaciones independientes')
		ON CONFLICT (nombre) DO NOTHING;`
		_, _ = d.db.Exec(seedQuery)
	}
}

// GetPatrocinantes obtiene la lista de todos los patrocinantes registrados.
func (d *DB) GetPatrocinantes() ([]models.Patrocinante, models.Respuesta) {
	var rp models.Respuesta
	list := make([]models.Patrocinante, 0)

	sqlQuery := `
		SELECT id, nombre, COALESCE(rif, ''), COALESCE(persona_contacto, ''), COALESCE(telefono, ''), 
		       COALESCE(email, ''), COALESCE(tipo, 'Fundación'), COALESCE(saldo_total_abonado, 0.00),
		       COALESCE(nro_becados, 0), COALESCE(observaciones, ''), COALESCE(activo, true), fecha_registro
		FROM medi001.patrocinantes
		ORDER BY nombre ASC;`

	rows, err := d.db.Query(sqlQuery)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al consultar patrocinantes: " + err.Error()
		return list, rp
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Patrocinante
		errScan := rows.Scan(
			&p.ID, &p.Nombre, &p.RIF, &p.PersonaContacto, &p.Telefono, &p.Email,
			&p.Tipo, &p.SaldoTotalAbonado, &p.NroBecados, &p.Observaciones, &p.Activo, &p.FechaRegistro,
		)
		if errScan == nil {
			list = append(list, p)
		}
	}

	rp.Status = 200
	rp.Mensaje = fmt.Sprintf("%d patrocinantes encontrados", len(list))
	return list, rp
}

// PostPatrocinante registra o actualiza un patrocinante en la base de datos.
func (d *DB) PostPatrocinante(p models.Patrocinante) models.Respuesta {
	var rp models.Respuesta

	if p.FechaRegistro.IsZero() {
		p.FechaRegistro = time.Now()
	}

	sqlQuery := `
		INSERT INTO medi001.patrocinantes (
			nombre, rif, persona_contacto, telefono, email, tipo, observaciones, activo, fecha_registro
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (nombre) DO UPDATE SET
			rif = EXCLUDED.rif,
			persona_contacto = EXCLUDED.persona_contacto,
			telefono = EXCLUDED.telefono,
			email = EXCLUDED.email,
			tipo = EXCLUDED.tipo,
			observaciones = EXCLUDED.observaciones,
			activo = EXCLUDED.activo
		RETURNING id;`

	var newID int
	err := d.db.QueryRow(
		sqlQuery,
		p.Nombre, p.RIF, p.PersonaContacto, p.Telefono, p.Email, p.Tipo, p.Observaciones, p.Activo, p.FechaRegistro,
	).Scan(&newID)

	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al registrar patrocinante: " + err.Error()
		return rp
	}

	rp.Status = 200
	rp.Mensaje = fmt.Sprintf("Patrocinante registrado exitosamente con ID %d", newID)
	return rp
}

// DeletePatrocinante elimina un patrocinante por ID.
func (d *DB) DeletePatrocinante(id int) models.Respuesta {
	var rp models.Respuesta

	sqlQuery := `DELETE FROM medi001.patrocinantes WHERE id = $1;`
	res, err := d.db.Exec(sqlQuery, id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al eliminar patrocinante: " + err.Error()
		return rp
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		rp.Status = 404
		rp.Mensaje = "Patrocinante no encontrado"
		return rp
	}

	rp.Status = 200
	rp.Mensaje = "Patrocinante eliminado correctamente"
	return rp
}
