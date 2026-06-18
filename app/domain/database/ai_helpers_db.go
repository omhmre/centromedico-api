package database

import (
	"database/sql"
	"fmt"
	"time"
	"omhmre.com/centromedico/app/domain/models"
	"omhmre.com/centromedico/app/domain/utils"
)

const sqlGetCitasDoctorFecha = `
SELECT
    c.id AS cita_id,
    c.iddoctor,
    d.nombres AS especialista,
    d.espec AS especialidad,
    c.cedula AS paciente_cedula,
    p.nombres AS paciente,
    c.motivo,
    c.inicio,
    c.fin,
    c.diagnostico,
    c.status AS cita_status,
    c.color,
    c.montoref,
    c.tasa,
    c.montobs,
    c.pagado,
    c.saldo
FROM
    medi001.citas c
INNER JOIN
    medi001.doctores d ON c.iddoctor = d.id
INNER JOIN
    medi001.pacientes p ON c.cedula = p.cedula
WHERE
    c.iddoctor = $1 AND c.inicio::date = $2::date
ORDER BY
    c.inicio ASC;`

func (d *DB) GetDoctoresPorEspecialidad(especialidad string) ([]models.DoctoresModel, models.Respuesta) {
	var rp models.Respuesta
	rows, err := d.db.Query(sqlGetDoctores + " WHERE LOWER(d.espec) = LOWER($1)", especialidad)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener doctores por especialidad: " + err.Error()
		return nil, rp
	}
	defer rows.Close()

	doctores := []models.DoctoresModel{}
	for rows.Next() {
		var doctor models.DoctoresModel
		err := rows.Scan(
			&doctor.Id,
			&doctor.Nombres,
			&doctor.Espec,
			&doctor.Dir,
			&doctor.Tlf,
			&doctor.Correo,
			&doctor.Whatsapp,
			&doctor.Instagram,
			&doctor.Tasapago,
		)
		if err != nil {
			continue
		}
		doctores = append(doctores, doctor)
	}
	rp.Status = 200
	rp.Mensaje = "Doctores obtenidos correctamente"
	return doctores, rp
}

func (d *DB) GetDoctor(medicoID int) (models.DoctoresModel, models.Respuesta) {
	var rp models.Respuesta
	var doctor models.DoctoresModel
	row := d.db.QueryRow(sqlGetDoctores + " WHERE d.id = $1", medicoID)
	err := row.Scan(
		&doctor.Id,
		&doctor.Nombres,
		&doctor.Espec,
		&doctor.Dir,
		&doctor.Tlf,
		&doctor.Correo,
		&doctor.Whatsapp,
		&doctor.Instagram,
		&doctor.Tasapago,
	)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener doctor: " + err.Error()
		return doctor, rp
	}
	rp.Status = 200
	rp.Mensaje = "Doctor obtenido correctamente"
	return doctor, rp
}

func (d *DB) GetCitasDoctorFecha(medicoID int, fecha string) ([]models.CitaModel, models.Respuesta) {
	var rp models.Respuesta
	var citas []models.CitaModel

	rows, err := d.db.Query(sqlGetCitasDoctorFecha, medicoID, fecha)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener citas del doctor por fecha: " + err.Error()
		return nil, rp
	}
	defer rows.Close()

	for rows.Next() {
		var cita models.CitaModel
		err := rows.Scan(
			&cita.Id,
			&cita.IdDoctor,
			&cita.Especialista,
			&cita.Especialidad,
			&cita.Cedula,
			&cita.Paciente,
			&cita.Motivo,
			&cita.Inicio,
			&cita.Fin,
			&cita.Diagnostico,
			&cita.Status,
			&cita.Color,
			&cita.Montoref,
			&cita.Tasa,
			&cita.Montobs,
			&cita.Pagado,
			&cita.Saldo,
		)
		if err != nil {
			continue
		}
		citas = append(citas, cita)
	}

	rp.Status = 200
	rp.Mensaje = "Citas obtenidas correctamente"
	return citas, rp
}

func (d *DB) GetPacientePorCedula(cedula string) (models.PacientesModel, models.Respuesta) {
	var rp models.Respuesta
	var paciente models.PacientesModel

	row := d.db.QueryRow(sqlGetPacientes + " WHERE p.cedula = $1", cedula)
	err := row.Scan(
		&paciente.Id,
		&paciente.Cedula,
		&paciente.Nombres,
		&paciente.Fenac,
		&paciente.Representante,
		&paciente.Whatsapp,
		&paciente.Direccion,
		&paciente.Correo,
		&paciente.Diagnostico,
		&paciente.CXC,
	)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener paciente: " + err.Error()
		return paciente, rp
	}

	// Cargar citas
	citas, rpCitas := d.GetCitasPaciente(paciente)
	if rpCitas.Status == 200 {
		paciente.Citas = citas
	}

	// Cargar precios por especialidad
	precios := []models.PrecioEspecialidad{}
	precio := models.PrecioEspecialidad{}
	rowsPrecios, errPrecios := d.db.Query(sqlGetPreciosPorPaciente, paciente.Id)
	if errPrecios == nil {
		defer rowsPrecios.Close()
		for rowsPrecios.Next() {
			if errScan := rowsPrecios.Scan(&precio.Especialidad, &precio.Precio); errScan == nil {
				precios = append(precios, precio)
			}
		}
	}
	paciente.Precios = precios

	// Cargar especialistas
	especialistasVistos := make(map[int]models.EspecialistaAtencion)
	for _, cita := range paciente.Citas {
		if cita.Status == "Completada" {
			if _, ok := especialistasVistos[cita.IdDoctor]; !ok {
				especialistasVistos[cita.IdDoctor] = models.EspecialistaAtencion{
					ID:           cita.IdDoctor,
					Nombres:      cita.Especialista,
					Especialidad: cita.Especialidad,
				}
			}
		}
	}
	paciente.Especialistas = make([]models.EspecialistaAtencion, 0, len(especialistasVistos))
	for _, esp := range especialistasVistos {
		paciente.Especialistas = append(paciente.Especialistas, esp)
	}

	rp.Status = 200
	rp.Mensaje = "Paciente obtenido correctamente"
	return paciente, rp
}

func (d *DB) AddEmailConfig(i models.EmailConfig) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(`INSERT INTO empre001.emailconfig (smtp, puerto, usuario, clave, tls) VALUES ($1, $2, $3, $4, $5);`,
		i.Smtp, i.Port, i.Usuario, i.Clave, i.Tls)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo agregar configuracion de correo: " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Status = 200
		rp.Mensaje = "Configuracion de correo agregada correctamente"
	}
	return rp
}

func (d *DB) DelEmailConfig(i models.Id) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(`DELETE FROM empre001.emailconfig WHERE id = $1;`, i.Id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo eliminar configuracion de correo: " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Status = 200
		rp.Mensaje = "Configuracion de correo eliminada correctamente"
	}
	return rp
}

func (d *DB) ensureLicenseTable() {
	_, err := d.db.Exec(`
	CREATE TABLE IF NOT EXISTS seguridad.licencia (
		id SERIAL PRIMARY KEY,
		key_hash VARCHAR(255) NOT NULL,
		license_key TEXT NOT NULL,
		client_name VARCHAR(255) NOT NULL,
		rif VARCHAR(50) NOT NULL,
		valid_until TIMESTAMP NOT NULL,
		is_premium BOOLEAN NOT NULL,
		activated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
	);`)
	if err != nil {
		utils.CreateLog("Error al crear tabla seguridad.licencia: " + err.Error())
	}
}

func (d *DB) GetActiveLicense() (models.License, error) {
	d.ensureLicenseTable()
	var l models.License
	row := d.db.QueryRow(`
		SELECT id, key_hash, license_key, client_name, rif, valid_until, is_premium, activated_at 
		FROM seguridad.licencia 
		WHERE valid_until > NOW() 
		ORDER BY id DESC LIMIT 1;`)
	err := row.Scan(&l.Id, &l.KeyHash, &l.LicenseKey, &l.ClientName, &l.Rif, &l.ValidUntil, &l.IsPremium, &l.ActivatedAt)
	if err != nil {
		return l, err
	}
	return l, nil
}

func (d *DB) SaveLicense(licenseKey, keyHash, clientName, rif string, validUntil time.Time, isPremium bool) models.Respuesta {
	d.ensureLicenseTable()
	var rp models.Respuesta
	_, err := d.db.Exec(`
		INSERT INTO seguridad.licencia (license_key, key_hash, client_name, rif, valid_until, is_premium)
		VALUES ($1, $2, $3, $4, $5, $6);`,
		licenseKey, keyHash, clientName, rif, validUntil, isPremium)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al guardar la licencia: " + err.Error()
		return rp
	}
	rp.Status = 200
	rp.Mensaje = "Licencia guardada correctamente"
	return rp
}

func (d *DB) GetParametroValor(parametro string) string {
	var valor string
	err := d.db.QueryRow(`SELECT valor FROM seguridad.parametros WHERE parametro = $1;`, parametro).Scan(&valor)
	if err != nil {
		return ""
	}
	return valor
}

func (d *DB) UpdPacienteMatricula(paciente models.PacientesModel) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(`UPDATE medi001.pacientes SET matricula = $2 WHERE id = $1;`, paciente.Id, paciente.Matricula)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al actualizar matricula del paciente: " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Status = 200
		rp.Mensaje = "Matricula del paciente actualizada correctamente"
	} else {
		rp.Status = 404
		rp.Mensaje = "Paciente no encontrado"
	}
	return rp
}

func (d *DB) UpdPacienteStatus(paciente models.PacientesModel) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(`UPDATE medi001.pacientes SET status = $2 WHERE id = $1;`, paciente.Id, paciente.Status)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al actualizar status del paciente: " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Status = 200
		rp.Mensaje = "Status del paciente actualizado correctamente"
	} else {
		rp.Status = 404
		rp.Mensaje = "Paciente no encontrado"
	}
	return rp
}

func (d *DB) FetchExchangeRate() (float64, models.Respuesta) {
	var rp models.Respuesta
	var tasa float64
	err := d.db.QueryRow(`select tasabs from empre001.divisas where id = 1;`).Scan(&tasa)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener tasa de cambio: " + err.Error()
		return 0, rp
	}
	rp.Status = 200
	rp.Mensaje = "Tasa de cambio obtenida correctamente"
	return tasa, rp
}

func (d *DB) UpdateUnpaidAppointmentsVESRate(rate float64) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(`
		UPDATE medi001.citas 
		SET tasa = $1, montobs = montoref * $1 
		WHERE status != 'Completada' OR pagado < montoref;`, rate)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al actualizar la tasa de las citas: " + err.Error()
		return rp
	}
	datos, _ := resp.RowsAffected()
	rp.Status = 200
	rp.Mensaje = fmt.Sprintf("Se actualizo la tasa a %.2f para %d citas no pagadas.", rate, datos)
	return rp
}

func (d *DB) GetAntecedentes(idPaciente int) (models.Antecedentes, models.Respuesta) {
	var rp models.Respuesta
	var a models.Antecedentes
	row := d.db.QueryRow(`
		SELECT id_paciente, COALESCE(medicos, ''), COALESCE(quirurgicos, ''), COALESCE(alergicos, ''), 
		       COALESCE(familiares, ''), COALESCE(habitos, ''), COALESCE(otros, ''), COALESCE(ultima_actualizacion, NOW())
		FROM medi001.paciente_antecedentes WHERE id_paciente = $1;`, idPaciente)
	err := row.Scan(&a.IdPaciente, &a.Medicos, &a.Quirurgicos, &a.Alergicos, &a.Familiares, &a.Habitos, &a.Otros, &a.UltimaActualizacion)
	if err != nil {
		if err == sql.ErrNoRows {
			a.IdPaciente = idPaciente
			rp.Status = 200
			rp.Mensaje = "No se encontraron antecedentes para el paciente, devolviendo vacio"
			return a, rp
		}
		rp.Status = 500
		rp.Mensaje = "Error al obtener antecedentes: " + err.Error()
		return a, rp
	}
	rp.Status = 200
	rp.Mensaje = "Antecedentes obtenidos correctamente"
	return a, rp
}

func (d *DB) UpsertAntecedentes(ant models.Antecedentes) models.Respuesta {
	var rp models.Respuesta
	_, err := d.db.Exec(`
		INSERT INTO medi001.paciente_antecedentes (id_paciente, medicos, quirurgicos, alergicos, familiares, habitos, otros, ultima_actualizacion)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
		ON CONFLICT (id_paciente) DO UPDATE SET
			medicos = EXCLUDED.medicos,
			quirurgicos = EXCLUDED.quirurgicos,
			alergicos = EXCLUDED.alergicos,
			familiares = EXCLUDED.familiares,
			habitos = EXCLUDED.habitos,
			otros = EXCLUDED.otros,
			ultima_actualizacion = NOW();`,
		ant.IdPaciente, ant.Medicos, ant.Quirurgicos, ant.Alergicos, ant.Familiares, ant.Habitos, ant.Otros)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al guardar antecedentes: " + err.Error()
		return rp
	}
	rp.Status = 200
	rp.Mensaje = "Antecedentes guardados correctamente"
	return rp
}

func (d *DB) GetSignosVitales(idPaciente int) ([]models.SignosVitales, models.Respuesta) {
	var rp models.Respuesta
	rows, err := d.db.Query(`
		SELECT id, id_paciente, id_cita, fecha, COALESCE(tension_arterial, ''), COALESCE(frecuencia_cardiaca, 0), 
		       COALESCE(frecuencia_respiratoria, 0), COALESCE(temperatura, 0.0), COALESCE(saturacion_oxigeno, 0), 
		       COALESCE(peso, 0.0), COALESCE(talla, 0.0), COALESCE(imc, 0.0), COALESCE(notas, ''), COALESCE(usuario_operacion, '')
		FROM medi001.paciente_signos_vitales WHERE id_paciente = $1 ORDER BY fecha DESC;`, idPaciente)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener signos vitales: " + err.Error()
		return nil, rp
	}
	defer rows.Close()

	var list []models.SignosVitales
	for rows.Next() {
		var v models.SignosVitales
		err := rows.Scan(
			&v.Id, &v.IdPaciente, &v.IdCita, &v.Fecha, &v.TensionArterial, &v.FrecuenciaCardiaca,
			&v.FrecuenciaRespiratoria, &v.Temperatura, &v.SaturacionOxigeno, &v.Peso, &v.Talla,
			&v.Imc, &v.Notas, &v.UsuarioOperacion,
		)
		if err != nil {
			continue
		}
		list = append(list, v)
	}
	rp.Status = 200
	rp.Mensaje = "Signos vitales obtenidos correctamente"
	return list, rp
}

func (d *DB) PostSignosVitales(v models.SignosVitales) models.Respuesta {
	var rp models.Respuesta
	_, err := d.db.Exec(`
		INSERT INTO medi001.paciente_signos_vitales (id_paciente, id_cita, fecha, tension_arterial, 
			frecuencia_cardiaca, frecuencia_respiratoria, temperatura, saturacion_oxigeno, peso, talla, imc, notas, usuario_operacion)
		VALUES ($1, $2, NOW(), $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`,
		v.IdPaciente, v.IdCita, v.TensionArterial, v.FrecuenciaCardiaca, v.FrecuenciaRespiratoria, v.Temperatura,
		v.SaturacionOxigeno, v.Peso, v.Talla, v.Imc, v.Notas, v.UsuarioOperacion)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al agregar signos vitales: " + err.Error()
		return rp
	}
	rp.Status = 200
	rp.Mensaje = "Signos vitales agregados correctamente"
	return rp
}

func (d *DB) PostInformeMedico(i models.InformeMedico) models.Respuesta {
	var rp models.Respuesta
	if i.Id != nil && *i.Id > 0 {
		_, err := d.db.Exec(`
			UPDATE medi001.informe_medico 
			SET id_doctor = $2, id_cita = $3, diagnostico = $4, evolucion = $5, plan = $6, 
			    recomendaciones = $7, contenido = $8, entregado = $9, fecha_entrega = $10, 
			    modificado_post_entrega = $11, usuario_operacion = $12
			WHERE id = $1;`,
			i.Id, i.IdDoctor, i.IdCita, i.Diagnostico, i.Evolucion, i.Plan, i.Recomendaciones,
			i.Contenido, i.Entregado, i.FechaEntrega, i.ModificadoPostEntrega, i.UsuarioOperacion)
		if err != nil {
			rp.Status = 500
			rp.Mensaje = "Error al actualizar informe medico: " + err.Error()
			return rp
		}
		rp.Status = 200
		rp.Mensaje = "Informe medico actualizado correctamente"
	} else {
		_, err := d.db.Exec(`
			INSERT INTO medi001.informe_medico (id_paciente, id_doctor, id_cita, fecha, diagnostico, 
				evolucion, plan, recomendaciones, contenido, entregado, fecha_entrega, modificado_post_entrega, usuario_operacion)
			VALUES ($1, $2, $3, NOW(), $4, $5, $6, $7, $8, $9, $10, $11, $12);`,
			i.IdPaciente, i.IdDoctor, i.IdCita, i.Diagnostico, i.Evolucion, i.Plan, i.Recomendaciones,
			i.Contenido, i.Entregado, i.FechaEntrega, i.ModificadoPostEntrega, i.UsuarioOperacion)
		if err != nil {
			rp.Status = 500
			rp.Mensaje = "Error al insertar informe medico: " + err.Error()
			return rp
		}
		rp.Status = 200
		rp.Mensaje = "Informe medico creado correctamente"
	}
	return rp
}

func (d *DB) GetInformesMedico(idPaciente int) ([]models.InformeMedico, models.Respuesta) {
	var rp models.Respuesta
	rows, err := d.db.Query(`
		SELECT i.id, i.id_paciente, i.fecha, i.id_doctor, i.id_cita, COALESCE(i.diagnostico, ''), 
		       COALESCE(i.evolucion, ''), COALESCE(i.plan, ''), COALESCE(i.recomendaciones, ''), 
		       COALESCE(i.contenido, ''), COALESCE(i.entregado, false), i.fecha_entrega, 
		       COALESCE(i.modificado_post_entrega, false), COALESCE(i.usuario_operacion, ''),
		       d.nombres, d.espec, d.es_medico, d.titulo, d.titulo_academico, d.num_mpps, d.num_cm, d.rif
		FROM medi001.informe_medico i
		INNER JOIN medi001.doctores d ON i.id_doctor = d.id
		WHERE i.id_paciente = $1 ORDER BY i.fecha DESC;`, idPaciente)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener informes medicos: " + err.Error()
		return nil, rp
	}
	defer rows.Close()

	var list []models.InformeMedico
	for rows.Next() {
		var i models.InformeMedico
		err := rows.Scan(
			&i.Id, &i.IdPaciente, &i.Fecha, &i.IdDoctor, &i.IdCita, &i.Diagnostico,
			&i.Evolucion, &i.Plan, &i.Recomendaciones, &i.Contenido, &i.Entregado, &i.FechaEntrega,
			&i.ModificadoPostEntrega, &i.UsuarioOperacion,
			&i.DoctorNombre, &i.DoctorEspec, &i.EsMedico, &i.DoctorTitulo, &i.DoctorTituloAcademico,
			&i.DoctorMPPS, &i.DoctorCM, &i.DoctorRIF,
		)
		if err != nil {
			continue
		}
		list = append(list, i)
	}
	rp.Status = 200
	rp.Mensaje = "Informes medicos obtenidos correctamente"
	return list, rp
}

func (d *DB) MarkInformeAsDelivered(id int) models.Respuesta {
	var rp models.Respuesta
	_, err := d.db.Exec(`
		UPDATE medi001.informe_medico 
		SET entregado = true, fecha_entrega = NOW() 
		WHERE id = $1;`, id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al marcar informe como entregado: " + err.Error()
		return rp
	}
	rp.Status = 200
	rp.Mensaje = "Informe marcado como entregado correctamente"
	return rp
}

func (d *DB) GetMedicalHistoryTimeline(idPaciente int) ([]models.HistoryTimelineItem, models.Respuesta) {
	var rp models.Respuesta
	var timeline []models.HistoryTimelineItem

	rowsCitas, err := d.db.Query(`
		SELECT c.inicio, 'Consulta' AS tipo, COALESCE(c.motivo, '') AS detalle, d.nombres, d.espec, c.id
		FROM medi001.citas c
		INNER JOIN medi001.doctores d ON c.iddoctor = d.id
		INNER JOIN medi001.pacientes p ON c.cedula = p.cedula
		WHERE p.id = $1 ORDER BY c.inicio DESC;`, idPaciente)
	if err == nil {
		defer rowsCitas.Close()
		for rowsCitas.Next() {
			var item models.HistoryTimelineItem
			rowsCitas.Scan(&item.Fecha, &item.Tipo, &item.Detalle, &item.Doctor, &item.Especialidad, &item.IdReferencia)
			timeline = append(timeline, item)
		}
	}

	rowsInformes, err := d.db.Query(`
		SELECT i.fecha, 'Informe' AS tipo, COALESCE(i.diagnostico, '') AS detalle, d.nombres, d.espec, i.id
		FROM medi001.informe_medico i
		INNER JOIN medi001.doctores d ON i.id_doctor = d.id
		WHERE i.id_paciente = $1 ORDER BY i.fecha DESC;`, idPaciente)
	if err == nil {
		defer rowsInformes.Close()
		for rowsInformes.Next() {
			var item models.HistoryTimelineItem
			rowsInformes.Scan(&item.Fecha, &item.Tipo, &item.Detalle, &item.Doctor, &item.Especialidad, &item.IdReferencia)
			timeline = append(timeline, item)
		}
	}

	rp.Status = 200
	rp.Mensaje = "Linea de tiempo obtenida correctamente"
	return timeline, rp
}

func (d *DB) GetPatientMedicalInsights(idPaciente int) (map[string]interface{}, models.Respuesta) {
	var rp models.Respuesta
	insights := make(map[string]interface{})

	var tArterial, notas string
	var fCardiaca, fRespiratoria, satOxigeno int
	var temp, peso, talla, imc float64
	var fecha time.Time

	row := d.db.QueryRow(`
		SELECT fecha, tension_arterial, frecuencia_cardiaca, frecuencia_respiratoria, temperatura, saturacion_oxigeno, peso, talla, imc, notas
		FROM medi001.paciente_signos_vitales 
		WHERE id_paciente = $1 ORDER BY fecha DESC LIMIT 1;`, idPaciente)
	err := row.Scan(&fecha, &tArterial, &fCardiaca, &fRespiratoria, &temp, &satOxigeno, &peso, &talla, &imc, &notas)
	if err == nil {
		insights["has_vitals"] = true
		insights["latest_vitals_date"] = fecha
		insights["tension_arterial"] = tArterial
		insights["frecuencia_cardiaca"] = fCardiaca
		insights["frecuencia_respiratoria"] = fRespiratoria
		insights["temperatura"] = temp
		insights["saturacion_oxigeno"] = satOxigeno
		insights["peso"] = peso
		insights["talla"] = talla
		insights["imc"] = imc
		insights["notas_vitales"] = notas
	} else {
		insights["has_vitals"] = false
	}

	insights["patient_summary"] = "Datos generales clinicos procesados."

	rp.Status = 200
	rp.Mensaje = "Insights obtenidos correctamente"
	return insights, rp
}
