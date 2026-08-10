package database

import (
	"fmt"
	"time"

	"omhmre.com/centromedico/app/domain/models"
	"omhmre.com/centromedico/app/domain/utils"
)

// InitAbonosTables crea las tablas medi001.paciente_abonos y medi001.paciente_consumos si no existen.
func (d *DB) InitAbonosTables() {
	if _, err := d.db.Exec(`CREATE SCHEMA IF NOT EXISTS medi001;`); err != nil {
		utils.CreateLog("Error creando esquema medi001: " + err.Error())
	}

	queryAbonos := `
	CREATE TABLE IF NOT EXISTS medi001.paciente_abonos (
		id SERIAL PRIMARY KEY,
		cedula_paciente VARCHAR(20) NOT NULL,
		patrocinante VARCHAR(100) NOT NULL DEFAULT 'Particular',
		fecha_abono TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
		monto NUMERIC(12, 2) NOT NULL DEFAULT 0.00,
		tasa NUMERIC(12, 4) NOT NULL DEFAULT 1.0000,
		metodo_pago VARCHAR(50) DEFAULT 'Transferencia',
		referencia VARCHAR(100) DEFAULT '',
		observaciones TEXT DEFAULT '',
		creado_por VARCHAR(100) DEFAULT '',
		fecha_creacion TIMESTAMP WITH TIME ZONE DEFAULT NOW()
	);`

	queryConsumos := `
	CREATE TABLE IF NOT EXISTS medi001.paciente_consumos (
		id SERIAL PRIMARY KEY,
		id_abono INT NULL REFERENCES medi001.paciente_abonos(id) ON DELETE SET NULL,
		cedula_paciente VARCHAR(20) NOT NULL,
		id_cita INT NULL,
		fecha_consumo TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
		especialidad VARCHAR(100) NOT NULL DEFAULT 'General',
		servicio VARCHAR(150) DEFAULT '',
		monto NUMERIC(12, 2) NOT NULL DEFAULT 0.00,
		observaciones TEXT DEFAULT '',
		creado_por VARCHAR(100) DEFAULT '',
		fecha_creacion TIMESTAMP WITH TIME ZONE DEFAULT NOW()
	);`

	if _, err := d.db.Exec(queryAbonos); err != nil {
		utils.CreateLog("Error creando tabla paciente_abonos: " + err.Error())
	}
	if _, err := d.db.Exec(queryConsumos); err != nil {
		utils.CreateLog("Error creando tabla paciente_consumos: " + err.Error())
	}
}

// GetAbonos obtiene la lista de abonos filtrados por cedula y/o patrocinante.
func (d *DB) GetAbonos(cedula string, patrocinante string) ([]models.PacienteAbono, models.Respuesta) {
	var rp models.Respuesta
	list := make([]models.PacienteAbono, 0)

	sqlQuery := `
		SELECT a.id, a.cedula_paciente, COALESCE(p.nombres, ''), a.patrocinante, a.fecha_abono, a.monto, a.tasa,
		       a.metodo_pago, a.referencia, a.observaciones, COALESCE(a.creado_por, ''), a.fecha_creacion
		FROM medi001.paciente_abonos a
		LEFT JOIN medi001.pacientes p ON (a.cedula_paciente = p.cedula OR a.cedula_paciente = p.id::text)
		WHERE ($1 = '' OR a.cedula_paciente = $1 OR p.cedula = $1 OR p.id::text = $1)
		  AND ($2 = '' OR LOWER(a.patrocinante) LIKE LOWER($2))
		ORDER BY a.fecha_abono DESC;`

	patFilter := ""
	if patrocinante != "" {
		patFilter = "%" + patrocinante + "%"
	}

	rows, err := d.db.Query(sqlQuery, cedula, patFilter)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al consultar abonos: " + err.Error()
		return list, rp
	}
	defer rows.Close()

	for rows.Next() {
		var a models.PacienteAbono
		errScan := rows.Scan(
			&a.ID, &a.CedulaPaciente, &a.NombrePaciente, &a.Patrocinante, &a.FechaAbono, &a.Monto, &a.Tasa,
			&a.MetodoPago, &a.Referencia, &a.Observaciones, &a.CreadoPor, &a.FechaCreacion,
		)
		if errScan == nil {
			list = append(list, a)
		}
	}

	rp.Status = 200
	rp.Mensaje = fmt.Sprintf("%d abonos encontrados", len(list))
	return list, rp
}

// PostAbono registra un nuevo abono.
func (d *DB) PostAbono(a models.PacienteAbono) models.Respuesta {
	var rp models.Respuesta

	if a.FechaAbono.IsZero() {
		a.FechaAbono = time.Now()
	}
	if a.Tasa <= 0 {
		a.Tasa = 1.0
	}
	if a.Patrocinante == "" {
		a.Patrocinante = "Particular"
	}

	sqlInsert := `
		INSERT INTO medi001.paciente_abonos
		(cedula_paciente, patrocinante, fecha_abono, monto, tasa, metodo_pago, referencia, observaciones, creado_por)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id;`

	var newID int
	err := d.db.QueryRow(
		sqlInsert,
		a.CedulaPaciente, a.Patrocinante, a.FechaAbono, a.Monto, a.Tasa,
		a.MetodoPago, a.Referencia, a.Observaciones, a.CreadoPor,
	).Scan(&newID)

	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al guardar el abono: " + err.Error()
		return rp
	}

	rp.Status = 200
	rp.Mensaje = fmt.Sprintf("Abono N° %d registrado correctamente", newID)
	return rp
}

// DeleteAbono elimina un abono por ID.
func (d *DB) DeleteAbono(id int) models.Respuesta {
	var rp models.Respuesta
	_, err := d.db.Exec(`DELETE FROM medi001.paciente_abonos WHERE id = $1;`, id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al eliminar abono: " + err.Error()
		return rp
	}
	rp.Status = 200
	rp.Mensaje = "Abono eliminado exitosamente"
	return rp
}

// GetConsumos obtiene la lista de consumos filtrados por cedula.
func (d *DB) GetConsumos(cedula string) ([]models.PacienteConsumo, models.Respuesta) {
	var rp models.Respuesta
	list := make([]models.PacienteConsumo, 0)

	sqlQuery := `
		SELECT c.id, c.id_abono, c.cedula_paciente, COALESCE(p.nombres, ''), c.id_cita, c.fecha_consumo,
		       c.especialidad, c.servicio, c.monto, c.observaciones, COALESCE(c.creado_por, ''), c.fecha_creacion
		FROM medi001.paciente_consumos c
		LEFT JOIN medi001.pacientes p ON (c.cedula_paciente = p.cedula OR c.cedula_paciente = p.id::text)
		WHERE ($1 = '' OR c.cedula_paciente = $1 OR p.cedula = $1 OR p.id::text = $1)
		ORDER BY c.fecha_consumo DESC;`

	rows, err := d.db.Query(sqlQuery, cedula)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al consultar consumos: " + err.Error()
		return list, rp
	}
	defer rows.Close()

	for rows.Next() {
		var c models.PacienteConsumo
		errScan := rows.Scan(
			&c.ID, &c.IDAbono, &c.CedulaPaciente, &c.NombrePaciente, &c.IDCita, &c.FechaConsumo,
			&c.Especialidad, &c.Servicio, &c.Monto, &c.Observaciones, &c.CreadoPor, &c.FechaCreacion,
		)
		if errScan == nil {
			list = append(list, c)
		}
	}

	rp.Status = 200
	rp.Mensaje = fmt.Sprintf("%d consumos encontrados", len(list))
	return list, rp
}

// PostConsumo registra un nuevo consumo descontado.
func (d *DB) PostConsumo(c models.PacienteConsumo) models.Respuesta {
	var rp models.Respuesta

	if c.FechaConsumo.IsZero() {
		c.FechaConsumo = time.Now()
	}

	sqlInsert := `
		INSERT INTO medi001.paciente_consumos
		(id_abono, cedula_paciente, id_cita, fecha_consumo, especialidad, servicio, monto, observaciones, creado_por)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id;`

	var newID int
	err := d.db.QueryRow(
		sqlInsert,
		c.IDAbono, c.CedulaPaciente, c.IDCita, c.FechaConsumo,
		c.Especialidad, c.Servicio, c.Monto, c.Observaciones, c.CreadoPor,
	).Scan(&newID)

	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al registrar el consumo: " + err.Error()
		return rp
	}

	rp.Status = 200
	rp.Mensaje = fmt.Sprintf("Consumo N° %d registrado correctamente", newID)
	return rp
}

// DeleteConsumo elimina un consumo por ID.
func (d *DB) DeleteConsumo(id int) models.Respuesta {
	var rp models.Respuesta
	_, err := d.db.Exec(`DELETE FROM medi001.paciente_consumos WHERE id = $1;`, id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al eliminar consumo: " + err.Error()
		return rp
	}
	rp.Status = 200
	rp.Mensaje = "Consumo eliminado exitosamente"
	return rp
}

// GetEstadoCuentaAbonos genera el estado de cuenta y rendición de cuentas para Digitel / Paciente.
func (d *DB) GetEstadoCuentaAbonos(cedula string, patrocinante string, desde string, hasta string) (models.EstadoCuentaAbonos, models.Respuesta) {
	var rp models.Respuesta
	var ec models.EstadoCuentaAbonos
	ec.Abonos = make([]models.PacienteAbono, 0)
	ec.Consumos = make([]models.PacienteConsumo, 0)
	ec.CedulaPaciente = cedula
	ec.Patrocinante = patrocinante

	// Obtener Nombre del paciente si cedula está presente
	if cedula != "" {
		_ = d.db.QueryRow(`SELECT COALESCE(nombres, '') FROM medi001.pacientes WHERE cedula = $1;`, cedula).Scan(&ec.NombrePaciente)
	}

	// 1. Cargar Abonos
	abonos, _ := d.GetAbonos(cedula, patrocinante)
	ec.Abonos = abonos

	// Si patrocinante está vacío y el paciente posee abonos, usar el patrocinante del primer abono
	if ec.Patrocinante == "" && len(ec.Abonos) > 0 {
		for _, a := range ec.Abonos {
			if a.Patrocinante != "" {
				ec.Patrocinante = a.Patrocinante
				break
			}
		}
	}

	// 2. Cargar Consumos
	consumos, _ := d.GetConsumos(cedula)
	ec.Consumos = consumos

	// 3. Calcular Totales
	var totAbonado float64 = 0.0
	for _, a := range ec.Abonos {
		totAbonado += a.Monto
	}

	var totConsumido float64 = 0.0
	for _, c := range ec.Consumos {
		totConsumido += c.Monto
	}

	ec.TotalAbonado = totAbonado
	ec.TotalConsumido = totConsumido
	ec.SaldoDisponible = totAbonado - totConsumido

	rp.Status = 200
	rp.Mensaje = "Estado de cuenta generado correctamente"
	return ec, rp
}
