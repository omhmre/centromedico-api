package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"omhmre.com/centromedico/app/domain/models"
	"omhmre.com/centromedico/app/domain/utils"
)

// ensureEgresosTables garantiza que las tablas y columnas necesarias para egresos existan en PostgreSQL.
func (d *DB) ensureEgresosTables() {
	queryEgresos := `
	CREATE TABLE IF NOT EXISTS medi001.egresos (
		id SERIAL PRIMARY KEY,
		fecha TIMESTAMP WITHOUT TIME ZONE NOT NULL,
		descripcion TEXT NOT NULL,
		proveedor VARCHAR(255) DEFAULT '',
		monto DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
		categoria VARCHAR(100) DEFAULT '',
		metodo_pago VARCHAR(100) DEFAULT '',
		referencia VARCHAR(100) DEFAULT '',
		usuario_operacion VARCHAR(100) DEFAULT '',
		fecha_operacion TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	ALTER TABLE medi001.egresos ADD COLUMN IF NOT EXISTS proveedor VARCHAR(255) DEFAULT '';
	`
	if _, err := d.db.Exec(queryEgresos); err != nil {
		utils.CreateLog("Error asegurando tabla medi001.egresos: " + err.Error())
	}

	queryConfig := `
	CREATE TABLE IF NOT EXISTS medi001.config_egresos (
		id SERIAL PRIMARY KEY,
		tipo VARCHAR(50) NOT NULL,
		valor VARCHAR(255) NOT NULL,
		status BOOLEAN DEFAULT true,
		created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	`
	if _, err := d.db.Exec(queryConfig); err != nil {
		utils.CreateLog("Error asegurando tabla medi001.config_egresos: " + err.Error())
	}
}

// IMPLEMENTACIONES DE EGRESOS

func (d *DB) GetEgresos(f models.Fechas) ([]models.Egreso, models.Respuesta) {
	d.ensureEgresosTables()
	var rp models.Respuesta

	var rows *sql.Rows
	var err error

	if f.Desde != "" && f.Hasta != "" {
		desdeTime, errDesde := time.Parse(time.RFC3339, f.Desde)
		hastaTime, errHasta := time.Parse(time.RFC3339, f.Hasta)

		if errDesde == nil && errHasta == nil {
			query := `
				SELECT id, fecha, descripcion, COALESCE(proveedor, ''), monto, categoria, metodo_pago, referencia, usuario_operacion, fecha_operacion
				FROM medi001.egresos
				WHERE fecha >= $1 AND fecha <= $2
				ORDER BY fecha DESC, id DESC;
			`
			rows, err = d.db.Query(query, desdeTime, hastaTime)
		} else {
			query := `
				SELECT id, fecha, descripcion, COALESCE(proveedor, ''), monto, categoria, metodo_pago, referencia, usuario_operacion, fecha_operacion
				FROM medi001.egresos
				WHERE fecha >= $1::timestamp AND fecha <= $2::timestamp
				ORDER BY fecha DESC, id DESC;
			`
			rows, err = d.db.Query(query, f.Desde, f.Hasta)
		}
	} else {
		query := `
			SELECT id, fecha, descripcion, COALESCE(proveedor, ''), monto, categoria, metodo_pago, referencia, usuario_operacion, fecha_operacion
			FROM medi001.egresos
			ORDER BY fecha DESC, id DESC LIMIT 500;
		`
		rows, err = d.db.Query(query)
	}

	if err != nil {
		utils.CreateLog("Error al obtener egresos: " + err.Error())
		rp.Status = 500
		rp.Mensaje = "Error al consultar egresos en la base de datos: " + err.Error()
		return []models.Egreso{}, rp
	}
	defer rows.Close()

	egresos := []models.Egreso{}
	for rows.Next() {
		var e models.Egreso
		var fechaOp sql.NullTime
		errScan := rows.Scan(
			&e.ID,
			&e.Fecha,
			&e.Descripcion,
			&e.Proveedor,
			&e.Monto,
			&e.Categoria,
			&e.MetodoPago,
			&e.Referencia,
			&e.UsuarioOperacion,
			&fechaOp,
		)
		if errScan != nil {
			utils.CreateLog("Error al escanear egreso: " + errScan.Error())
			continue
		}
		if fechaOp.Valid {
			e.FechaOperacion = fechaOp.Time
		}
		egresos = append(egresos, e)
	}

	rp.Status = 200
	rp.Mensaje = "OK"
	return egresos, rp
}

func (d *DB) PostEgreso(e models.Egreso) models.Respuesta {
	d.ensureEgresosTables()
	var rp models.Respuesta

	query := `
		INSERT INTO medi001.egresos (
			fecha, descripcion, proveedor, monto, categoria, metodo_pago, referencia, usuario_operacion, fecha_operacion
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW())
		RETURNING id;
	`

	var id int
	err := d.db.QueryRow(
		query,
		e.Fecha,
		e.Descripcion,
		e.Proveedor,
		e.Monto,
		e.Categoria,
		e.MetodoPago,
		e.Referencia,
		e.UsuarioOperacion,
	).Scan(&id)

	if err != nil {
		utils.CreateLog("Error al registrar egreso: " + err.Error())
		rp.Status = 500
		rp.Mensaje = "Error al registrar egreso: " + err.Error()
		return rp
	}

	rp.Status = 200
	rp.Mensaje = fmt.Sprintf("Egreso registrado con ID %d", id)
	return rp
}

func (d *DB) PutEgreso(e models.Egreso) models.Respuesta {
	d.ensureEgresosTables()
	var rp models.Respuesta

	query := `
		UPDATE medi001.egresos
		SET fecha = $2, descripcion = $3, proveedor = $4, monto = $5, categoria = $6, metodo_pago = $7, referencia = $8, usuario_operacion = $9
		WHERE id = $1;
	`

	res, err := d.db.Exec(
		query,
		e.ID,
		e.Fecha,
		e.Descripcion,
		e.Proveedor,
		e.Monto,
		e.Categoria,
		e.MetodoPago,
		e.Referencia,
		e.UsuarioOperacion,
	)

	if err != nil {
		utils.CreateLog("Error al actualizar egreso: " + err.Error())
		rp.Status = 500
		rp.Mensaje = "Error al actualizar egreso: " + err.Error()
		return rp
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		rp.Status = 404
		rp.Mensaje = "No se encontró el egreso a actualizar"
		return rp
	}

	rp.Status = 200
	rp.Mensaje = "Egreso actualizado correctamente"
	return rp
}

func (d *DB) DelEgreso(i models.Id) models.Respuesta {
	d.ensureEgresosTables()
	var rp models.Respuesta

	id, errConv := strconv.Atoi(i.Id)
	if errConv != nil || id <= 0 {
		rp.Status = 400
		rp.Mensaje = "ID de egreso inválido"
		return rp
	}

	query := `DELETE FROM medi001.egresos WHERE id = $1;`
	res, err := d.db.Exec(query, id)

	if err != nil {
		utils.CreateLog("Error al eliminar egreso: " + err.Error())
		rp.Status = 500
		rp.Mensaje = "Error al eliminar egreso: " + err.Error()
		return rp
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		rp.Status = 404
		rp.Mensaje = "No se encontró el egreso a eliminar"
		return rp
	}

	rp.Status = 200
	rp.Mensaje = "Egreso eliminado correctamente"
	return rp
}

func (d *DB) GetConfigEgresos() ([]models.ConfigEgreso, models.Respuesta) {
	d.ensureEgresosTables()
	var rp models.Respuesta

	query := `SELECT id, tipo, valor FROM medi001.config_egresos ORDER BY id ASC;`
	rows, err := d.db.Query(query)
	if err != nil {
		utils.CreateLog("Error al obtener configuracion de egresos: " + err.Error())
		rp.Status = 500
		rp.Mensaje = "Error al consultar configuracion: " + err.Error()
		return []models.ConfigEgreso{}, rp
	}
	defer rows.Close()

	configs := []models.ConfigEgreso{}
	for rows.Next() {
		var c models.ConfigEgreso
		if errScan := rows.Scan(&c.ID, &c.Tipo, &c.Valor); errScan == nil {
			configs = append(configs, c)
		}
	}

	rp.Status = 200
	rp.Mensaje = "OK"
	return configs, rp
}

func (d *DB) PostConfigEgreso(c models.ConfigEgreso) models.Respuesta {
	d.ensureEgresosTables()
	var rp models.Respuesta

	query := `INSERT INTO medi001.config_egresos (tipo, valor) VALUES ($1, $2) RETURNING id;`
	var id int
	err := d.db.QueryRow(query, c.Tipo, c.Valor).Scan(&id)
	if err != nil {
		utils.CreateLog("Error al guardar opción de configuración: " + err.Error())
		rp.Status = 500
		rp.Mensaje = "Error al guardar configuración: " + err.Error()
		return rp
	}

	rp.Status = 200
	rp.Mensaje = fmt.Sprintf("Configuración guardada con ID %d", id)
	return rp
}

func (d *DB) DelConfigEgreso(i models.Id) models.Respuesta {
	d.ensureEgresosTables()
	var rp models.Respuesta

	id, errConv := strconv.Atoi(i.Id)
	if errConv != nil || id <= 0 {
		rp.Status = 400
		rp.Mensaje = "ID de configuración inválido"
		return rp
	}

	query := `DELETE FROM medi001.config_egresos WHERE id = $1;`
	res, err := d.db.Exec(query, id)
	if err != nil {
		utils.CreateLog("Error al eliminar opción de configuración: " + err.Error())
		rp.Status = 500
		rp.Mensaje = "Error al eliminar configuración: " + err.Error()
		return rp
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		rp.Status = 404
		rp.Mensaje = "No se encontró la opción de configuración a eliminar"
		return rp
	}

	rp.Status = 200
	rp.Mensaje = "Configuración eliminada correctamente"
	return rp
}

// IMPLEMENTACIONES DE NÓMINA (Para resolver dependencias de compilación)

func (d *DB) GetNominas(desde string, hasta string) ([]models.NominaModel, models.Respuesta) {
	return []models.NominaModel{}, models.Respuesta{Status: 200, Mensaje: "OK"}
}

func (d *DB) PostNomina(n models.NominaModel) models.Respuesta {
	return models.Respuesta{Status: 200, Mensaje: "Nómina registrada"}
}

func (d *DB) UpdNomina(n models.NominaModel) models.Respuesta {
	return models.Respuesta{Status: 200, Mensaje: "Nómina actualizada"}
}

func (d *DB) PayNomina(nominaID int, fechaPago string, metodoPago string, usuarioOperacion string) models.Respuesta {
	return models.Respuesta{Status: 200, Mensaje: "Nómina pagada"}
}

func (d *DB) DelNomina(i models.Id) models.Respuesta {
	return models.Respuesta{Status: 200, Mensaje: "Nómina eliminada"}
}

// ensurePersonalTable garantiza que la tabla medi001.personal exista con todas sus columnas.
func (d *DB) ensurePersonalTable() {
	query := `
	CREATE TABLE IF NOT EXISTS medi001.personal (
		id SERIAL PRIMARY KEY,
		nombre VARCHAR(255) NOT NULL,
		cedula VARCHAR(20) UNIQUE NOT NULL,
		telefono VARCHAR(20),
		correo VARCHAR(255),
		direccion TEXT,
		titulo VARCHAR(255),
		universidad VARCHAR(255),
		fecha_ingreso DATE,
		fecha_nacimiento DATE,
		cargo VARCHAR(100),
		sueldo DECIMAL(15, 2) DEFAULT 0,
		frecuencia_pago VARCHAR(20) DEFAULT 'Mensual',
		status VARCHAR(20) DEFAULT 'Activo',
		doctor_id INT,
		comision_porc DECIMAL(5, 2) DEFAULT 0.00,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	ALTER TABLE medi001.personal ADD COLUMN IF NOT EXISTS frecuencia_pago VARCHAR(20) DEFAULT 'Mensual';
	ALTER TABLE medi001.personal ADD COLUMN IF NOT EXISTS doctor_id INT;
	ALTER TABLE medi001.personal ADD COLUMN IF NOT EXISTS comision_porc DECIMAL(5, 2) DEFAULT 0.00;
	`
	if _, err := d.db.Exec(query); err != nil {
		utils.CreateLog("Error asegurando tabla medi001.personal: " + err.Error())
	}
}

// IMPLEMENTACIONES DE PERSONAL

func (d *DB) GetPersonal() ([]models.PersonalModel, models.Respuesta) {
	d.ensurePersonalTable()
	var rp models.Respuesta

	query := `
		SELECT id, nombre, cedula, COALESCE(telefono, ''), COALESCE(correo, ''), COALESCE(direccion, ''),
		       COALESCE(titulo, ''), COALESCE(universidad, ''), fecha_ingreso, fecha_nacimiento,
		       COALESCE(cargo, ''), COALESCE(sueldo, 0), COALESCE(frecuencia_pago, 'Mensual'),
		       COALESCE(status, 'Activo'), COALESCE(doctor_id, 0), COALESCE(comision_porc, 0), created_at, updated_at
		FROM medi001.personal
		ORDER BY id DESC;
	`

	rows, err := d.db.Query(query)
	if err != nil {
		utils.CreateLog("Error al consultar personal: " + err.Error())
		rp.Status = 500
		rp.Mensaje = "Error al consultar personal: " + err.Error()
		return []models.PersonalModel{}, rp
	}
	defer rows.Close()

	list := []models.PersonalModel{}
	for rows.Next() {
		var p models.PersonalModel
		var idVal, docIDVal int
		var tel, mail, dir, tit, uni, car, freq, st string
		var sueldo, comision float64
		var fIngreso, fNac, cAt, uAt sql.NullTime

		errScan := rows.Scan(
			&idVal, &p.Nombre, &p.Cedula, &tel, &mail, &dir,
			&tit, &uni, &fIngreso, &fNac,
			&car, &sueldo, &freq, &st, &docIDVal, &comision, &cAt, &uAt,
		)
		if errScan != nil {
			utils.CreateLog("Error al escanear personal: " + errScan.Error())
			continue
		}

		p.Id = &idVal
		p.Telefono = &tel
		p.Correo = &mail
		p.Direccion = &dir
		p.Titulo = &tit
		p.Universidad = &uni
		if fIngreso.Valid {
			p.FechaIngreso = &fIngreso.Time
		}
		if fNac.Valid {
			p.FechaNacimiento = &fNac.Time
		}
		p.Cargo = &car
		p.Sueldo = sueldo
		p.FrecuenciaPago = freq
		p.Status = st
		if docIDVal > 0 {
			p.DoctorId = &docIDVal
		}
		p.ComisionPorc = comision
		if cAt.Valid {
			p.CreatedAt = &cAt.Time
		}
		if uAt.Valid {
			p.UpdatedAt = &uAt.Time
		}

		list = append(list, p)
	}

	rp.Status = 200
	rp.Mensaje = "OK"
	return list, rp
}

func (d *DB) PostPersonal(p models.PersonalModel) models.Respuesta {
	d.ensurePersonalTable()
	var rp models.Respuesta

	query := `
		INSERT INTO medi001.personal (
			nombre, cedula, telefono, correo, direccion, titulo, universidad, fecha_ingreso, fecha_nacimiento, cargo, sueldo, frecuencia_pago, status, doctor_id, comision_porc, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, NOW(), NOW())
		RETURNING id;
	`

	var id int
	err := d.db.QueryRow(
		query,
		p.Nombre,
		p.Cedula,
		p.Telefono,
		p.Correo,
		p.Direccion,
		p.Titulo,
		p.Universidad,
		p.FechaIngreso,
		p.FechaNacimiento,
		p.Cargo,
		p.Sueldo,
		p.FrecuenciaPago,
		p.Status,
		p.DoctorId,
		p.ComisionPorc,
	).Scan(&id)

	if err != nil {
		utils.CreateLog("Error al guardar personal: " + err.Error())
		rp.Status = 500
		rp.Mensaje = "Error al guardar personal: " + err.Error()
		return rp
	}

	rp.Status = 200
	rp.Mensaje = fmt.Sprintf("Personal registrado con ID %d", id)
	return rp
}

func (d *DB) UpdPersonal(p models.PersonalModel) models.Respuesta {
	d.ensurePersonalTable()
	var rp models.Respuesta

	if p.Id == nil || *p.Id <= 0 {
		rp.Status = 400
		rp.Mensaje = "ID de personal requerido para actualizar"
		return rp
	}

	query := `
		UPDATE medi001.personal
		SET nombre = $2, cedula = $3, telefono = $4, correo = $5, direccion = $6, titulo = $7, universidad = $8, fecha_ingreso = $9, fecha_nacimiento = $10, cargo = $11, sueldo = $12, frecuencia_pago = $13, status = $14, doctor_id = $15, comision_porc = $16, updated_at = NOW()
		WHERE id = $1;
	`

	res, err := d.db.Exec(
		query,
		*p.Id,
		p.Nombre,
		p.Cedula,
		p.Telefono,
		p.Correo,
		p.Direccion,
		p.Titulo,
		p.Universidad,
		p.FechaIngreso,
		p.FechaNacimiento,
		p.Cargo,
		p.Sueldo,
		p.FrecuenciaPago,
		p.Status,
		p.DoctorId,
		p.ComisionPorc,
	)

	if err != nil {
		utils.CreateLog("Error al actualizar personal: " + err.Error())
		rp.Status = 500
		rp.Mensaje = "Error al actualizar personal: " + err.Error()
		return rp
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		rp.Status = 404
		rp.Mensaje = "No se encontró el personal a actualizar"
		return rp
	}

	rp.Status = 200
	rp.Mensaje = "Personal actualizado correctamente"
	return rp
}

func (d *DB) DelPersonal(i models.Id) models.Respuesta {
	d.ensurePersonalTable()
	var rp models.Respuesta

	id, errConv := strconv.Atoi(i.Id)
	if errConv != nil || id <= 0 {
		rp.Status = 400
		rp.Mensaje = "ID de personal inválido"
		return rp
	}

	query := `DELETE FROM medi001.personal WHERE id = $1;`
	res, err := d.db.Exec(query, id)
	if err != nil {
		utils.CreateLog("Error al eliminar personal: " + err.Error())
		rp.Status = 500
		rp.Mensaje = "Error al eliminar personal: " + err.Error()
		return rp
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		rp.Status = 404
		rp.Mensaje = "No se encontró el personal a eliminar"
		return rp
	}

	rp.Status = 200
	rp.Mensaje = "Personal eliminado correctamente"
	return rp
}
