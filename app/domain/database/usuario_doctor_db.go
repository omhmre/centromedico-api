package database

import (
	"omhmre.com/centromedico/app/domain/models"
	"omhmre.com/centromedico/app/domain/utils"
)

// GetDoctoresPorUsuario obtiene la lista de especialistas (doctores) asignados a un usuario administrativo.
func (d *DB) GetDoctoresPorUsuario(userID int64) ([]models.DoctoresModel, models.Respuesta) {
	var rp models.Respuesta
	doctores := []models.DoctoresModel{}

	query := `
		SELECT 
			d.id, 
			COALESCE(d.nombres, '') as nombres, 
			COALESCE(d.espec, '') as espec, 
			COALESCE(d.dir, '') as dir, 
			COALESCE(d.tlf, '') as tlf, 
			COALESCE(d.correo, '') as correo, 
			COALESCE(d.whatsapp, '') as whatsapp, 
			COALESCE(d.instagram, '') as instagram, 
			COALESCE(d.tasapago, 0.0) as tasapago,
			COALESCE(d.days_of_week, '[]'::jsonb) as days_of_week, 
			COALESCE(d.start_time, '08:00') as start_time, 
			COALESCE(d.end_time, '18:00') as end_time, 
			COALESCE(d.slot_duration, 45) as slot_duration, 
			COALESCE(d.monto_cita, 0.0) as monto_cita, 
			COALESCE(d.es_medico, false) as es_medico,
			COALESCE(d.titulo, '') as titulo, 
			COALESCE(d.titulo_academico, '') as titulo_academico, 
			COALESCE(d.num_mpps, '') as num_mpps, 
			COALESCE(d.num_cm, '') as num_cm, 
			COALESCE(d.rif, '') as rif, 
			COALESCE(d.servicios, '[]'::jsonb) as servicios, 
			COALESCE(d.cedula, '') as cedula,
			COALESCE(to_char(d.fecha_nacimiento, 'YYYY-MM-DD'), '') as fecha_nacimiento, 
			COALESCE(to_char(d.fecha_ingreso, 'YYYY-MM-DD'), '') as fecha_ingreso, 
			COALESCE(d.sueldo, 0.0) as sueldo, 
			COALESCE(d.frecuencia_pago, 'Mensual') as frecuencia_pago, 
			COALESCE(d.activo, true) as activo,
			COALESCE(d.fecha_retiro, '') as fecha_retiro, 
			COALESCE(d.motivo_retiro, '') as motivo_retiro
		FROM seguridad.usuario_doctor ud
		INNER JOIN medi001.doctores d ON ud.id_doctor = d.id
		WHERE ud.id_usuario = $1
		ORDER BY d.nombres;
	`

	rows, err := d.db.Query(query, userID)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener doctores asignados: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return doctores, rp
	}
	defer rows.Close()

	for rows.Next() {
		doc, err := scanDoctorRow(rows)
		if err != nil {
			rp.Status = 500
			rp.Mensaje = "Error al escanear doctor asignado: " + err.Error()
			utils.CreateLog(rp.Mensaje)
			return doctores, rp
		}
		doctores = append(doctores, doc)
	}

	rp.Status = 200
	rp.Mensaje = "Doctores asignados obtenidos correctamente"
	return doctores, rp
}

// SaveDoctoresUsuario reemplaza las asignaciones de especialistas para un usuario administrativo.
func (d *DB) SaveDoctoresUsuario(userID int64, doctorIDs []int) models.Respuesta {
	var rp models.Respuesta

	// Iniciar transacción
	tx, err := d.db.Begin()
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al iniciar transacción: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return rp
	}

	// Defer rollback en caso de error
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 1. Eliminar asignaciones anteriores
	_, err = tx.Exec("DELETE FROM seguridad.usuario_doctor WHERE id_usuario = $1", userID)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al eliminar asignaciones previas: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return rp
	}

	// 2. Insertar las nuevas asignaciones
	if len(doctorIDs) > 0 {
		stmt, errPrepare := tx.Prepare("INSERT INTO seguridad.usuario_doctor (id_usuario, id_doctor) VALUES ($1, $2)")
		if errPrepare != nil {
			err = errPrepare
			rp.Status = 500
			rp.Mensaje = "Error al preparar inserción de asignación: " + err.Error()
			utils.CreateLog(rp.Mensaje)
			return rp
		}
		defer stmt.Close()

		for _, docID := range doctorIDs {
			_, err = stmt.Exec(userID, docID)
			if err != nil {
				rp.Status = 500
				rp.Mensaje = "Error al guardar asignación para doctor ID: " + err.Error()
				utils.CreateLog(rp.Mensaje)
				return rp
			}
		}
	}

	// 3. Confirmar la transacción
	err = tx.Commit()
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al confirmar transaccion de asignaciones: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return rp
	}

	rp.Status = 200
	rp.Mensaje = "Asignaciones de especialistas guardadas exitosamente"
	return rp
}
