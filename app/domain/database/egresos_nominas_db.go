package database

import (
	"omhmre.com/centromedico/app/domain/models"
)

// IMPLEMENTACIONES DE EGRESOS (Para resolver dependencias de compilación)

func (d *DB) GetEgresos(f models.Fechas) ([]models.Egreso, models.Respuesta) {
	return []models.Egreso{}, models.Respuesta{Status: 200, Mensaje: "OK"}
}

func (d *DB) PostEgreso(e models.Egreso) models.Respuesta {
	return models.Respuesta{Status: 200, Mensaje: "Egreso registrado"}
}

func (d *DB) PutEgreso(e models.Egreso) models.Respuesta {
	return models.Respuesta{Status: 200, Mensaje: "Egreso actualizado"}
}

func (d *DB) DelEgreso(i models.Id) models.Respuesta {
	return models.Respuesta{Status: 200, Mensaje: "Egreso eliminado"}
}

func (d *DB) GetConfigEgresos() ([]models.ConfigEgreso, models.Respuesta) {
	return []models.ConfigEgreso{}, models.Respuesta{Status: 200, Mensaje: "OK"}
}

func (d *DB) PostConfigEgreso(c models.ConfigEgreso) models.Respuesta {
	return models.Respuesta{Status: 200, Mensaje: "Configuración registrada"}
}

func (d *DB) DelConfigEgreso(i models.Id) models.Respuesta {
	return models.Respuesta{Status: 200, Mensaje: "Configuración eliminada"}
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

// IMPLEMENTACIONES DE PERSONAL (Para resolver dependencias de compilación)

func (d *DB) GetPersonal() ([]models.PersonalModel, models.Respuesta) {
	return []models.PersonalModel{}, models.Respuesta{Status: 200, Mensaje: "OK"}
}

func (d *DB) PostPersonal(p models.PersonalModel) models.Respuesta {
	return models.Respuesta{Status: 200, Mensaje: "Personal registrado"}
}

func (d *DB) UpdPersonal(p models.PersonalModel) models.Respuesta {
	return models.Respuesta{Status: 200, Mensaje: "Personal actualizado"}
}

func (d *DB) DelPersonal(i models.Id) models.Respuesta {
	return models.Respuesta{Status: 200, Mensaje: "Personal eliminado"}
}
