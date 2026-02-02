package database

import (
	"context"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"gopkg.in/mail.v2"
	"omhmre.com/centromedico/app/domain/models"
	"omhmre.com/centromedico/app/domain/utils"
)

type PostDB interface {
	Open() error
	Close() error
	Ping() error
	GetInventarios(compacto int) ([]models.Inventario, models.Respuesta)
	GetInventarioMenu() ([]models.InventarioNombre, models.Respuesta)
	AddInventario(i models.Inventario) models.Respuesta
	AddItemsInventario(i []models.ItemsInventario) models.Respuesta
	AddPresenInventario(i []models.PresenInventario) models.Respuesta
	UpdInventario(i models.Inventario) models.Respuesta
	DelInventario(i models.Inventario) models.Respuesta
	GetMenu() ([]models.Menu, models.Respuesta)
	GetMenuListado() ([]models.ClaseMenu, models.Respuesta)
	GetClases() ([]models.Clase, models.Respuesta)
	AddClase(i models.Clase) models.Respuesta
	UpdClase(i models.Clase) models.Respuesta
	DelClase(i models.Clase) models.Respuesta
	AddMenu(m models.Menu) models.Respuesta
	AddAllMenu(m []models.Menu) models.Respuesta
	UpdMenu(m models.Menu) models.Respuesta
	DelMenu(m models.Menu) models.Respuesta
	DelAllMenu() models.Respuesta
	GetEmpre() ([]models.Empresa, models.Respuesta)
	AddEmpre(e models.Empresa) models.Respuesta
	UpdEmpresa(e models.Empresa) models.Respuesta
	DelEmpresa(e models.Empresa) models.Respuesta
	GetMenuCompleto() (models.RespMenuCompleto, models.Respuesta)
	GetPrefacturas() ([]models.Prefactura, models.Respuesta)
	GetPrefactura(p models.Prefactura) ([]models.Prefactura, models.Respuesta)
	PostPrefactura(p models.Prefactura) (models.Prefactura, models.Respuesta)
	UpdPrefactura(p models.Prefactura) models.Respuesta
	GetClientes() ([]models.Cliente, models.Respuesta)
	UpdCliente(e models.Cliente) models.Respuesta
	DelCliente(e models.Cliente) models.Respuesta
	AddCliente(e models.Cliente) models.Respuesta
	GetMesas() ([]models.Mesas, models.Respuesta)
	AddMesa(e models.Mesas) models.Respuesta
	DelMesa(e models.Mesas) models.Respuesta
	UpdMesa(e models.Mesas) models.Respuesta
	GetMesoneros() ([]models.Mesoneros, models.Respuesta)
	AddMesonero(e models.Mesoneros) models.Respuesta
	UpdMesonero(e models.Mesoneros) models.Respuesta
	DelMesonero(e models.Mesoneros) models.Respuesta
	AbrirMesa(e models.Mesas) models.Respuesta
	GetInstrumentos() ([]models.InstrumentosPagos, models.Respuesta)
	CleanMesa(e models.Mesas) models.Respuesta
	PostFactura(p models.IdFactura) models.Respuesta
	GetFacturas() ([]models.Factura, models.Respuesta)
	UpdAnularFactura(i models.Factura) models.Respuesta
	PostVentasFactura(datos DT) ([]models.ResVentasDia, models.Respuesta)
	PostVentasProductos(datos DT) ([]models.ResVentasProductos, models.Respuesta)
	UpdDivisa(i models.Divisas) models.Respuesta
	PostPagos(p models.Pagos) models.Respuesta
	GetDetPagosFecha(p models.Fechas) ([]models.DetPago, models.Respuesta)
	GetResDetPagos(p models.Fechas) ([]models.ResumenDetPago, models.Respuesta)
	GetDivisas() ([]models.Divisas, models.Respuesta)
	GetVendedores() ([]models.Vendedor, models.Respuesta)
	GetFacturaId(p models.Factura) (models.Factura, models.Respuesta)
	GetUsers() ([]models.Usuario, models.Respuesta)
	AddUsuario(i models.NuevoUsuario) models.Respuesta
	DelUsuario(i models.Id) models.Respuesta
	UpdateUsuario(u models.Usuario) models.Respuesta
	Login(u models.LoginUsuario) models.LoginData
	ChangePassword(u models.LoginUsuario) models.Respuesta
	GetTopVentas() ([]models.TopVentas, models.Respuesta)
	GetVentasMes() ([]models.VentasMensual, models.Respuesta)
	BackupDatabase() error
	GetProveedores() ([]models.Proveedor, models.Respuesta)
	GetCxcResumen(e models.Fechas) ([]models.CxcResumen, models.Respuesta)
	GetCompras() ([]models.Compra, models.Respuesta)
	GetParametros() ([]models.Parametros, models.Respuesta)
	AddParametro(i models.Parametros) models.Respuesta
	UpdParametro(i models.Parametros) models.Respuesta
	PostProveedor(p models.Proveedor) models.Respuesta
	UpdProveedor(e models.Proveedor) models.Respuesta
	DelProveedor(e models.Id) models.Respuesta
	AddCompra(c models.Compra) models.Respuesta
	GetEmailConfig() ([]models.EmailConfig, models.Respuesta)
	SendMail(f models.MailSend)
	UpdEmailConfig(i models.EmailConfig) models.Respuesta
	GetVentas(f models.Fechas) ([]models.VentasResumen, models.Respuesta)
	PostNotaEntrega(p models.IdFactura) models.Respuesta
	GetInventarioFormal() ([]models.Inventario, models.Respuesta)
	GetNotasEntrega() ([]models.Factura, models.Respuesta)
	GetCxcVencida() ([]models.CxcVencida, models.Respuesta)
	DelCompra(e models.Id) models.Respuesta
	GetPresupuestos() ([]models.Presupuestos, models.Respuesta)
	PostPresupuesto(p models.Id) models.Respuesta
	GetPresupuestoId(p models.Presupuestos) (models.Presupuestos, models.Respuesta)
	DelPresupuesto(e models.Id) models.Respuesta
	GetDoctores() ([]models.DoctoresModel, models.Respuesta)
	UpdDoctores(i models.DoctoresModel) models.Respuesta
	GetPacientes() ([]models.PacientesModel, models.Respuesta)
	PostDoctor(i models.DoctoresModel) models.Respuesta
	DelDoctor(i models.DoctoresModel) models.Respuesta
	PostPaciente(i models.PacientesModel) models.Respuesta
	UpdPaciente(i models.PacientesModel) models.Respuesta
	DelPaciente(i models.PacientesModel) models.Respuesta
	UpsertPrecioEspecialidad(p models.PrecioEspecialidad) models.Respuesta
	DelPrecioEspecialidad(p models.PrecioEspecialidad) models.Respuesta
	GetPayments(p models.Id) ([]models.Payments, models.Respuesta)
	PostPayments(p []models.Payments) models.Respuesta
	GetRelPagos(p models.Fechas) ([]models.RelPagos, models.Respuesta)
	DelPayment(i models.Id) models.Respuesta
	// Citas
	AddCita(cita []models.CitaModel) models.Respuesta
	UpdateCita(cita models.CitaModel) models.Respuesta
	GetCitas() ([]models.CitaModel, models.Respuesta)
	DelCita(e models.IdCitas) models.Respuesta
	GetCitasPaciente(p models.PacientesModel) ([]models.CitaModel, models.Respuesta)
	GetCitasFecha(p models.Fechas) ([]models.CitaModel, models.Respuesta)
	UpdateDiagnosticoCita(cita models.CitaModel) models.Respuesta
	FetchExchangeRate() (float64, models.Respuesta)
	UpdateUnpaidAppointmentsVESRate(newRate float64) models.Respuesta
	// Informes Medicos
	GetInformesMedicos(idPaciente int) ([]models.InformeMedico, models.Respuesta)
	GetInformeMedico(id int) (models.InformeMedico, models.Respuesta)
	AddInformeMedico(i models.InformeMedico) models.Respuesta
	UpdInformeMedico(i models.InformeMedico) models.Respuesta
	DelInformeMedico(id int) models.Respuesta
}

type DB struct {
	db   *sql.DB
	Conn *sql.DB
}
type DT struct {
	Desde string `json:"desde"`
	Hasta string `json:"hasta"`
}

func (db *DB) Ping() error {
	// Ejemplo para PostgreSQL:
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := db.Conn.PingContext(ctx); err != nil {
		return fmt.Errorf("database ping failed: %w", err)
	}
	return nil
}

// roundFloat rounds a float64 to a specified number of decimal places.
func roundFloat(val float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func (d *DB) GetRelPagos(p models.Fechas) ([]models.RelPagos, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetRelPagos, p.Desde, p.Hasta)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Pagos Registrados! " + err.Error()
		utils.CreateLog(err.Error())
		return nil, rp
	}
	defer rows.Close()
	payments := []models.RelPagos{}
	payment := models.RelPagos{}
	var montoCita sql.NullFloat64
	var formaPago sql.NullString
	var montoDoctor sql.NullFloat64

	for rows.Next() {
		err2 :=
			rows.Scan(
				&payment.Doctor_id,
				&payment.Doctor_name,
				&payment.Cita_id,
				&payment.Paciente_nombre,
				&payment.Fecha_pago,
				&montoCita,
				&formaPago,
				&payment.Saldo,
				&payment.Pago_doctor,
				&montoDoctor,
			)
		if err2 != nil {
			utils.CreateLog(err2.Error())
			continue // Skip to the next row on scan error
		}

		if montoCita.Valid {
			payment.Monto_cita = montoCita.Float64
		} else {
			payment.Monto_cita = 0 // Default to 0 if NULL
		}

		if formaPago.Valid {
			payment.Forma_pago = formaPago.String
		} else {
			payment.Forma_pago = "" // Default to empty string if NULL
		}

		if montoDoctor.Valid {
			payment.Monto_doctor = montoDoctor.Float64
		} else {
			payment.Monto_doctor = 0 // Default to 0 if NULL
		}

		payments = append(payments, payment)
	}
	rp.Status = 10
	rp.Mensaje = "Pagos listados correctamente!"
	return payments, rp
}

func (d *DB) PostPayments(p []models.Payments) models.Respuesta {
	var rp models.Respuesta
	var rowsAffected = 0

	for _, e := range p {
		_, err := d.db.Exec(sqlPostPayments, e.Appointmentid, e.Paymentmethod, e.Amount, e.Currency, e.Reference,
			e.Date, e.Status, e.Notes)
		if err != nil {
			rp.Status = 501
			rp.Mensaje = "No se pudo Agregar la Informacion del Pago. " + err.Error()
			utils.CreateLog("No se pudo Agregar la Informacion del Pago. " + err.Error())
			return rp
		} else {
			rowsAffected += 1
		}
	}
	rp.Status = 200
	rp.Mensaje = strconv.FormatInt(int64(rowsAffected), 10) + " pagos agregados correctamente"
	utils.CreateLog(rp.Mensaje)
	// datos, err1 := resp.RowsAffected()
	// if err1 != nil {
	// 	rp.Status = 502
	// 	rp.Mensaje = err1.Error()
	// } else if datos > 0 {
	// 	rp.Status = 200
	// 	rp.Mensaje = strconv.FormatInt(datos, 10) + " Paciente Agregado Correctamente"
	// } else {
	// 	rp.Status = 201
	// 	rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	// }
	return rp
}

func (d *DB) GetPayments(p models.Id) ([]models.Payments, models.Respuesta) {
	var rp models.Respuesta
	var strSql = ""
	if p.Id != "-1" {
		strSql = sqlGetPaymentsByCita
	} else {
		strSql = sqlGetPayments
	}

	rows, err := d.db.Query(strSql, p.Id)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Pagos Registrados! " + err.Error()
		utils.CreateLog(err.Error())
		return nil, rp
	}
	defer rows.Close()
	payments := []models.Payments{}
	payment := models.Payments{}
	for rows.Next() {
		err2 :=
			rows.Scan(
				&payment.Id,
				&payment.Appointmentid,
				&payment.Paymentmethod,
				&payment.Amount,
				&payment.Currency,
				&payment.Reference,
				&payment.Date,
				&payment.Status,
				&payment.Notes,
			)
		if err2 != nil {
			utils.CreateLog(err2.Error())
		}
		payments = append(payments, payment)
	}
	rp.Status = 10
	rp.Mensaje = "Pagos listados correctamente!"
	return payments, rp
}

func (d *DB) DelPaciente(i models.PacientesModel) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlDelPaciente, i.Id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Eliminar el Paciente. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Eliminado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se agrego producto!"
	}
	return rp
}

func (d *DB) UpdPaciente(i models.PacientesModel) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlUpdPaciente,
		i.Id,            // $1
		i.Cedula,        // $2
		i.Nombres,       // $3
		i.Fenac,         // $4
		i.Representante, // $5
		i.Whatsapp,      // $6
		i.Direccion,     // $7
		i.Correo,        // $8
		i.Diagnostico,   // $9
		i.CXC,           // $10
	)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion del Paciente. " + err.Error()
		utils.CreateLog(err.Error())
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registros Actualizados Correctamente"
		rp.Status = 200
		utils.CreateLog(rp.Mensaje)
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) PostPaciente(i models.PacientesModel) models.Respuesta {
	var rp models.Respuesta

	i.CreatedAt = time.Now()

	// Iniciar transacción para garantizar la atomicidad de la operación (insertar y luego actualizar si es necesario).
	tx, err := d.db.Begin()
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al iniciar la transacción: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return rp
	}

	var newID int
	// La consulta de inserción devuelve el nuevo ID. Usamos QueryRow para capturarlo.
	err = tx.QueryRow(sqlPostPaciente, i.Cedula, i.Nombres, i.Fenac, i.Representante, i.Whatsapp,
		i.Direccion, i.Correo, i.Diagnostico, i.CXC, i.CreatedAt).Scan(&newID)

	if err != nil {
		tx.Rollback() // Revertir la transacción en caso de error
		rp.Status = 501
		rp.Mensaje = "No se pudo Agregar la Informacion del Paciente. " + err.Error()
		utils.CreateLog("No se pudo Agregar la Informacion del Paciente. " + err.Error())
		return rp
	}

	// Si la cédula viene vacía, la actualizamos con el ID recién generado.
	// Se verifica que el puntero no sea nulo antes de desreferenciarlo.
	// Si es nulo o si la cadena está vacía, se actualiza.
	updateCedula := false
	if i.Cedula == nil || (i.Cedula != nil && *i.Cedula == "") {
		updateCedula = true
	}
	if updateCedula {
		newCedula := strconv.Itoa(newID)
		_, errUpdate := tx.Exec("UPDATE medi001.pacientes SET cedula = $1 WHERE id = $2", newCedula, newID)
		if errUpdate != nil {
			tx.Rollback() // Revertir si la actualización falla
			rp.Status = 500
			rp.Mensaje = "Error al actualizar la cédula del paciente: " + errUpdate.Error()
			utils.CreateLog(rp.Mensaje)
			return rp
		}
	}

	// Si todo fue bien, confirmar la transacción.
	if err := tx.Commit(); err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al confirmar la transacción: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return rp
	}

	rp.Status = 200
	rp.Mensaje = "Paciente Agregado Correctamente con ID: " + strconv.Itoa(newID)
	return rp
}

func (d *DB) GetPacientes() ([]models.PacientesModel, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetPacientes)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener pacientes: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return nil, rp
	}
	defer rows.Close()

	var pacientes []models.PacientesModel
	for rows.Next() {
		var paciente models.PacientesModel
		err := rows.Scan(
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
			&paciente.CreatedAt,
		)
		if err != nil {
			rp.Status = 500
			rp.Mensaje = "Error al escanear paciente: " + err.Error()
			utils.CreateLog(rp.Mensaje)
			return nil, rp
		}
		pacientes = append(pacientes, paciente)
	}

	rp.Status = 200
	rp.Mensaje = "Pacientes listados correctamente!"
	return pacientes, rp
}

// UpsertPrecioEspecialidad inserta o actualiza un precio para un paciente y especialidad.
func (d *DB) UpsertPrecioEspecialidad(p models.PrecioEspecialidad) models.Respuesta {
	var rp models.Respuesta

	if p.IDPaciente == 0 || p.Especialidad == "" {
		rp.Status = 400
		rp.Mensaje = "ID de paciente y especialidad son requeridos."
		return rp
	}

	resp, err := d.db.Exec(sqlUpsertPrecioEspecialidad, p.IDPaciente, p.Especialidad, p.Precio)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo guardar el precio por especialidad: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return rp
	}

	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Status = 200
		rp.Mensaje = "Precio por especialidad guardado correctamente."
	} else {
		// En un UPSERT, si no hay cambios, RowsAffected puede ser 0.
		// Esto no es necesariamente un error, significa que el registro ya existía con el mismo precio.
		rp.Status = 200
		rp.Mensaje = "El precio por especialidad no ha cambiado."
	}
	return rp
}

// DelPrecioEspecialidad elimina un precio personalizado para un paciente y especialidad.
func (d *DB) DelPrecioEspecialidad(p models.PrecioEspecialidad) models.Respuesta {
	var rp models.Respuesta

	if p.IDPaciente == 0 || p.Especialidad == "" {
		rp.Status = 400
		rp.Mensaje = "ID de paciente y especialidad son requeridos."
		return rp
	}

	resp, err := d.db.Exec(sqlDelPrecioEspecialidad, p.IDPaciente, p.Especialidad)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo eliminar el precio por especialidad: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return rp
	}

	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Status = 200
		rp.Mensaje = "Precio por especialidad eliminado correctamente."
	} else {
		rp.Status = 404 // Not Found
		rp.Mensaje = "No se encontró el precio por especialidad para eliminar."
	}
	return rp
}

func (d *DB) UpdDoctores(i models.DoctoresModel) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlUpdDoctores, i.Id, i.Nombres, i.Espec, i.Dir, i.Tlf, i.Correo, i.Whatsapp,
		i.Instagram, i.Tasapago)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion de Doctores. " + err.Error()
		utils.CreateLog(err.Error())
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registros Actualizados Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) GetDoctores() ([]models.DoctoresModel, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetDoctores)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Productos Registrados! " + err.Error()
		// utils.CreateLog(err.Error())
		return nil, rp
	}
	defer rows.Close()
	doctores := []models.DoctoresModel{}
	doctor := models.DoctoresModel{}
	for rows.Next() {
		err2 :=
			rows.Scan(
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
		if err2 != nil {
			utils.CreateLog(err2.Error())
		}
		doctores = append(doctores, doctor)
	}
	rp.Status = 10
	rp.Mensaje = "Doctores listado correctamente!"
	return doctores, rp
}

func (d *DB) PostDoctor(i models.DoctoresModel) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlPostDoctor, i.Nombres, i.Espec, i.Dir, i.Tlf,
		i.Correo, i.Whatsapp, i.Instagram, i.Tasapago)
	if err != nil {
		rp.Status = 501
		rp.Mensaje = "No se pudo Agregar la Informacion del Especialista. " + err.Error()
		utils.CreateLog("No se pudo Agregar la Informacion del Especialista. " + err.Error())
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Status = 200
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Especialista Agregado Correctamente"
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) DelDoctor(i models.DoctoresModel) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlDelDoctor, i.Id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Eliminar el Especialista. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Eliminado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se agrego producto!"
	}
	return rp
}

func (d *DB) GetParametros() ([]models.Parametros, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetParametros)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Parametros Registrados! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	parametros := []models.Parametros{}
	parametro := models.Parametros{}
	for rows.Next() {
		err2 :=
			rows.Scan(
				&parametro.Id,
				&parametro.Parametro,
				&parametro.Descripcion,
				&parametro.Valor,
				&parametro.Valores,
				&parametro.Descvalor,
			)
		if err2 != nil {
			utils.CreateLog(err2.Error())
		}
		parametros = append(parametros, parametro)
	}
	rp.Status = 10
	rp.Mensaje = "Parametros listado correctamente!"
	return parametros, rp
}

func (d *DB) GetEmailConfig() ([]models.EmailConfig, models.Respuesta) {
	var rp models.Respuesta
	emailConfig := models.EmailConfig{}
	emailConfigs := []models.EmailConfig{}

	rows, err := d.db.Query(sqlGetEmailConfig)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Parametros Registrados! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	// emailConfigs := []models.EmailConfig{}

	for rows.Next() {
		err2 :=
			rows.Scan(
				&emailConfig.Id,
				&emailConfig.Smtp,
				&emailConfig.Port,
				&emailConfig.Usuario,
				&emailConfig.Clave,
				&emailConfig.Tls,
			)
		if err2 != nil {
			utils.CreateLog(err2.Error())
		}
		emailConfigs = append(emailConfigs, emailConfig)
	}
	rp.Status = 10
	rp.Mensaje = "Parametros listado correctamente!"
	return emailConfigs, rp
}

func (d *DB) SendMail(f models.MailSend) {
	m := mail.NewMessage()

	// Get Email config
	eml, resp := d.GetEmailConfig()
	// utils.CreateLog(resp.Mensaje)

	if resp.Status != 10 {
		utils.CreateLog(resp.Mensaje)
	}
	// utils.CreateLog("smtp " + eml[0].Smtp + "puerto " + strconv.Itoa(eml[0].Port) + "usuario " + eml[0].Usuario + " clave " + eml[0].Clave)
	// Destinatarios
	m.SetHeader("To", f.To)
	m.SetHeader("From", eml[0].Usuario)
	m.SetHeader("Subject", f.Subject)

	// Cuerpo del correo
	m.SetBody("text/plain", f.Body)

	// Adjuntar archivo PDF
	m.Attach(f.Archivo)

	// Configuración del servidor SMTP
	// d := mail.NewDialer("smtp.gmail.com", 587, "omhmre@gmail.com", "kxjs haaz cbfr mdtb")
	dd := mail.NewDialer(eml[0].Smtp, eml[0].Port, eml[0].Usuario, eml[0].Clave)

	// Habilitar SSL
	dd.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Enviar el correo
	if err := dd.DialAndSend(m); err != nil {
		utils.CreateLog(err.Error())
	}
}

func (d *DB) UpdEmailConfig(i models.EmailConfig) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlUpdEmailConfig, i.Id, i.Smtp, i.Port, i.Usuario, i.Clave, i.Tls)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion de Correo. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registros Actualizados Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) AddParametro(i models.Parametros) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlAddParametros, i.Parametro, i.Descripcion, i.Valor, i.Valores, i.Descvalor)
	if err != nil {
		rp.Status = 501
		rp.Mensaje = "No se pudo Agregar la Informacion del Parametro. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Status = 200
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Parametro Agregado Correctamente"
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) UpdParametro(i models.Parametros) models.Respuesta {
	var rp models.Respuesta
	switch i.Valor {
	case 1:
		i.Descvalor = "Comercial"
	case 2:
		i.Descvalor = "Restaurant"
	case 3:
		i.Descvalor = "Laboratorio"
	default:
		i.Descvalor = "Comercial"
	}

	resp, err := d.db.Exec(sqlUpdateParametros, i.Parametro, i.Descripcion, i.Valor, i.Valores, i.Descvalor)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion del Parametro. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registros Actualizados Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) Open() error {
	pg, err := sql.Open("postgres", dbinfo)
	if err != nil {
		utils.CreateLog(err.Error())
		return err
	}
	utils.CreateLog("Conectado a la base de datos " + DB_NAME)
	d.db = pg
	d.Conn = pg // Asignar la conexión también al campo 'Conn'

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

// BackupDatabase creates a backup of the PostgreSQL database
func (d *DB) BackupDatabase() error {
	// Set the backup file name and path
	backupFile := fmt.Sprintf("mangoadmin_db_backup_%s.sql", time.Now().Format("20060102150405"))
	backupPath := "./backups/" + backupFile

	// Create the backups directory if it doesn't exist
	err := os.MkdirAll("./backups", 0755)
	if err != nil {
		return err
	}

	// Execute the pg_dump command
	// cmd := exec.Command("pg_dump", "-U", DB_USER, "--password", DB_PASSWORD, "-d", DB_NAME, "-f", backupPath)
	cmd := exec.Command("pg_dump", "-h", DB_SERVER, "-p", DB_PORT, "-U", DB_USER, "-d", DB_NAME, "-f", backupPath)
	// Set the PGPASSWORD environment variable for the command
	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", DB_PASSWORD))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	// Log the backup output
	utils.CreateLog(string(output))

	return nil
}

func (d *DB) GetInventarios(compacto int) ([]models.Inventario, models.Respuesta) {
	var rp models.Respuesta

	var strSql string
	switch compacto {
	case 0:
		// Inventario Completo
		strSql = sqlGetInventario
	case 1:
		// Inventario Resumido
		strSql = sqlGetInventarioCompacto
	case 2:
		// Inventario Formal para Reportes
		strSql = sqlGetInventarioFormal
	}
	rows, err := d.db.Query(strSql)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Productos Registrados! " + err.Error()
		// utils.CreateLog(err.Error())
		return nil, rp
	}
	defer rows.Close()
	inventarios := []models.Inventario{}
	producto := models.Inventario{}
	for rows.Next() {
		err2 :=
			rows.Scan(
				&producto.Codigo,
				&producto.Nombre,
				&producto.Marca,
				&producto.Unidad,
				&producto.Costo,
				&producto.Costoa,
				&producto.Costopr,
				&producto.Precio1,
				&producto.Precio2,
				&producto.Precio3,
				&producto.Cantidad,
				&producto.Enser,
				&producto.Exento,
				&producto.Clasif,
				&producto.Tipo,
				&producto.Empaque,
				&producto.Cantemp,
				&producto.Pedido,
				&producto.Disponible,
				&producto.Preciom1,
				&producto.Preciom2,
				&producto.Preciom3,
				&producto.Costodolar,
				&producto.Dirfoto,
				&producto.Foto,
				&producto.Descripcion,
				&producto.Codservicio,
				&producto.Preciovar,
				&producto.Compuesto,
				&producto.Mateprima,
				&producto.Global,
				&producto.Cantvar,
				&producto.Espresent,
			)
		if err2 != nil {
			utils.CreateLog(err2.Error())
		}

		items := []models.ItemsInventario{}
		item := models.ItemsInventario{}
		if producto.Compuesto {
			itemsinventario, err := d.db.Query(sqlGetItemsInventario, producto.Codigo)
			if err != nil {
				rp.Status = 502
				rp.Mensaje = "No Hay Productos Registrados! " + err.Error()
				return nil, rp
			}
			defer itemsinventario.Close()
			for itemsinventario.Next() {
				itemsinventario.Scan(
					&item.Codinventario,
					&item.Coditem,
					&item.Nombre,
					&item.Cantidad,
				)
				items = append(items, item)
			}

		} else {
			items = []models.ItemsInventario{}
		}
		producto.Items = items
		presentaciones := []models.PresenInventario{}
		presentacion := models.PresenInventario{}
		if producto.Espresent {
			presInventario, err := d.db.Query(sqlGetPresenInventario, producto.Codigo)
			if err != nil {
				rp.Status = 502
				rp.Mensaje = "No Hay Productos Registrados! " + err.Error()
				return nil, rp
			}
			defer presInventario.Close()
			for presInventario.Next() {
				presInventario.Scan(
					&presentacion.Id,
					&presentacion.Codinv,
					&presentacion.Presentacion,
					&presentacion.Cantidad,
					&presentacion.Precio,
				)
				presentacion.Items = []models.ItemsPresentacion{}
				presentaciones = append(presentaciones, presentacion)
			}
		} else {
			presentaciones = []models.PresenInventario{}
		}
		producto.Presentaciones = presentaciones

		inventarios = append(inventarios, producto)
	}
	rp.Status = 10
	rp.Mensaje = "Inventario listado correctamente!"
	return inventarios, rp
}

func (d *DB) GetInventarioFormal() ([]models.Inventario, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetInventario)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Productos Registrados! " + err.Error()
		// utils.CreateLog(err.Error())
		return nil, rp
	}
	defer rows.Close()
	inventarios := []models.Inventario{}
	producto := models.Inventario{}
	for rows.Next() {
		err2 :=
			rows.Scan(
				&producto.Codigo,
				&producto.Nombre,
				&producto.Unidad,
				&producto.Cantidad,
				&producto.Costo,
				&producto.Precio1,
				&producto.Preciom1,
			)
		if err2 != nil {
			utils.CreateLog(err2.Error())
		}
		inventarios = append(inventarios, producto)
	}
	rp.Status = 10
	rp.Mensaje = "Inventario listado correctamente!"
	return inventarios, rp
}

func (d *DB) AddInventario(i models.Inventario) models.Respuesta {
	var rp models.Respuesta
	// fmt.Println("datos de inventario")
	// fmt.Println(i)
	resp, err := d.db.Exec(sqlAddInventario, i.Codigo, i.Nombre, i.Marca, i.Unidad, i.Costo, i.Costoa, i.Costopr, i.Precio1, i.Precio2,
		i.Precio3, i.Cantidad, i.Enser, i.Exento, i.Clasif, i.Tipo, i.Empaque, i.Cantemp, i.Pedido, i.Disponible, i.Preciom1, i.Preciom2,
		i.Preciom3, i.Costodolar, i.Dirfoto, i.Foto, i.Descripcion, i.Preciovar, i.Compuesto, i.Mateprima, i.Espresent)
	if err != nil {
		rp.Status = 501
		rp.Mensaje = "No se pudo Agregar la Informacion del Producto. " + err.Error()
		utils.CreateLog("No se pudo Agregar la Informacion del Producto. " + err.Error())
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Status = 200
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Producto Agregado Correctamente"
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) AddItemsInventario(i []models.ItemsInventario) models.Respuesta {
	var rp models.Respuesta
	nroItems := int64(0)

	d.db.Exec(sqlDelItemsInventario, i[0].Codinventario)
	if len(i) > 0 {
		for _, item := range i {
			resp, err := d.db.Exec(sqlAddItemsInventario, item.Codinventario, item.Coditem, item.Nombre, item.Cantidad)
			if err != nil {
				rp.Status = 501
				rp.Mensaje = "No se pudo Agregar la Informacion del Producto. " + err.Error()
				return rp
			}

			datos, err1 := resp.RowsAffected()
			if err1 != nil {
				rp.Status = 502
				rp.Mensaje = err1.Error()
			}
			nroItems += datos
		}
	}
	rp.Status = 200
	rp.Mensaje = strconv.FormatInt(nroItems, 10) + " Items Agregados Correctamente"
	return rp
}

func (d *DB) AddPresenInventario(i []models.PresenInventario) models.Respuesta {
	var rp models.Respuesta
	nroItems := int64(0)

	_, err := d.db.Exec(sqlDelPresenInventario, i[0].Codinv)
	if err != nil {
		utils.CreateLog(err.Error())
	}

	var items []models.ItemsPresentacion

	for _, item := range i {
		var miId = 0
		resp, err := d.db.Query(sqlAddPresenInventario, item.Codinv, item.Presentacion, item.Cantidad, item.Precio)
		if err != nil {
			rp.Status = 501
			rp.Mensaje = "No se pudo Agregar la Informacion del Producto. " + err.Error()
			utils.CreateLog("No se pudo Agregar la Informacion del Producto. " + err.Error())
			return rp
		}

		// datos, err1 := resp.RowsAffected()
		// if err1 != nil {
		// 	rp.Status = 502
		// 	rp.Mensaje = err1.Error()
		// 	utils.CreateLog(err1.Error())
		// }
		nroItems += 1

		for resp.Next() {
			resp.Scan(
				&miId,
			)
		}
		items = item.Items
		for _, v := range items {
			_, err2 := d.db.Exec(sqlAddItemsPres, miId, v.Codinv, v.Cantidad)
			if err2 != nil {
				utils.CreateLog(err2.Error())
			}
		}
	}
	rp.Mensaje = strconv.FormatInt(nroItems, 10) + " Items Agregados Correctamente"
	return rp
}

func (d *DB) UpdInventario(i models.Inventario) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlUpdInventario, i.Codigo, i.Nombre, i.Marca, i.Unidad, i.Costo, i.Costoa, i.Costopr,
		i.Precio1, i.Precio2, i.Precio3, i.Cantidad, i.Enser, i.Exento, i.Clasif, i.Tipo, i.Empaque, i.Cantemp, i.Pedido, i.Disponible,
		i.Preciom1, i.Preciom2, i.Preciom3, i.Costodolar, i.Dirfoto, i.Foto, i.Descripcion, i.Preciovar, i.Compuesto, i.Mateprima, i.Espresent)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion del Producto. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registros Actualizados Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
		// fmt.Println(rp.Mensaje)
	}
	return rp
}

func (d *DB) DelInventario(i models.Inventario) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlDelInventario, i.Codigo)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Eliminar la Informacion del Producto. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Eliminado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201 // No se encontró el registro
		rp.Mensaje = "No se encontró el paciente para eliminar."
	}
	return rp
}

func (d *DB) GetInventarioMenu() ([]models.InventarioNombre, models.Respuesta) {
	// var rp models.Respuesta
	var rp models.Respuesta
	rows, err := d.db.Query(sqlGetInventarioNombre)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Productos Registrados! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	inventarios := []models.InventarioNombre{}
	inventario := models.InventarioNombre{}
	for rows.Next() {
		rows.Scan(
			&inventario.Codigo,
			&inventario.Nombre,
			&inventario.Preciom1,
		)
		inventarios = append(inventarios, inventario)
	}
	rp.Status = 10
	rp.Mensaje = "Inventario listado correctamente!"
	return inventarios, rp
}

func (d *DB) GetMenu() ([]models.Menu, models.Respuesta) {
	// var rp models.Respuesta
	var rp models.Respuesta
	rows, err := d.db.Query(sqlGetMenu)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Menu Registrado! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	menues := []models.Menu{}
	menu := models.Menu{}
	for rows.Next() {
		rows.Scan(
			&menu.Codigo,
			&menu.Nombre,
			&menu.Descripcion,
			&menu.Idclase,
			&menu.Precio1,
			&menu.Precio2,
			&menu.Precio3,
			&menu.Cantidad,
			&menu.Preciom1,
			&menu.Preciom2,
			&menu.Preciom3,
			&menu.Dirfoto,
			&menu.Foto,
		)
		menues = append(menues, menu)
	}
	rp.Status = 10
	rp.Mensaje = "Menu listado correctamente!"
	return menues, rp
}

func (d *DB) GetMenuListado() ([]models.ClaseMenu, models.Respuesta) {
	// var rp models.Respuesta
	var rp models.Respuesta
	rows, err := d.db.Query(sqlGetClase)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Menu Registrado! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	clases := []models.ClaseMenu{}
	clase := models.ClaseMenu{}
	for rows.Next() {
		rows.Scan(
			&clase.Id,
			&clase.Nombre,
		)
		rowsMenu, err1 := d.db.Query(sqlGetMenuClase, clase.Id)
		if err1 != nil {
			rp.Status = 502
			rp.Mensaje = "No Hay Menu Registrado! " + err1.Error()
			return nil, rp
		}
		defer rowsMenu.Close()
		menues := []models.Menu{}
		menu := models.Menu{}
		for rowsMenu.Next() {
			rowsMenu.Scan(
				&menu.Codigo,
				&menu.Nombre,
				&menu.Descripcion,
				&menu.Idclase,
				&menu.Precio1,
				&menu.Precio2,
				&menu.Precio3,
				&menu.Cantidad,
				&menu.Preciom1,
				&menu.Preciom2,
				&menu.Preciom3,
				&menu.Dirfoto,
				&menu.Foto,
			)
			menues = append(menues, menu)
		}
		clase.Menu = menues
		clases = append(clases, clase)
	}
	rp.Status = 10
	rp.Mensaje = "Menu listado correctamente!"
	return clases, rp
}

func (d *DB) GetClases() ([]models.Clase, models.Respuesta) {
	var rp models.Respuesta
	rows, err := d.db.Query(sqlGetClase)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Menu Registrado! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	clases := []models.Clase{}
	clase := models.Clase{}
	for rows.Next() {
		rows.Scan(
			&clase.Id,
			&clase.Nombre,
		)
		clases = append(clases, clase)
	}
	rp.Status = 10
	rp.Mensaje = "Menu listado correctamente!"
	return clases, rp
}

func (d *DB) AddClase(i models.Clase) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlAddClase, i.Nombre)
	if err != nil {
		rp.Status = 501
		rp.Mensaje = "No se pudo Agregar la Informacion de Clase de Menu. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Status = 200
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Clase Agregada Correctamente"
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) UpdClase(i models.Clase) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlUpdClase, i.Id, i.Nombre)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion de la Clase Menu. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Actualizado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) DelClase(i models.Clase) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlDelClase, i.Id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Eliminar la Informacion de la Clase Menu. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Eliminado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se agrego la clase!"
	}
	return rp
}

func (d *DB) AddMenu(m models.Menu) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlAddMenu, m.Codigo, m.Nombre, m.Descripcion, m.Idclase, m.Precio1, m.Precio2, m.Precio3, m.Cantidad, m.Preciom1, m.Preciom2, m.Preciom3, m.Dirfoto, m.Foto)
	if err != nil {
		rp.Status = 501
		rp.Mensaje = "No se pudo Agregar la Informacion del Menu. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Status = 200
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Menu Agregado Correctamente"
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) AddAllMenu(m []models.Menu) models.Respuesta {
	var rp models.Respuesta
	_, errDel := d.db.Query(sqlDelAllMenu)
	if errDel != nil {
		rp.Status = 503
		rp.Mensaje = "Ocurrio un Error al intentar eliminar todos los menus: " + errDel.Error()
		return rp
	}
	for _, mn := range m {
		resp, err := d.db.Exec(sqlAddMenu, mn.Codigo, mn.Nombre, mn.Descripcion, mn.Idclase, mn.Precio1, mn.Precio2, mn.Precio3, mn.Cantidad, mn.Preciom1, mn.Preciom2, mn.Preciom3, mn.Dirfoto, mn.Foto)
		if err != nil {
			rp.Status = 501
			rp.Mensaje = "No se pudo Agregar la Informacion del Menu. " + err.Error()
			// fmt.Println(rp.Mensaje)
			return rp
		}
		datos, err1 := resp.RowsAffected()
		if err1 != nil {
			rp.Status = 502
			rp.Mensaje = err1.Error()
		} else if datos > 0 {
			rp.Status = 200
			rp.Mensaje = strconv.FormatInt(datos, 10) + " Menu Agregado Correctamente"
		} else {
			rp.Status = 201
			rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
		}
	}
	return rp
}

func (d *DB) UpdMenu(m models.Menu) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlUpdMenu, m.Codigo, m.Nombre, m.Descripcion, m.Idclase, m.Precio1, m.Precio2, m.Precio3, m.Cantidad, m.Preciom1, m.Preciom2, m.Preciom3, m.Dirfoto, m.Foto)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion del Menu. " + err.Error()
		return rp
	}

	_, err2 := d.db.Exec(sqlUpdInventCod, m.Codigo, m.Nombre, m.Descripcion, m.Precio1, m.Precio2, m.Precio3, m.Preciom1, m.Preciom2, m.Preciom3, m.Dirfoto, m.Foto)
	if err2 != nil {
		rp.Status = 501
		rp.Mensaje = "No se pudo Actualizar la Informacion del Producto en la base de datos. " + err2.Error()
		return rp
	}

	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 503
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Actualizado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) DelMenu(m models.Menu) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlDelMenu, m.Codigo)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Eliminar la Informacion del Producto en el Menu. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Eliminado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se agrego la clase!"
	}
	return rp
}

func (d *DB) DelAllMenu() models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlDelAllMenu)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Eliminar la Informacion del Menu. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Eliminado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se agrego la clase!"
	}
	return rp
}

func (d *DB) GetEmpre() ([]models.Empresa, models.Respuesta) {
	var rp models.Respuesta
	rows, err := d.db.Query(sqlGetEmpre)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Empresa Registrada! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	empresas := []models.Empresa{}
	empresa := models.Empresa{}
	for rows.Next() {
		rows.Scan(
			&empresa.ID,
			&empresa.Rif,
			&empresa.Rasocial,
			&empresa.Dirfisc,
			&empresa.Ciudad,
			&empresa.Estado,
			&empresa.Telf,
			&empresa.Logo,
			&empresa.Comercial,
			&empresa.Slogan,
			&empresa.Iva,
			&empresa.Correo,
			&empresa.Instagram,
			&empresa.Whatsapp,
		)
		empresas = append(empresas, empresa)
	}
	rp.Status = 10
	rp.Mensaje = "Empresas listada correctamente!"
	return empresas, rp
}

func (d *DB) AddEmpre(e models.Empresa) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlAddEmpresa,
		e.ID,
		e.Rif,
		e.Rasocial,
		e.Dirfisc,
		e.Ciudad,
		e.Estado,
		e.Telf,
		e.Logo,
		e.Comercial,
		e.Slogan,
		e.Iva,
		e.Correo,
		e.Instagram,
		e.Whatsapp,
	)
	if err != nil {
		rp.Status = 501
		rp.Mensaje = "No se pudo Agregar la Informacion de Empresa. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 502
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Status = 200
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Empresa Agregada Correctamente"
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) UpdEmpresa(e models.Empresa) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlUpdEmpresa, e.ID, e.Rif, e.Rasocial, e.Dirfisc, e.Ciudad, e.Estado, e.Telf, e.Logo, e.Comercial, e.Slogan, e.Iva, e.Correo,
		e.Instagram, e.Whatsapp)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion de la Empresa. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Actualizado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) DelEmpresa(e models.Empresa) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlDelEmpresa, e.ID)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Eliminar la Informacion de la Empresa. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Eliminado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se agrego la Empresa!"
	}
	return rp
}

func (d *DB) GetMenuCompleto() (models.RespMenuCompleto, models.Respuesta) {
	var rp models.Respuesta
	var rmc models.RespMenuCompleto
	var mc models.MenuCompleto
	miEmpre, errorEmpre := d.GetEmpre()
	if errorEmpre.Status != 10 {
		return rmc, errorEmpre
	}
	mc.MiEmpresa = miEmpre[0]
	rows, err := d.db.Query(sqlGetClase)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Menu Registrado! " + err.Error()
		return rmc, rp
	}
	defer rows.Close()
	clases := []models.ClaseMenu{}
	clase := models.ClaseMenu{}
	for rows.Next() {
		rows.Scan(
			&clase.Id,
			&clase.Nombre,
		)
		rowsMenu, err1 := d.db.Query(sqlGetMenuClase, clase.Id)
		if err1 != nil {
			rp.Status = 502
			rp.Mensaje = "No Hay Menu Registrado! " + err1.Error()
			return rmc, rp
		}
		defer rowsMenu.Close()
		menues := []models.Menu{}
		menu := models.Menu{}
		for rowsMenu.Next() {
			rowsMenu.Scan(
				&menu.Codigo,
				&menu.Nombre,
				&menu.Descripcion,
				&menu.Idclase,
				&menu.Precio1,
				&menu.Precio2,
				&menu.Precio3,
				&menu.Cantidad,
				&menu.Preciom1,
				&menu.Preciom2,
				&menu.Preciom3,
				&menu.Dirfoto,
				&menu.Foto,
			)
			menues = append(menues, menu)
		}
		clase.Menu = menues
		clases = append(clases, clase)
	}

	mc.MiClaseMenu = clases
	rmc.Data = mc
	rp.Status = 10
	rp.Mensaje = "Menu listado correctamente!"
	return rmc, rp
}

func (d *DB) GetPrefacturas() ([]models.Prefactura, models.Respuesta) {
	// var rp models.Respuesta
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetPrefacturas)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Facturas Registradas! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	prefacturas := []models.Prefactura{}
	pf := models.Prefactura{}
	for rows.Next() {
		rows.Scan(
			&pf.Id,
			&pf.Idcliente,
			&pf.Fecha,
			&pf.Subtotal,
			&pf.Dscto,
			&pf.Mototal,
			&pf.Deimp,
			&pf.Tasaimp,
			&pf.Moimp,
			&pf.Moneto,
			&pf.Idvendedor,
			&pf.Idsesion,
		)
		prefacturas = append(prefacturas, pf)
	}
	rp.Status = 10
	rp.Mensaje = "Prefacturas listadas correctamente!"
	return prefacturas, rp
}

func (d *DB) GetPrefactura(p models.Prefactura) ([]models.Prefactura, models.Respuesta) {
	var rp models.Respuesta

	// utils.CreateLog("funcion Get prefactura")
	// utils.CreateLog(p.Id)
	rows, err := d.db.Query(sqlGetPrefactura, p.Id)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Facturas Registradas! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	prefacturas := []models.Prefactura{}
	prefactura := models.Prefactura{}
	for rows.Next() {
		rows.Scan(
			&prefactura.Id,
			&prefactura.Idcliente,
			&prefactura.Cliente,
			&prefactura.Fecha,
			&prefactura.Subtotal,
			&prefactura.Dscto,
			&prefactura.Mototal,
			&prefactura.Deimp,
			&prefactura.Tasaimp,
			&prefactura.Moimp,
			&prefactura.Moneto,
			&prefactura.Monetodiv,
			&prefactura.Idvendedor,
			&prefactura.Idsesion,
			&prefactura.Idmesa,
			&prefactura.Idmesonero,
			&prefactura.Pagado,
			&prefactura.Porpagar,
			&prefactura.Cambio,
			&prefactura.Tasadiv,
			&prefactura.Cxcbs,
			&prefactura.Cxcdiv,
		)
		losItems, _ := d.db.Query(sqlGetItemsFacturas, p.Id)
		defer losItems.Close()
		misItems := []models.Item{}
		mIt := models.Item{}
		for losItems.Next() {
			losItems.Scan(
				&mIt.Idprefact,
				&mIt.Codprod,
				&mIt.Producto,
				&mIt.Cant,
				&mIt.Precio,
				&mIt.Subtotal,
				&mIt.Descuento,
				&mIt.Neto,
				&mIt.Descripcion,
				&mIt.Cantmp,
				&mIt.Cantpres,
			)
			misItems = append(misItems, mIt)
		}
		prefactura.Items = misItems
		prefacturas = append(prefacturas, prefactura)
	}
	rp.Status = 10
	rp.Mensaje = "Prefactura listada correctamente!"
	return prefacturas, rp
}

func (d *DB) PostPrefactura(p models.Prefactura) (models.Prefactura, models.Respuesta) {
	var rp models.Respuesta

	var items = p.Items

	if p.Id == -1 {
		var miId = 0
		resp, err := d.db.Query(sqlPostPrefacturaNueva, p.Idcliente, p.Fecha, p.Diasvence, p.Vence, p.Subtotal, p.Dscto, p.Mototal, p.Deimp, p.Tasaimp,
			p.Moimp, p.Moneto, p.Idvendedor, p.Idsesion, p.Idmesa, p.Idmesonero, p.Porpagar, p.Tasadiv, p.Monetodiv, p.Condiciones)
		if err != nil {
			rp.Status = 501
			rp.Mensaje = "No se pudo Agregar la Informacion de Factura. " + err.Error()
			utils.CreateLog("No se pudo Agregar la Informacion de Factura. " + err.Error())
			return p, rp
		}

		for resp.Next() {
			resp.Scan(
				&miId,
			)
		}
		items = p.Items

		for _, v := range items {
			_, err2 := d.db.Exec(sqlAddItems, miId, v.Codprod, v.Cant, v.Precio, v.Subtotal, v.Descuento, v.Neto, v.Descripcion, v.Cantmp, v.Cantpres, v.Iditempres, v.Producto)
			if err2 != nil {
				utils.CreateLog(err2.Error())
			}
		}
		_, err1 := d.db.Exec(sqlAbrirMesa, 1, p.Mototal, miId, p.Idcliente, p.Idmesonero, p.Idmesa)
		if err1 != nil {
			utils.CreateLog(err1.Error())
		}
		p.Id = miId
	} else {
		_, err := d.db.Exec(sqlPostPrefactura, p.Id, p.Idcliente, p.Fecha, p.Subtotal, p.Dscto, p.Mototal, p.Deimp, p.Tasaimp,
			p.Moimp, p.Moneto, p.Idvendedor, p.Idsesion, p.Idmesa, p.Idmesonero, p.Porpagar, p.Tasadiv, p.Monetodiv)
		if err != nil {
			rp.Status = 501
			rp.Mensaje = "No se pudo Agregar la Informacion de Factura. " + err.Error()
			utils.CreateLog("No se pudo Agregar la Informacion de Factura. " + err.Error())
			return p, rp
		}
		items = p.Items
		_, err3 := d.db.Exec(sqlDelItems, p.Id)
		if err3 != nil {
			utils.CreateLog("error al eliminar item prefactura " + err3.Error())
		}
		for _, v := range items {
			_, err2 := d.db.Exec(sqlAddItems, p.Id, v.Codprod, v.Cant, v.Precio, v.Subtotal, v.Descuento, v.Neto, v.Descripcion, v.Cantmp, v.Cantpres, v.Iditempres)
			if err2 != nil {
				utils.CreateLog(err2.Error())
			}
		}

		_, err4 := d.db.Exec(sqlActualizarMesa, p.Mototal, p.Idcliente, p.Idmesonero, p.Idmesa)
		if err4 != nil {
			utils.CreateLog(err4.Error())
		}
	}
	return p, rp
}

func (d *DB) UpdPrefactura(p models.Prefactura) models.Respuesta {
	var rp models.Respuesta

	resp, err := d.db.Exec(sqlPutPrefactura, p.Id, p.Idcliente, p.Fecha, p.Subtotal, p.Dscto, p.Mototal, p.Deimp, p.Tasaimp, p.Moimp, p.Moneto, p.Idvendedor, p.Idsesion)
	if err != nil {
		rp.Status = 501
		rp.Mensaje = "No se pudo agregar la Informacion de Factura. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Actualizado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) GetClientes() ([]models.Cliente, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetClientes)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Clientes Registrados! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	clientes := []models.Cliente{}
	cliente := models.Cliente{}
	for rows.Next() {
		rows.Scan(
			&cliente.Id,
			&cliente.Tipo,
			&cliente.Nombre,
			&cliente.Rif,
			&cliente.Dirfiscal,
			&cliente.Ciudad,
			&cliente.Estado,
			&cliente.Telf,
			&cliente.Correo,
			&cliente.Twitter,
			&cliente.Facebook,
			&cliente.Whatsapp,
			&cliente.Instagram,
			&cliente.Status,
			&cliente.Clasif,
			&cliente.Dscto,
			&cliente.Cred,
			&cliente.Diascr,
			&cliente.Cxcbs,
			&cliente.Persconta,
			&cliente.Tlfconta,
			&cliente.Codvend,
			&cliente.Cxcdiv,
		)

		items := []models.ItemsCxcClientes{}
		item := models.ItemsCxcClientes{}
		itemsCxc, err := d.db.Query(sqlGetItemsCxcCliente, cliente.Id)
		if err != nil {
			rp.Status = 502
			rp.Mensaje = "No Hay Clientes Registrados! " + err.Error()
			return nil, rp
		}
		defer itemsCxc.Close()
		for itemsCxc.Next() {
			itemsCxc.Scan(
				&item.Id,
				&item.Idfact,
				&item.Fecha,
				&item.Montobs,
				&item.Cobradobs,
				&item.Saldobs,
				&item.Montodiv,
				&item.Cobradodiv,
				&item.Saldodiv,
			)
			items = append(items, item)
		}
		cliente.Items = items
		clientes = append(clientes, cliente)
	}
	rp.Status = 10
	rp.Mensaje = "Clientes listads correctamente!"
	return clientes, rp
}

func (d *DB) UpdCliente(e models.Cliente) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlUpdCliente,
		e.Id,
		e.Tipo,
		e.Nombre,
		e.Rif,
		e.Dirfiscal,
		e.Ciudad,
		e.Estado,
		e.Telf,
		e.Correo,
		e.Twitter,
		e.Facebook,
		e.Whatsapp,
		e.Instagram,
		e.Status,
		e.Clasif,
		e.Dscto,
		e.Cred,
		e.Diascr,
		e.Cxcbs,
		e.Persconta,
		e.Tlfconta,
		e.Codvend,
		e.Cxcdiv,
	)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion del Cliente. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Actualizado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) DelCliente(e models.Cliente) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlDelCliente, e.Id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Eliminar Cliente. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Eliminado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro el cliente!"
	}
	return rp
}

func (d *DB) AddCliente(e models.Cliente) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlPostCliente,
		e.Id,
		e.Tipo,
		e.Nombre,
		e.Rif,
		e.Dirfiscal,
		e.Ciudad,
		e.Estado,
		e.Telf,
		e.Correo,
		e.Twitter,
		e.Facebook,
		e.Whatsapp,
		e.Instagram,
		e.Status,
		e.Clasif,
		e.Dscto,
		e.Cred,
		e.Diascr,
		e.Cxcbs,
		e.Persconta,
		e.Tlfconta,
		e.Codvend,
		e.Cxcdiv,
	)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo incluir el Cliente. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Actualizado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) GetMesas() ([]models.Mesas, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query("Select * from empre001.mesas order by id")
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Mesas Registrados! " + err.Error()
		return nil, rp
	}

	defer rows.Close()
	mesas := []models.Mesas{}
	mesa := models.Mesas{}
	for rows.Next() {
		rows.Scan(
			&mesa.Id,
			&mesa.Nombre,
			&mesa.Subtotal,
			&mesa.Abierta,
			&mesa.Idprefactura,
			&mesa.Idcliente,
			&mesa.Cliente,
			&mesa.Idmesonero,
			&mesa.Mesonero,
			&mesa.Inicio,
			&mesa.Fin,
		)
		mesas = append(mesas, mesa)
	}
	rp.Status = 10
	rp.Mensaje = "Mesas listadas correctamente!"
	return mesas, rp
}

func (d *DB) AddMesa(e models.Mesas) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlAddMesa,
		e.Id,
		e.Nombre,
	)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo incluir Mesa. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Actualizado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) UpdMesa(e models.Mesas) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlUpdMesa,
		e.Id,
		e.Nombre,
	)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion de la mesa. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Actualizado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) DelMesa(e models.Mesas) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlDelMesa, e.Id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Eliminar Mesa. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Eliminado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro el cliente!"
	}
	return rp
}

func (d *DB) GetMesoneros() ([]models.Mesoneros, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query("Select * from empre001.mesoneros")
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Mesas Registrados! " + err.Error()
		return nil, rp
	}

	defer rows.Close()
	mesoneros := []models.Mesoneros{}
	mesonero := models.Mesoneros{}
	for rows.Next() {
		rows.Scan(
			&mesonero.Id,
			&mesonero.Nombre,
			&mesonero.Direccion,
			&mesonero.Telefono,
		)
		mesoneros = append(mesoneros, mesonero)
	}
	rp.Status = 10
	rp.Mensaje = "Mesas listadas correctamente!"
	return mesoneros, rp
}

func (d *DB) AddMesonero(e models.Mesoneros) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlAddMesonero, e.Id, e.Nombre, e.Direccion, e.Telefono)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo incluir Mesonero. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Actualizado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) UpdMesonero(e models.Mesoneros) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlUpdMesonero, e.Id, e.Nombre, e.Direccion, e.Telefono)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion del mesonero. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Actualizado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) DelMesonero(e models.Mesoneros) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlDelMesonero, e.Id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Eliminar Mesonero. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Eliminado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro el cliente!"
	}
	return rp
}

func (d *DB) AbrirMesa(e models.Mesas) models.Respuesta {
	// utils.CreateLog("funcion abrir mesa")
	var rp models.Respuesta
	// 0: cerrar mesa, 1:abrir mesa
	var strSql = ""
	if e.Abierta == 0 {
		strSql = sqlCerrarMesa
	} else {
		strSql = sqlAbrirMesa
	}
	resp, err := d.db.Exec(strSql, e.Id, e.Idprefactura, e.Idcliente, e.Cliente, e.Idmesonero, e.Mesonero)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo procesar la Informacion de la Mesa. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Actualizado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) GetInstrumentos() ([]models.InstrumentosPagos, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetInstrumentos)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Instrumentos de Pagos Registrados! " + err.Error()
		return nil, rp
	}

	defer rows.Close()
	instrumentos := []models.InstrumentosPagos{}
	instrumento := models.InstrumentosPagos{}
	for rows.Next() {
		rows.Scan(
			&instrumento.Id,
			&instrumento.Descripcion,
			&instrumento.Tasa,
			&instrumento.Simbolo,
		)
		instrumentos = append(instrumentos, instrumento)
	}
	rp.Status = 10
	rp.Mensaje = "Instrumentos de Pagos listados correctamente!"
	return instrumentos, rp
}

func (d *DB) CleanMesa(e models.Mesas) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlCerrarMesa, e.Id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion de la mesa. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Mesa Actualizada Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) PostFactura(p models.IdFactura) models.Respuesta {
	var rp models.Respuesta

	pagadoStr := strconv.FormatFloat(p.Pagado, 'f', 2, 64)
	rows, err := d.db.Query(sqlPostFactura, p.Id, pagadoStr)
	if err != nil {
		utils.CreateLog("No se pudo Agregar la Informacion de Factura. " + err.Error())
		rp.Status = 501
		rp.Mensaje = "No se pudo Agregar la Informacion de Factura. " + err.Error()
		return rp
	}

	defer rows.Close()
	var miId = 0
	for rows.Next() {
		rows.Scan(
			&miId,
		)
	}
	_, errDet := d.db.Exec(sqlPostDetFactura, miId, p.Id)
	if errDet != nil {
		utils.CreateLog("No se pudo Agregar items de la Factura. " + errDet.Error())
		rp.Status = 503
		rp.Mensaje += "Error al agregar detalle: " + errDet.Error()
	}

	// procesando pagos
	var pg = p.Pagos
	var itemsPagos = pg.Items
	var pagoId = 0

	pago, err := d.db.Query(sqlPostPagos, pg.Idcliente, pg.Monto, pg.Dscto, pg.Total, pg.Idsesion)
	if err != nil {
		rp.Status = 501
		rp.Mensaje = "No se pudo Agregar la Informacion de Pago. " + err.Error()
		return rp
	}
	defer pago.Close()
	for pago.Next() {
		pago.Scan(
			&pagoId,
		)
	}
	for _, v := range itemsPagos {
		_, err2 := d.db.Exec(sqlAddItemsPago, pagoId, v.Codpago, v.Depago, v.Cant, v.Tasa, v.Subtotal, miId)
		if err2 != nil {
			utils.CreateLog(err2.Error())
		}
	}
	// fin proceso pagos

	rp.Status = 200
	rp.Mensaje = "Factura agregada "

	return rp
}

func (d *DB) PostNotaEntrega(p models.IdFactura) models.Respuesta {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlPostEntrega, p.Id, p.Pagado)
	if err != nil {
		utils.CreateLog("No se pudo Agregar la Informacion de la Nota de Entrega. " + err.Error())
		rp.Status = 501
		rp.Mensaje = "No se pudo Agregar la Informacion de la Nota de Entrega. " + err.Error()
		return rp
	}

	defer rows.Close()
	var miId = 0
	for rows.Next() {
		rows.Scan(
			&miId,
		)
	}
	_, errDet := d.db.Exec(sqlPostDetEntrega, miId, p.Id)
	if errDet != nil {
		utils.CreateLog("No se pudo Agregar la Informacion de Factura. " + errDet.Error())
		rp.Status = 503
		rp.Mensaje += "Error al agregar detalle: " + errDet.Error()
	}

	// procesando pagos
	// var pg = p.Pagos
	// var itemsPagos = pg.Items
	// var pagoId = 0

	// pago, err := d.db.Query(sqlPostPagos, pg.Idcliente, pg.Monto, pg.Dscto, pg.Total, pg.Idsesion)
	// if err != nil {
	// 	rp.Status = 501
	// 	rp.Mensaje = "No se pudo Agregar la Informacion de Pago. " + err.Error()
	// 	return rp
	// }
	// defer pago.Close()
	// for pago.Next() {
	// 	pago.Scan(
	// 		&pagoId,
	// 	)
	// }
	// for _, v := range itemsPagos {
	// 	_, err2 := d.db.Exec(sqlAddItemsPago, pagoId, v.Codpago, v.Depago, v.Cant, v.Tasa, v.Subtotal, miId)
	// 	if err2 != nil {
	// 		utils.CreateLog(err2.Error())
	// 	}
	// }
	// fin proceso pagos

	rp.Status = miId

	rp.Mensaje = "Nota de Entrega nro " + strconv.Itoa(miId) + " Agregada"

	return rp
}

func (d *DB) PostPresupuesto(p models.Id) models.Respuesta {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlPostPresupuesto, p.Id)
	if err != nil {
		utils.CreateLog("No se pudo Agregar la Informacion del Presupuesto. " + err.Error())
		rp.Status = 501
		rp.Mensaje = "No se pudo Agregar la Informacion del Presupuesto. " + err.Error()
		return rp
	}

	defer rows.Close()
	var miId = 0
	for rows.Next() {
		rows.Scan(
			&miId,
		)
	}
	_, errDet := d.db.Exec(sqlPostDetPresupuesto, miId, p.Id)
	if errDet != nil {
		utils.CreateLog("No se pudo Agregar Items del Presupuesto. " + errDet.Error())
		rp.Status = 503
		rp.Mensaje += "Error al agregar item: " + errDet.Error()
	}
	rp.Status = miId
	rp.Mensaje = "Presupuesto nro " + strconv.Itoa(miId) + " agregado"
	return rp
}

func (d *DB) GetFacturas() ([]models.Factura, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetFacturas + " order by a.id desc")
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Facturas Registradas! " + err.Error()
		utils.CreateLog("No Hay Facturas Registradas! " + err.Error())
		return nil, rp
	}
	defer rows.Close()
	facturas := []models.Factura{}
	factura := models.Factura{}
	for rows.Next() {
		errFact := rows.Scan(
			&factura.Id,
			&factura.Idcliente,
			&factura.Cliente,
			&factura.Dirfiscal,
			&factura.Rif,
			&factura.Persconta,
			&factura.Tlfconta,
			&factura.Fecha,
			&factura.Subtotal,
			&factura.Dscto,
			&factura.Mototal,
			&factura.Deimp,
			&factura.Tasaimp,
			&factura.Moimp,
			&factura.Moneto,
			&factura.Tasadiv,
			&factura.Monetodiv,
			&factura.Idvendedor,
			&factura.Vendedor,
			&factura.Idsesion,
			&factura.Idmesa,
			&factura.Idmesonero,
			&factura.Pagado,
			&factura.Porpagar,
			&factura.Cambio,
			&factura.Cxcbs,
			&factura.Cxcdiv,
			&factura.Valido,
			&factura.Esdivisa,
		)
		if errFact != nil {
			utils.CreateLog("Error en Factura! " + errFact.Error())
		}
		detFacts := []models.ItemsFactura{}
		detFact := models.ItemsFactura{}
		rowsDet, errDet := d.db.Query(sqlGetDetFact, factura.Id)
		if errDet != nil {
			utils.CreateLog("No Hay Facturas Registradas! " + errDet.Error())
			// return nil, rp
		}
		defer rowsDet.Close()
		for rowsDet.Next() {
			// idfact, codprod, deprod, cant, precio, subtotal, descuento, neto, descripcion, subtotaldiv, cantmp
			rowsDet.Scan(
				&detFact.Idfact,
				&detFact.Codprod,
				&detFact.Producto,
				&detFact.Cant,
				&detFact.Precio,
				&detFact.Subtotal,
				&detFact.Descuento,
				&detFact.Neto,
				&detFact.Descripcion,
				&detFact.Cantmp,
				&detFact.Iditempres,
			)
			detFacts = append(detFacts, detFact)
		}
		// utils.CreateLog(factura.Id)
		factura.Items = detFacts

		detPagos := []models.DetPago{}
		detPago := models.DetPago{}
		rowsPagos, errPagos := d.db.Query(sqlGetDetPagosFact, factura.Id)
		if errPagos != nil {
			utils.CreateLog("No Hay Facturas Registradas! " + errPagos.Error())
		}
		defer rowsPagos.Close()
		for rowsPagos.Next() {
			// d.id, d.idpago, d.idinstpago, i.descripcion, d.comenta, d.monto, d.tasa, d.total
			rowsPagos.Scan(
				&detPago.Id,
				&detPago.Idpago,
				&detPago.Idinstpago,
				&detPago.Descripcion,
				&detPago.Comenta,
				&detPago.Monto,
				&detPago.Tasa,
				&detPago.Total,
				&detPago.Idfact,
			)
			detPagos = append(detPagos, detPago)
		}
		factura.Pagos = detPagos
		facturas = append(facturas, factura)
	}
	rp.Status = 10
	rp.Mensaje = "Facturas listadas correctamente!"
	return facturas, rp
}

func (d *DB) GetPresupuestos() ([]models.Presupuestos, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetPresupuestos + " order by a.id desc")
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Presupuestos Registrados! " + err.Error()
		utils.CreateLog("No Hay Presupuestos Registrados! " + err.Error())
		return nil, rp
	}
	defer rows.Close()
	facturas := []models.Presupuestos{}
	factura := models.Presupuestos{}
	for rows.Next() {
		errFact := rows.Scan(
			&factura.Id,
			&factura.Idcliente,
			&factura.Cliente,
			&factura.Dirfiscal,
			&factura.Rif,
			&factura.Persconta,
			&factura.Tlfconta,
			&factura.Fecha,
			&factura.Diasvence,
			&factura.Vence,
			&factura.Subtotal,
			&factura.Dscto,
			&factura.Mototal,
			&factura.Deimp,
			&factura.Tasaimp,
			&factura.Moimp,
			&factura.Moneto,
			&factura.Tasadiv,
			&factura.Monetodiv,
			&factura.Idvendedor,
			&factura.Vendedor,
			&factura.Idsesion,
			&factura.Idmesa,
			&factura.Idmesonero,
			&factura.Condiciones,
		)
		if errFact != nil {
			utils.CreateLog("Error en Presupuestos! " + errFact.Error())
		}
		detFacts := []models.ItemsPresupuestos{}
		detFact := models.ItemsPresupuestos{}
		rowsDet, errDet := d.db.Query(sqlGetDetPre, factura.Id)
		if errDet != nil {
			utils.CreateLog("No Hay Presupuestos Registrados! " + errDet.Error())
			// return nil, rp
		}
		defer rowsDet.Close()
		for rowsDet.Next() {
			// idfact, codprod, deprod, cant, precio, subtotal, descuento, neto, descripcion, subtotaldiv, cantmp
			rowsDet.Scan(
				&detFact.Idpre,
				&detFact.Codprod,
				&detFact.Producto,
				&detFact.Cant,
				&detFact.Precio,
				&detFact.Subtotal,
				&detFact.Descuento,
				&detFact.Neto,
				&detFact.Descripcion,
				&detFact.Cantmp,
				&detFact.Iditempres,
			)
			detFacts = append(detFacts, detFact)
		}
		// utils.CreateLog(factura.Id)
		factura.Items = detFacts
		facturas = append(facturas, factura)
	}
	rp.Status = 10
	rp.Mensaje = "Presupuestos listados correctamente!"
	return facturas, rp
}

func (d *DB) GetPresupuestoId(p models.Presupuestos) (models.Presupuestos, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetPresupuestos+sqlWherePresupuestos, p.Id)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Presupuestos Registrados! " + err.Error()
		utils.CreateLog("No Hay Presupuestos Registrados! " + err.Error())
		return p, rp
	}
	defer rows.Close()
	// facturas := []models.Presupuestos{}
	factura := models.Presupuestos{}
	for rows.Next() {
		errFact := rows.Scan(
			&factura.Id,
			&factura.Idcliente,
			&factura.Cliente,
			&factura.Dirfiscal,
			&factura.Rif,
			&factura.Persconta,
			&factura.Tlfconta,
			&factura.Fecha,
			&factura.Diasvence,
			&factura.Vence,
			&factura.Subtotal,
			&factura.Dscto,
			&factura.Mototal,
			&factura.Deimp,
			&factura.Tasaimp,
			&factura.Moimp,
			&factura.Moneto,
			&factura.Tasadiv,
			&factura.Monetodiv,
			&factura.Idvendedor,
			&factura.Vendedor,
			&factura.Idsesion,
			&factura.Idmesa,
			&factura.Idmesonero,
			&factura.Condiciones,
		)
		if errFact != nil {
			utils.CreateLog("Error en Presupuestos! " + errFact.Error())
		}
		detFacts := []models.ItemsPresupuestos{}
		detFact := models.ItemsPresupuestos{}
		rowsDet, errDet := d.db.Query(sqlGetDetPre, factura.Id)
		if errDet != nil {
			utils.CreateLog("No Hay Presupuestos Registrados! " + errDet.Error())
			// return nil, rp
		}
		defer rowsDet.Close()
		for rowsDet.Next() {
			// idfact, codprod, deprod, cant, precio, subtotal, descuento, neto, descripcion, subtotaldiv, cantmp
			rowsDet.Scan(
				&detFact.Idpre,
				&detFact.Codprod,
				&detFact.Producto,
				&detFact.Cant,
				&detFact.Precio,
				&detFact.Subtotal,
				&detFact.Descuento,
				&detFact.Neto,
				&detFact.Descripcion,
				&detFact.Cantmp,
				&detFact.Iditempres,
			)
			detFacts = append(detFacts, detFact)
		}
		// utils.CreateLog(factura.Id)
		factura.Items = detFacts
		// facturas = append(facturas, factura)
	}
	rp.Status = 10
	rp.Mensaje = "Presupuestos listados correctamente!"
	return factura, rp
}

func (d *DB) GetNotasEntrega() ([]models.Factura, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetNotasEntrega + " order by a.id desc")
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Notas de Entrega Registradas! " + err.Error()
		utils.CreateLog("No Hay Notas de Entrega Registradas! " + err.Error())
		return nil, rp
	}
	defer rows.Close()
	facturas := []models.Factura{}
	factura := models.Factura{}
	for rows.Next() {
		errFact := rows.Scan(
			&factura.Id,
			&factura.Idcliente,
			&factura.Cliente,
			&factura.Dirfiscal,
			&factura.Rif,
			&factura.Persconta,
			&factura.Tlfconta,
			&factura.Fecha,
			&factura.Subtotal,
			&factura.Dscto,
			&factura.Mototal,
			&factura.Deimp,
			&factura.Tasaimp,
			&factura.Moimp,
			&factura.Moneto,
			&factura.Tasadiv,
			&factura.Monetodiv,
			&factura.Idvendedor,
			&factura.Vendedor,
			&factura.Idsesion,
			&factura.Idmesa,
			&factura.Idmesonero,
			&factura.Pagado,
			&factura.Porpagar,
			&factura.Cambio,
			&factura.Cxcbs,
			&factura.Cxcdiv,
			&factura.Valido,
			&factura.Esdivisa,
		)
		if errFact != nil {
			utils.CreateLog("Error en Factura! " + errFact.Error())
		}
		detFacts := []models.ItemsFactura{}
		detFact := models.ItemsFactura{}
		rowsDet, errDet := d.db.Query(sqlGetDetNotasEnt, factura.Id)
		if errDet != nil {
			utils.CreateLog("No Hay Notas de Entrega Registradas! " + errDet.Error())
			// return nil, rp
		}
		defer rowsDet.Close()
		for rowsDet.Next() {
			// idfact, codprod, deprod, cant, precio, subtotal, descuento, neto, descripcion, subtotaldiv, cantmp
			rowsDet.Scan(
				&detFact.Idfact,
				&detFact.Codprod,
				&detFact.Producto,
				&detFact.Cant,
				&detFact.Precio,
				&detFact.Subtotal,
				&detFact.Descuento,
				&detFact.Neto,
				&detFact.Descripcion,
				&detFact.Cantmp,
				&detFact.Iditempres,
			)
			detFacts = append(detFacts, detFact)
		}
		// utils.CreateLog(factura.Id)
		factura.Items = detFacts

		// detPagos := []models.DetPago{}
		// detPago := models.DetPago{}
		// rowsPagos, errPagos := d.db.Query(sqlGetDetPagosFact, factura.Id)
		// if errPagos != nil {
		// 	utils.CreateLog("No Hay Facturas Registradas! " + errPagos.Error())
		// }
		// defer rowsPagos.Close()
		// for rowsPagos.Next() {
		// 	// d.id, d.idpago, d.idinstpago, i.descripcion, d.comenta, d.monto, d.tasa, d.total
		// 	rowsPagos.Scan(
		// 		&detPago.Id,
		// 		&detPago.Idpago,
		// 		&detPago.Idinstpago,
		// 		&detPago.Descripcion,
		// 		&detPago.Comenta,
		// 		&detPago.Monto,
		// 		&detPago.Tasa,
		// 		&detPago.Total,
		// 		&detPago.Idfact,
		// 	)
		// 	detPagos = append(detPagos, detPago)
		// }
		// factura.Pagos = detPagos
		facturas = append(facturas, factura)
	}
	rp.Status = 10
	rp.Mensaje = "Notas de Entrega listadas correctamente!"
	return facturas, rp
}

func (d *DB) GetVentas(f models.Fechas) ([]models.VentasResumen, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetFacturas+sqlGetVentasFecha, f.Desde, f.Hasta)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Facturas Registradas! " + err.Error()
		utils.CreateLog("No Hay Facturas Registradas! " + err.Error())
		return nil, rp
	}
	defer rows.Close()
	facturas := []models.VentasResumen{}
	factura := models.VentasResumen{}
	for rows.Next() {
		errFact := rows.Scan(
			&factura.Id,
			&factura.Idcliente,
			&factura.Cliente,
			&factura.Dirfiscal,
			&factura.Persconta,
			&factura.Tlfconta,
			&factura.Rif,
			&factura.Fecha,
			&factura.Subtotal,
			&factura.Dscto,
			&factura.Mototal,
			&factura.Deimp,
			&factura.Tasaimp,
			&factura.Moimp,
			&factura.Moneto,
			&factura.Tasadiv,
			&factura.Monetodiv,
			&factura.Idvendedor,
			&factura.Vendedor,
			&factura.Idsesion,
			&factura.Idmesa,
			&factura.Idmesonero,
			&factura.Pagado,
			&factura.Porpagar,
			&factura.Cambio,
			&factura.Cxcbs,
			&factura.Cxcdiv,
			&factura.Valido,
			&factura.Esdivisa,
		)
		if errFact != nil {
			utils.CreateLog("Error en Factura! " + errFact.Error())
		}
		facturas = append(facturas, factura)
	}
	rp.Status = 10
	rp.Mensaje = "Facturas listadas correctamente!"
	return facturas, rp
}

func (d *DB) GetFactura(p models.Id) ([]models.Factura, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetFactura, p.Id)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Facturas Registradas! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	facturas := []models.Factura{}
	factura := models.Factura{}
	for rows.Next() {
		rows.Scan(
			&factura.Id,
			&factura.Idcliente,
			&factura.Fecha,
			&factura.Subtotal,
			&factura.Dscto,
			&factura.Mototal,
			&factura.Deimp,
			&factura.Tasaimp,
			&factura.Moimp,
			&factura.Moneto,
			&factura.Idvendedor,
			&factura.Idsesion,
			&factura.Pagado,
			&factura.Porpagar,
			&factura.Cambio,
		)
		facturas = append(facturas, factura)
	}
	rp.Status = 10
	rp.Mensaje = "Facturas listadas correctamente!"
	return facturas, rp
}

func (d *DB) UpdAnularFactura(i models.Factura) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlAnularFactura, i.Id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion de la Factura. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		_, err2 := d.db.Exec(sqlAnularInventario, i.Id)
		if err2 != nil {
			utils.CreateLog(err2.Error())
		}
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Actualizado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) PostVentasFactura(datos DT) ([]models.ResVentasDia, models.Respuesta) {
	var rp models.Respuesta
	rows, err := d.db.Query(sqlGetVentasFactura, datos.Desde, datos.Hasta)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Facturas Registradas! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	ventas := []models.ResVentasDia{}
	venta := models.ResVentasDia{}
	for rows.Next() {
		rows.Scan(
			&venta.Cant,
			&venta.Monto,
			&venta.Descuento,
			&venta.Subtotal,
			&venta.Subtotaldivisa,
		)
		ventas = append(ventas, venta)
	}
	rp.Status = 10
	rp.Mensaje = "Ventas listadas correctamente!"
	return ventas, rp
}

func (d *DB) PostVentasProductos(datos DT) ([]models.ResVentasProductos, models.Respuesta) {
	var rp models.Respuesta
	rows, err := d.db.Query(sqlGetVentasProductos, datos.Desde, datos.Hasta)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Facturas Registradas! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	ventas := []models.ResVentasProductos{}
	venta := models.ResVentasProductos{}
	for rows.Next() {
		rows.Scan(
			&venta.CodProd,
			&venta.Producto,
			&venta.Cantidad,
			&venta.Subtotal,
			&venta.Subtotaldivisa,
			&venta.Cantmp,
			&venta.Cantpres,
		)
		ventas = append(ventas, venta)
	}
	rp.Status = 10
	rp.Mensaje = "Ventas listadas correctamente!"
	return ventas, rp
}

func (d *DB) UpdDivisa(i models.Divisas) models.Respuesta {
	var rp models.Respuesta
	resp, err := d.db.Exec(sqlUpdDivisa, i.Tasabs, i.Fecha)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion de la Clase Menu. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Actualizado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	_, err2 := d.db.Exec(sqlUpdDivisaInst, i.Tasabs)
	if err2 != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion de la Clase Menu. " + err2.Error()
		return rp
	}
	return rp
}

func (d *DB) PostPagos(p models.Pagos) models.Respuesta {
	var rp models.Respuesta

	var items = p.Items
	var miId = 0

	pago, err := d.db.Query(sqlPostPagos, p.Idcliente, p.Monto, p.Dscto, p.Total, p.Idsesion)
	if err != nil {
		rp.Status = 501
		rp.Mensaje = "No se pudo Agregar la Informacion de Pago. " + err.Error()
		return rp
	}
	defer pago.Close()
	for pago.Next() {
		pago.Scan(
			&miId,
		)
	}
	for _, v := range items {
		_, err2 := d.db.Exec(sqlAddItemsPago, miId, v.Codpago, v.Depago, v.Cant, v.Tasa, v.Subtotal, v.Idfactura)
		if err2 != nil {
			utils.CreateLog(err2.Error())
		}
	}
	return rp
}

func (d *DB) GetDetPagosFecha(p models.Fechas) ([]models.DetPago, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetDetPagos, p.Desde, p.Hasta)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Pagos Registrados! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	pagos := []models.DetPago{}
	pago := models.DetPago{}
	for rows.Next() {
		rows.Scan(
			&pago.Id,
			&pago.Idpago,
			&pago.Idinstpago,
			&pago.Descripcion,
			&pago.Comenta,
			&pago.Monto,
			&pago.Tasa,
			&pago.Total,
			&pago.Idfact,
		)
		pagos = append(pagos, pago)
	}
	rp.Status = 10
	rp.Mensaje = "Pagos listados correctamente!"
	return pagos, rp
}

func (d *DB) GetResDetPagos(p models.Fechas) ([]models.ResumenDetPago, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetResumenDetPagos, p.Desde, p.Hasta)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Pagos Registrados! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	pagos := []models.ResumenDetPago{}
	pago := models.ResumenDetPago{}
	for rows.Next() {
		rows.Scan(
			&pago.Idinstpago,
			&pago.Descripcion,
			&pago.Cant,
			&pago.Montos,
			&pago.Tasa,
			&pago.Totales,
		)
		pagos = append(pagos, pago)
	}
	rp.Status = 10
	rp.Mensaje = "Pagos listados correctamente!"
	return pagos, rp
}

func (d *DB) GetDivisas() ([]models.Divisas, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetDivisas)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Divisas Registradas! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	divisas := []models.Divisas{}
	divisa := models.Divisas{}
	for rows.Next() {
		rows.Scan(
			&divisa.Id,
			&divisa.Divisa,
			&divisa.Simbolo,
			&divisa.Tasabs,
			&divisa.Fecha,
		)
		divisas = append(divisas, divisa)
	}
	rp.Status = 10
	rp.Mensaje = "Divisas listadas correctamente!"
	return divisas, rp
}

func (d *DB) GetVendedores() ([]models.Vendedor, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetVendedores)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Vendedores Registrados! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	vendedores := []models.Vendedor{}
	vendedor := models.Vendedor{}
	for rows.Next() {
		err3 := rows.Scan(
			&vendedor.Id,
			&vendedor.Cedula,
			&vendedor.Nombre,
			&vendedor.Direccion,
			&vendedor.Telefono,
			&vendedor.Correo,
			&vendedor.Codvend,
		)
		if err3 != nil {
			utils.CreateLog(err3.Error())
		}
		// utils.CreateLog(vendedor.Id,
		// 	vendedor.Cedula,
		// 	vendedor.Nombre,
		// 	vendedor.Direccion,
		// 	vendedor.Telefono,
		// 	vendedor.Correo,
		// 	vendedor.Codvend)
		vendedores = append(vendedores, vendedor)
		// rows.Err()
	}
	rp.Status = 10
	rp.Mensaje = "Vendedores listados correctamente!"
	return vendedores, rp
}

func (d *DB) GetFacturaId(p models.Factura) (models.Factura, models.Respuesta) {
	var rp models.Respuesta

	// utils.CreateLog("funcion Get factura")

	var strSql string
	var strSqlDet string
	if p.Idsesion == -1 {
		strSql = sqlGetInvoicePdf
		strSqlDet = sqlGetItemsFacturas
	} else {
		strSql = sqlGetNotaEntregaPdf
		strSqlDet = sqlGetItemsEntregas
	}

	rows, err := d.db.Query(strSql, p.Id)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Facturas Registradas! " + err.Error()
		utils.CreateLog("No Hay Facturas Registradas! " + err.Error())
		return p, rp
	}
	defer rows.Close()
	// facturas := []models.Factura{}
	factura := models.Factura{}
	for rows.Next() {
		err := rows.Scan(
			&factura.Id,
			&factura.Idcliente,
			&factura.Cliente,
			&factura.Dirfiscal,
			&factura.Rif,
			&factura.Persconta,
			&factura.Tlfconta,
			&factura.Fecha,
			&factura.Subtotal,
			&factura.Dscto,
			&factura.Mototal,
			&factura.Deimp,
			&factura.Tasaimp,
			&factura.Moimp,
			&factura.Moneto,
			&factura.Tasadiv,
			&factura.Monetodiv,
			&factura.Idvendedor,
			&factura.Vendedor,
			&factura.Idsesion,
			&factura.Idmesa,
			&factura.Idmesonero,
			&factura.Pagado,
			&factura.Porpagar,
			&factura.Cambio,
			&factura.Cxcbs,
			&factura.Cxcdiv,
		)
		if err != nil {
			utils.CreateLog("Error al obtener la factura " + err.Error())
		}
		losItems, _ := d.db.Query(strSqlDet, p.Id)
		defer losItems.Close()
		misItems := []models.ItemsFactura{}
		mIt := models.ItemsFactura{}
		for losItems.Next() {
			losItems.Scan(
				&mIt.Idfact,
				&mIt.Codprod,
				&mIt.Producto,
				&mIt.Cant,
				&mIt.Precio,
				&mIt.Subtotal,
				&mIt.Descuento,
				&mIt.Neto,
				&mIt.Descripcion,
				&mIt.Cantmp,
				&mIt.Cantpres,
				&mIt.Iditempres,
			)
			misItems = append(misItems, mIt)
		}
		factura.Items = misItems
		// facturas = append(facturas, factura)
	}
	rp.Status = 10
	rp.Mensaje = "Factura listada correctamente!"
	return factura, rp
}

func (d *DB) GetTopVentas() ([]models.TopVentas, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetTopVentas)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Ventas Registradas! " + err.Error()
		// utils.CreateLog(err.Error())
		return nil, rp
	}
	defer rows.Close()
	ventas := []models.TopVentas{}
	venta := models.TopVentas{}
	for rows.Next() {
		err2 :=
			rows.Scan(
				&venta.Codprod,
				&venta.Producto,
				&venta.Cantidad,
				&venta.Venta,
			)
		if err2 != nil {
			utils.CreateLog(err2.Error())
		}
		// utils.CreateLog(producto.Nombre, producto.Preciovar)
		ventas = append(ventas, venta)
	}
	rp.Status = 10
	rp.Mensaje = "Ventas listadas correctamente!"
	return ventas, rp
}

func (d *DB) GetVentasMes() ([]models.VentasMensual, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetVentasMes)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Ventas Registradas! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	ventas := []models.VentasMensual{}
	venta := models.VentasMensual{}
	for rows.Next() {
		err2 :=
			rows.Scan(
				&venta.Nro,
				&venta.Mes,
				&venta.Bs,
				&venta.Divisas,
			)
		if err2 != nil {
			utils.CreateLog(err2.Error())
		}
		ventas = append(ventas, venta)
	}
	rp.Status = 10
	rp.Mensaje = "Ventas listadas correctamente!"
	return ventas, rp
}

func (d *DB) GetProveedores() ([]models.Proveedor, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetProveedores)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay proveedores Registrados! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	proveedores := []models.Proveedor{}
	proveedor := models.Proveedor{}
	for rows.Next() {
		rows.Scan(
			&proveedor.Id,
			&proveedor.Tipo,
			&proveedor.Proveedor,
			&proveedor.Rif,
			&proveedor.Dirfiscal,
			&proveedor.Ciudad,
			&proveedor.Estado,
			&proveedor.Telf,
			&proveedor.Correo,
			&proveedor.Twitter,
			&proveedor.Facebook,
			&proveedor.Status,
			&proveedor.Obs,
			&proveedor.Clasif,
			&proveedor.Credito,
			&proveedor.Diascred,
			&proveedor.Cxp,
		)
		proveedores = append(proveedores, proveedor)
	}
	rp.Status = 10
	rp.Mensaje = "Proveedores listados correctamente!"
	return proveedores, rp
}

func (d *DB) PostProveedor(p models.Proveedor) models.Respuesta {
	var rp models.Respuesta

	resp, err := d.db.Exec(sqlAddProveedor, p.Tipo,
		p.Proveedor, p.Rif, p.Dirfiscal, p.Ciudad,
		p.Estado, p.Telf, p.Correo, p.Twitter, p.Facebook,
		p.Status, p.Obs, p.Clasif, p.Credito,
		p.Diascred, p.Cxp)

	if err != nil {
		rp.Status = 501
		rp.Mensaje = "No se pudo Agregar la Informacion de Proveedor. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Status = 200
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro agregado Correctamente"
	}
	return rp
}

func (d *DB) UpdProveedor(e models.Proveedor) models.Respuesta {
	var rp models.Respuesta

	resp, err := d.db.Exec(sqlUpdateProveedor,
		e.Id, e.Tipo,
		e.Proveedor, e.Rif, e.Dirfiscal, e.Ciudad,
		e.Estado, e.Telf, e.Correo, e.Twitter, e.Facebook,
		e.Status, e.Obs, e.Clasif, e.Credito,
		e.Diascred, e.Cxp)

	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion del proveedor. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Actualizado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) DelProveedor(e models.Id) models.Respuesta {
	var rp models.Respuesta

	log.Println("eliminar proveedor")
	log.Println(e.Id)
	resp, err := d.db.Exec(sqlDelProveedor, e.Id)

	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion del proveedor. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Eliminado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) GetCxcResumen(e models.Fechas) ([]models.CxcResumen, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetCxcResumen, e.Desde, e.Hasta)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Ventas Registradas! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	ventas := []models.CxcResumen{}
	venta := models.CxcResumen{}
	for rows.Next() {
		err2 :=
			rows.Scan(
				&venta.Id,
				&venta.Idfact,
				&venta.Codclie,
				&venta.Nombre,
				&venta.Saldobs,
				&venta.Saldodiv,
			)
		if err2 != nil {
			utils.CreateLog(err2.Error())
		}
		ventas = append(ventas, venta)
	}
	rp.Status = 10
	rp.Mensaje = "Ventas listadas correctamente!"
	return ventas, rp
}

func (d *DB) GetCxcVencida() ([]models.CxcVencida, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetCxcVencida)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Cxc Registradas! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	cxcs := []models.CxcVencida{}
	cxc := models.CxcVencida{}
	for rows.Next() {
		err2 :=
			rows.Scan(
				&cxc.Rango,
				&cxc.Saldo,
			)
		if err2 != nil {
			utils.CreateLog(err2.Error())
		}
		cxcs = append(cxcs, cxc)
	}
	rp.Status = 10
	rp.Mensaje = "Cuentas por cobrar listadas correctamente!"
	return cxcs, rp
}

func (d *DB) GetCompras() ([]models.Compra, models.Respuesta) {
	var rp models.Respuesta

	rows, err := d.db.Query(sqlGetCompras)
	if err != nil {
		rp.Status = 502
		rp.Mensaje = "No Hay Compras Registradas! " + err.Error()
		return nil, rp
	}
	defer rows.Close()
	compras := []models.Compra{}
	compra := models.Compra{}
	for rows.Next() {
		rows.Scan(
			&compra.Id,
			&compra.Idprov,
			&compra.Proveedor,
			&compra.Fecha,
			&compra.Subtotal,
			&compra.Dscto,
			&compra.Mototal,
			&compra.Deimp,
			&compra.Tasaimp,
			&compra.Moimp,
			&compra.Moneto,
			&compra.Idsesion,
		)
		compra.Items = []models.ItemCompra{}
		compras = append(compras, compra)
	}
	rp.Status = 10
	rp.Mensaje = "Compras listadas correctamente!"
	return compras, rp
}

func (d *DB) AddCompra(c models.Compra) models.Respuesta {
	var rp models.Respuesta

	// Verificar si el usuario está autenticado
	// if idsesion == "*" {
	//     rp.Status = 401
	//     rp.Mensaje = "Credenciales no válidas, debe iniciar sesión!"
	//     return rp
	// }

	// Obtener el ID de la compra
	var nId int
	err := d.db.QueryRow(sqlGetNextIdCompra).Scan(&nId)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener el ID de la compra: " + err.Error()
		return rp
	}

	// Insertar la compra
	_, err = d.db.Exec(sqlInsertCompra, nId,
		c.Idprov,
		c.Fecha,
		c.Subtotal,
		c.Dscto,
		c.Mototal,
		c.Deimp,
		c.Tasaimp,
		c.Moimp,
		c.Moneto,
		c.Idsesion,
	)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al insertar la compra: " + err.Error()
		return rp
	}

	// Insertar los detalles de la compra
	for _, item := range c.Items {
		_, err = d.db.Exec(sqlInsertDetCompra, nId, item.Codprod, item.Deprod, item.Cant, item.Costo, item.Subtotalcosto)
		if err != nil {
			rp.Status = 500
			rp.Mensaje = "Error al insertar el detalle de la compra: " + err.Error()
			return rp
		}
	}

	// Obtener la tasa actual
	var nTasa float64
	err = d.db.QueryRow(sqlGetTasaActual).Scan(&nTasa)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener la tasa actual: " + err.Error()
		return rp
	}

	// Calcular el monto en divisas
	montos := c.Moneto / nTasa

	// Obtener el ID de la cuenta por pagar
	var nIdCxp int
	err = d.db.QueryRow(sqlGetNextIdCxp).Scan(&nIdCxp)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener el ID de la cuenta por pagar: " + err.Error()
		return rp
	}

	// Insertar la cuenta por pagar
	_, err = d.db.Exec(sqlInsertCxp, nIdCxp, nId,
		c.Idprov, c.Fecha, c.Moneto, c.Moneto, montos, montos, c.Idsesion)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al insertar la cuenta por pagar: " + err.Error()
		return rp
	}

	rp.Status = 200
	rp.Mensaje = "Compra creada correctamente"
	return rp
}

func (d *DB) DelCompra(e models.Id) models.Respuesta {
	var rp models.Respuesta

	resp, err := d.db.Exec(sqlDelCompra, e.Id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo Actualizar la Informacion de la Compra. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Eliminado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) AddCita(citas []models.CitaModel) models.Respuesta {
	var rp models.Respuesta
	tx, err := d.db.Begin()
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al iniciar la transacción: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return rp
	}

	totalCitasAgregadas := 0
	// Por cada cita "base" que se recibe, se genera una serie recurrente.
	for _, citaBase := range citas {
		// Generar la serie para 1 año.
		// newUUID := uuid.New().String()
		// groupID := &newUUID

		weeks := citaBase.Weeks
		if weeks <= 0 {
			weeks = 52
		}

		var groupID *string
		if weeks > 1 {
			newUUID := uuid.New().String()
			groupID = &newUUID
		}

		// Bucle por la cantidad de semanas indicadas
		for i := 0; i < weeks; i++ {
			nuevaCita := citaBase
			nuevaCita.Inicio = citaBase.Inicio.AddDate(0, 0, 7*i)
			nuevaCita.Fin = citaBase.Fin.AddDate(0, 0, 7*i)
			nuevaCita.GroupID = groupID

			_, err := tx.Exec(sqlPostCita, nuevaCita.IdDoctor, nuevaCita.Cedula, nuevaCita.Motivo, nuevaCita.Inicio, nuevaCita.Fin, nuevaCita.Status, nuevaCita.Color,
				nuevaCita.Montoref, nuevaCita.Tasa, nuevaCita.Montobs, nuevaCita.Pagado, nuevaCita.GroupID)
			if err != nil {
				tx.Rollback()
				rp.Status = 500
				rp.Mensaje = "Error al agregar cita: " + err.Error()
				utils.CreateLog("Error al agregar cita: " + err.Error())
				return rp
			}
			totalCitasAgregadas++
		}
	}

	if err := tx.Commit(); err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al confirmar la transacción: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return rp
	}

	rp.Status = 200
	rp.Mensaje = fmt.Sprintf("%d cita(s) agregada(s) correctamente", totalCitasAgregadas)
	return rp
}

func (d *DB) UpdateCita(cita models.CitaModel) models.Respuesta {
	var rp models.Respuesta

	// Si se solicita actualizar la serie y existe un GroupID válido
	if cita.UpdateSeries && cita.GroupID != nil && *cita.GroupID != "" {
		// 1. Obtener la cita original para calcular el desplazamiento de tiempo (si hubo cambio de hora)
		var oldInicio time.Time
		err := d.db.QueryRow("SELECT inicio FROM medi001.citas WHERE id = $1", cita.Id).Scan(&oldInicio)
		if err != nil {
			rp.Status = 500
			rp.Mensaje = "Error al obtener la cita original para calcular desplazamiento: " + err.Error()
			utils.CreateLog(rp.Mensaje)
			return rp
		}

		// 2. Calcular la diferencia (delta) para desplazar toda la serie
		diff := cita.Inicio.Sub(oldInicio)
		intervalStr := fmt.Sprintf("%f seconds", diff.Seconds())

		// 3. Actualizar toda la serie: aplica el desplazamiento a las fechas y setea los nuevos valores
		const sqlUpdCitaSeries = `
			UPDATE medi001.citas 
			SET id_doctor=$1, cedula=$2, motivo=$3, status=$4, color=$5, montoref=$6, tasa=$7, montobs=$8, pagado=$9, 
				inicio = inicio + $10::interval, 
				fin = fin + $10::interval 
			WHERE group_id=$11`

		res, err := d.db.Exec(sqlUpdCitaSeries,
			cita.IdDoctor, cita.Cedula, cita.Motivo, cita.Status,
			cita.Color, cita.Montoref, cita.Tasa, cita.Montobs, cita.Pagado,
			intervalStr, cita.GroupID)

		if err != nil {
			rp.Status = 500
			rp.Mensaje = "Error al actualizar la serie de citas: " + err.Error()
			utils.CreateLog(rp.Mensaje)
			return rp
		}
		rows, _ := res.RowsAffected()
		rp.Status = 200
		rp.Mensaje = fmt.Sprintf("Serie actualizada correctamente (%d citas modificadas)", rows)
		return rp
	}

	// Comportamiento original: Actualizar solo la cita individual
	_, err := d.db.Exec(sqlUpdCita, cita.Id, cita.IdDoctor, cita.Cedula, cita.Motivo, cita.Inicio, cita.Fin, cita.Status,
		cita.Color, cita.Montoref, cita.Tasa, cita.Montobs, cita.Pagado, cita.GroupID)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al actualizar cita: " + err.Error()
		utils.CreateLog("Error al actualizar cita: " + err.Error())
		return rp
	}
	rp.Status = 200
	rp.Mensaje = "Cita actualizada correctamente"
	return rp
}

func (d *DB) UpdateDiagnosticoCita(cita models.CitaModel) models.Respuesta {
	var rp models.Respuesta

	// Asegúrate de tener definida la constante sqlUpdDiagnostico con tu consulta SQL
	_, err := d.db.Exec(sqlUpdDiagnostico, cita.Diagnostico, cita.Id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al actualizar diagnóstico: " + err.Error()
		utils.CreateLog("Error al actualizar diagnóstico: " + err.Error())
		return rp
	}

	rp.Status = 200
	rp.Mensaje = "Diagnóstico actualizado correctamente"
	return rp
}

func (d *DB) GetCitas() ([]models.CitaModel, models.Respuesta) {
	var rp models.Respuesta
	var citas []models.CitaModel

	rows, err := d.db.Query(sqlGetCitas)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener citas: " + err.Error()
		utils.CreateLog("Error al obteniendo las citas: " + rp.Mensaje)
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
			&cita.GroupID,
		)
		if err != nil {
			rp.Status = 500
			rp.Mensaje = "Error al escanear cita: " + err.Error()
			return nil, rp
		}
		citas = append(citas, cita)
	}

	rp.Status = 200
	rp.Mensaje = "Citas obtenidas correctamente"
	return citas, rp
}

func (d *DB) GetInformesMedicos(idPaciente int) ([]models.InformeMedico, models.Respuesta) {
	var rp models.Respuesta
	const sqlGetInformesPorPaciente = `SELECT id, id_paciente, fecha, id_doctor, id_cita, diagnostico, evolucion, plan, recomendaciones FROM medi001.informes_medicos WHERE id_paciente = $1 ORDER BY fecha DESC`

	rows, err := d.db.Query(sqlGetInformesPorPaciente, idPaciente)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener informes médicos: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return nil, rp
	}
	defer rows.Close()

	var informes []models.InformeMedico
	for rows.Next() {
		var informe models.InformeMedico
		var idCita sql.NullInt64
		err := rows.Scan(
			&informe.Id,
			&informe.IdPaciente,
			&informe.Fecha,
			&informe.IdDoctor,
			&idCita,
			&informe.Diagnostico,
			&informe.Evolucion,
			&informe.Plan,
			&informe.Recomendaciones,
		)
		if err != nil {
			rp.Status = 500
			rp.Mensaje = "Error al escanear informe médico: " + err.Error()
			utils.CreateLog(rp.Mensaje)
			return nil, rp
		}
		if idCita.Valid {
			informe.IdCita = int(idCita.Int64)
		}
		informes = append(informes, informe)
	}
	rp.Status = 200
	rp.Mensaje = "Informes médicos obtenidos correctamente"
	return informes, rp
}

func (d *DB) GetInformeMedico(id int) (models.InformeMedico, models.Respuesta) {
	var rp models.Respuesta
	var informe models.InformeMedico
	const sqlGetInforme = `SELECT id, id_paciente, fecha, id_doctor, id_cita, diagnostico, evolucion, plan, recomendaciones FROM medi001.informes_medicos WHERE id = $1`

	var idCita sql.NullInt64
	err := d.db.QueryRow(sqlGetInforme, id).Scan(
		&informe.Id,
		&informe.IdPaciente,
		&informe.Fecha,
		&informe.IdDoctor,
		&idCita,
		&informe.Diagnostico,
		&informe.Evolucion,
		&informe.Plan,
		&informe.Recomendaciones,
	)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener informe médico: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return informe, rp
	}
	if idCita.Valid {
		informe.IdCita = int(idCita.Int64)
	}
	rp.Status = 200
	rp.Mensaje = "Informe médico obtenido correctamente"
	return informe, rp
}

func (d *DB) AddInformeMedico(i models.InformeMedico) models.Respuesta {
	var rp models.Respuesta
	const sqlAddInforme = `INSERT INTO medi001.informes_medicos (id_paciente, fecha, id_doctor, id_cita, diagnostico, evolucion, plan, recomendaciones) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	var idCita interface{}
	if i.IdCita != 0 {
		idCita = i.IdCita
	} else {
		idCita = nil
	}

	var newID int
	err := d.db.QueryRow(sqlAddInforme, i.IdPaciente, i.Fecha, i.IdDoctor, idCita, i.Diagnostico, i.Evolucion, i.Plan, i.Recomendaciones).Scan(&newID)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al agregar informe médico: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return rp
	}
	rp.Status = 200
	rp.Mensaje = "Informe médico agregado correctamente con ID: " + strconv.Itoa(newID)
	return rp
}

func (d *DB) UpdInformeMedico(i models.InformeMedico) models.Respuesta {
	var rp models.Respuesta
	const sqlUpdInforme = `UPDATE medi001.informes_medicos SET fecha=$1, id_doctor=$2, id_cita=$3, diagnostico=$4, evolucion=$5, plan=$6, recomendaciones=$7 WHERE id=$8`

	var idCita interface{}
	if i.IdCita != 0 {
		idCita = i.IdCita
	} else {
		idCita = nil
	}

	res, err := d.db.Exec(sqlUpdInforme, i.Fecha, i.IdDoctor, idCita, i.Diagnostico, i.Evolucion, i.Plan, i.Recomendaciones, i.Id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al actualizar informe médico: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return rp
	}

	rows, _ := res.RowsAffected()
	if rows > 0 {
		rp.Status = 200
		rp.Mensaje = "Informe médico actualizado correctamente"
	} else {
		rp.Status = 404
		rp.Mensaje = "No se encontró el informe médico para actualizar"
	}
	return rp
}

func (d *DB) DelInformeMedico(id int) models.Respuesta {
	var rp models.Respuesta
	const sqlDelInforme = `DELETE FROM medi001.informes_medicos WHERE id=$1`

	res, err := d.db.Exec(sqlDelInforme, id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al eliminar informe médico: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return rp
	}
	rows, _ := res.RowsAffected()
	if rows > 0 {
		rp.Status = 200
		rp.Mensaje = "Informe médico eliminado correctamente"
	} else {
		rp.Status = 404
		rp.Mensaje = "No se encontró el informe médico para eliminar"
	}
	return rp
}

func (d *DB) GetCitasPaciente(p models.PacientesModel) ([]models.CitaModel, models.Respuesta) {
	var rp models.Respuesta
	var citas []models.CitaModel

	rows, err := d.db.Query(sqlGetCitaPaciente, p.Cedula)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener citas: " + err.Error()
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
			&cita.GroupID,
		)
		if err != nil {
			rp.Status = 500
			rp.Mensaje = "Error al escanear cita: " + err.Error()
			return nil, rp
		}
		citas = append(citas, cita)
	}

	rp.Status = 200
	rp.Mensaje = "Citas obtenidas correctamente"
	return citas, rp
}

func (d *DB) DelPresupuesto(e models.Id) models.Respuesta {

	var rp models.Respuesta

	var strSql = ""
	switch e.Id {
	case "-1":
		strSql = sqlDelPresupuesto
	case "-2":
		strSql = sqlDelPresupuesto + "WHERE vence < Now();"
	default:
		strSql = sqlDelPresupuesto + "WHERE id = " + e.Id
	}
	resp, err := d.db.Exec(strSql)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al eliminar el presupuesto: " + err.Error()
		utils.CreateLog("Error al eliminar el presupuesto: " + err.Error())
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Eliminado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) DelCita(e models.IdCitas) models.Respuesta {
	var rp models.Respuesta
	var sqlStr = ""
	if e.Tipo == 1 {
		sqlStr = sqlDelCita
	} else {
		sqlStr = sqlDelCitaAll
	}
	resp, err := d.db.Exec(sqlStr, e.Id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo eliminar la Cita. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Eliminado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

func (d *DB) DelPayment(e models.Id) models.Respuesta {
	var rp models.Respuesta

	resp, err := d.db.Exec(sqlDelPayments, e.Id)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "No se pudo eliminar el pago. " + err.Error()
		return rp
	}
	datos, err1 := resp.RowsAffected()
	if err1 != nil {
		rp.Status = 500
		rp.Mensaje = err1.Error()
	} else if datos > 0 {
		rp.Mensaje = strconv.FormatInt(datos, 10) + " Registro Eliminado Correctamente"
		rp.Status = 200
	} else {
		rp.Status = 201
		rp.Mensaje = "No se encontro ningun registro con los datos proporcionados!"
	}
	return rp
}

// FetchExchangeRate fetches the current VES exchange rate from an external API.
func (d *DB) FetchExchangeRate() (float64, models.Respuesta) {
	var rp models.Respuesta
	rp.Status = 500 // Default to error

	// The API URL provided by the user
	urlTasa := "https://api.exchangerate-api.com/v4/latest/USD"

	resp, err := http.Get(urlTasa)
	if err != nil {
		rp.Mensaje = "Error al realizar la solicitud HTTP a la API de tasas: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return 0, rp
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		rp.Mensaje = fmt.Sprintf("Error al obtener la tasa del API: %s", resp.Status)
		utils.CreateLog(rp.Mensaje)
		return 0, rp
	}

	var apiResponse struct {
		Rates map[string]float64 `json:"rates"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		rp.Mensaje = "Error al decodificar la respuesta de la API de tasas: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return 0, rp
	}

	vesRate, ok := apiResponse.Rates["VES"]
	if !ok {
		rp.Mensaje = "Tasa VES no encontrada en la respuesta del API."
		utils.CreateLog(rp.Mensaje)
		return 0, rp
	}

	// Round the rate to 2 decimal places as requested
	roundedRate := roundFloat(vesRate, 2)

	rp.Status = 200
	rp.Mensaje = "Tasa de cambio obtenida exitosamente."
	return roundedRate, rp
}

// UpdateUnpaidAppointmentsVESRate updates the montobs for unpaid, non-cancelled, non-exonerated appointments.
func (d *DB) UpdateUnpaidAppointmentsVESRate(newRate float64) models.Respuesta {
	var rp models.Respuesta
	res, err := d.db.Exec(sqlUpdateUnpaidAppointmentsVESRate, newRate)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al actualizar montos de citas: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return rp
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener filas afectadas después de la actualización: " + err.Error()
		utils.CreateLog(rp.Mensaje)
		return rp
	}

	rp.Status = 200
	rp.Mensaje = fmt.Sprintf("Montos de %d citas actualizados correctamente con la nueva tasa: %.2f", rowsAffected, newRate)
	utils.CreateLog(rp.Mensaje)
	return rp
}

func (d *DB) GetCitasFecha(p models.Fechas) ([]models.CitaModel, models.Respuesta) {
	var rp models.Respuesta
	var citas []models.CitaModel

	rows, err := d.db.Query(sqlGetCitasFecha, p.Desde, p.Hasta)
	if err != nil {
		rp.Status = 500
		rp.Mensaje = "Error al obtener citas por fecha: " + err.Error()
		utils.CreateLog(rp.Mensaje)
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
			&cita.GroupID,
		)
		if err != nil {
			rp.Status = 500
			rp.Mensaje = "Error al escanear cita: " + err.Error()
			return nil, rp
		}
		citas = append(citas, cita)
	}

	rp.Status = 200
	rp.Mensaje = "Citas obtenidas correctamente"
	return citas, rp
}
