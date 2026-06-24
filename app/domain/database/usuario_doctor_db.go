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
		SELECT d.id, d.nombres, d.espec, d.dir, d.tlf, d.correo, d.whatsapp, d.instagram, d.tasapago
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
		var doc models.DoctoresModel
		err := rows.Scan(
			&doc.Id,
			&doc.Nombres,
			&doc.Espec,
			&doc.Dir,
			&doc.Tlf,
			&doc.Correo,
			&doc.Whatsapp,
			&doc.Instagram,
			&doc.Tasapago,
		)
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
