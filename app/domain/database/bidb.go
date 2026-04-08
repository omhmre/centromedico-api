package database

import (
	"fmt"

	"omhmre.com/centromedico/app/domain/models"
	"omhmre.com/centromedico/app/domain/utils"
)

// GetBIResumenGeneral returns key KPIs for the given date range.
func (d *DB) GetBIResumenGeneral(desde, hasta string) (models.BIResumen, models.Respuesta) {
	var rp models.Respuesta
	var r models.BIResumen

	row := d.db.QueryRow(sqlBIResumenGeneral, desde, hasta)
	err := row.Scan(
		&r.TotalCitas,
		&r.Completadas,
		&r.Canceladas,
		&r.Pendientes,
		&r.TotalIngresosUSD,
		&r.IngresoPorCita,
		&r.PacientesUnicos,
		&r.TasaCompletadas,
	)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener resumen BI: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return r, rp
	}
	rp.Status = 200
	rp.Mensaje = "Resumen obtenido correctamente"
	return r, rp
}

// GetBICitasPorDia returns daily time series of appointments and revenue.
func (d *DB) GetBICitasPorDia(desde, hasta string) ([]models.BITimeSeries, models.Respuesta) {
	var rp models.Respuesta
	var series []models.BITimeSeries

	rows, err := d.db.Query(sqlBICitasPorDia, desde, hasta)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener serie temporal: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return series, rp
	}
	defer rows.Close()

	for rows.Next() {
		var item models.BITimeSeries
		if err := rows.Scan(&item.Fecha, &item.Citas, &item.Ingresos); err != nil {
			utils.CreateLog(fmt.Sprintf("BICitasPorDia scan error: %v", err))
			continue
		}
		series = append(series, item)
	}
	rp.Status = 200
	rp.Mensaje = "Serie temporal obtenida correctamente"
	return series, rp
}

// GetBICitasPorEspecialidad returns appointment stats grouped by specialty.
func (d *DB) GetBICitasPorEspecialidad(desde, hasta string) ([]models.BIEspecialidad, models.Respuesta) {
	var rp models.Respuesta
	var result []models.BIEspecialidad

	rows, err := d.db.Query(sqlBICitasPorEspecialidad, desde, hasta)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener datos por especialidad: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return result, rp
	}
	defer rows.Close()

	for rows.Next() {
		var item models.BIEspecialidad
		if err := rows.Scan(&item.Especialidad, &item.TotalCitas, &item.Completadas, &item.TotalIngresos, &item.TasaEficiencia); err != nil {
			utils.CreateLog(fmt.Sprintf("BIEspecialidad scan error: %v", err))
			continue
		}
		result = append(result, item)
	}
	rp.Status = 200
	rp.Mensaje = "Especialidades obtenidas correctamente"
	return result, rp
}

// GetBIRendimientoDoctor returns performance metrics grouped by doctor.
func (d *DB) GetBIRendimientoDoctor(desde, hasta string) ([]models.BIDoctor, models.Respuesta) {
	var rp models.Respuesta
	var result []models.BIDoctor

	rows, err := d.db.Query(sqlBIRendimientoDoctor, desde, hasta)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener rendimiento de doctores: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return result, rp
	}
	defer rows.Close()

	for rows.Next() {
		var item models.BIDoctor
		if err := rows.Scan(&item.Nombre, &item.TotalCitas, &item.Completadas, &item.Ingresos, &item.Eficiencia); err != nil {
			utils.CreateLog(fmt.Sprintf("BIDoctor scan error: %v", err))
			continue
		}
		result = append(result, item)
	}
	rp.Status = 200
	rp.Mensaje = "Rendimiento de doctores obtenido correctamente"
	return result, rp
}

// GetBIMetodosPago returns payment method breakdown with percentages.
func (d *DB) GetBIMetodosPago(desde, hasta string) ([]models.BIMetodoPago, models.Respuesta) {
	var rp models.Respuesta
	var result []models.BIMetodoPago

	rows, err := d.db.Query(sqlBIMetodosPago, desde, hasta)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener métodos de pago: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return result, rp
	}
	defer rows.Close()

	for rows.Next() {
		var item models.BIMetodoPago
		if err := rows.Scan(&item.Metodo, &item.Total, &item.Porcentaje); err != nil {
			utils.CreateLog(fmt.Sprintf("BIMetodoPago scan error: %v", err))
			continue
		}
		result = append(result, item)
	}
	rp.Status = 200
	rp.Mensaje = "Métodos de pago obtenidos correctamente"
	return result, rp
}

// GetBIHeatmap returns appointment density by day-of-week and hour.
func (d *DB) GetBIHeatmap(desde, hasta string) ([]models.BIHeatmapCell, models.Respuesta) {
	var rp models.Respuesta
	var result []models.BIHeatmapCell

	rows, err := d.db.Query(sqlBIHeatmap, desde, hasta)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener heatmap: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return result, rp
	}
	defer rows.Close()

	for rows.Next() {
		var item models.BIHeatmapCell
		if err := rows.Scan(&item.DiaSemana, &item.Hora, &item.Cantidad); err != nil {
			utils.CreateLog(fmt.Sprintf("BIHeatmap scan error: %v", err))
			continue
		}
		result = append(result, item)
	}
	rp.Status = 200
	rp.Mensaje = "Heatmap obtenido correctamente"
	return result, rp
}
