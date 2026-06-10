package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"omhmre.com/centromedico/app/domain/database"
)

type App struct {
	Router *mux.Router
	DB     database.PostDB
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}
	enableCORS(a.Router)
	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/health", a.HealthHandler()).Methods("GET")
	a.Router.HandleFunc("/menuweb", a.MenuWeb()).Methods("GET")
	// Inventario
	a.Router.HandleFunc("/getinventario", a.GetInventario()).Methods("GET")
	a.Router.HandleFunc("/getinventarioformal", a.GetInventarioFormal()).Methods("GET")
	a.Router.HandleFunc("/getinvcompacto", a.GetInventarioCompacto()).Methods("GET")
	a.Router.HandleFunc("/getinventariomenu", a.GetInventarioMenu()).Methods("GET")
	a.Router.HandleFunc("/postinventario", a.AddInventario()).Methods("POST")
	a.Router.HandleFunc("/postitemsinventario", a.AddItemsInventario()).Methods("POST")
	a.Router.HandleFunc("/postpresentaciones", a.AddPresenInventario()).Methods("POST")
	a.Router.HandleFunc("/putinventario", a.UpdInventario()).Methods("PUT")
	a.Router.HandleFunc("/delinventario", a.DelInventario()).Methods("DELETE")
	// Menu
	a.Router.HandleFunc("/getmenu", a.GetMenu()).Methods("GET")
	a.Router.HandleFunc("/getmenuclases", a.GetMenuClases()).Methods("GET")
	a.Router.HandleFunc("/postmenu", a.AddMenu()).Methods("POST")
	a.Router.HandleFunc("/postallmenu", a.AddAllMenu()).Methods("POST")
	a.Router.HandleFunc("/putmenu", a.UpdMenu()).Methods("PUT")
	a.Router.HandleFunc("/delmenu", a.DelMenu()).Methods("DELETE")
	a.Router.HandleFunc("/delallmenu", a.DelAllMenu()).Methods("DELETE")
	a.Router.HandleFunc("/getmenucompleto", a.GetMenuCompleto()).Methods("GET")
	// Clases
	a.Router.HandleFunc("/getclases", a.GetClases()).Methods("GET")
	a.Router.HandleFunc("/postclases", a.AddClase()).Methods("POST")
	a.Router.HandleFunc("/putclases", a.UpdClase()).Methods("PUT")
	a.Router.HandleFunc("/delclases", a.DelClase()).Methods("DELETE")
	// Empresas
	a.Router.HandleFunc("/getempre", a.GetEmpre()).Methods("GET")
	a.Router.HandleFunc("/postempre", a.AddEmpresa()).Methods("POST")
	a.Router.HandleFunc("/updempre", a.UpdEmpresa()).Methods("PUT")
	a.Router.HandleFunc("/delempre", a.DelEmpresa()).Methods("DELETE")
	// Images
	a.Router.HandleFunc("/getimages", a.Images()).Methods("GET")
	// PreFacturas
	a.Router.HandleFunc("/getprefacturas", a.GetPrefacturas()).Methods("GET")
	a.Router.HandleFunc("/getprefactura", a.GetPrefactura()).Methods("POST")
	a.Router.HandleFunc("/postprefactura", a.PostPreFactura()).Methods("POST")
	a.Router.HandleFunc("/putprefactura", a.UpdPreFactura()).Methods("POST")
	a.Router.HandleFunc("/getprefactura", a.GetPreFactura()).Methods("GET")
	// Clientes
	a.Router.HandleFunc("/getclientes", a.GetClientes()).Methods("GET")
	a.Router.HandleFunc("/postcliente", a.AddCliente()).Methods("POST")
	a.Router.HandleFunc("/putclientes", a.UpdCliente()).Methods("PUT")
	a.Router.HandleFunc("/delclientes", a.DelCliente()).Methods("DELETE")
	a.Router.HandleFunc("/getcxcs", a.GetCxcVencida()).Methods("GET")
	// Mesas
	a.Router.HandleFunc("/getmesas", a.GetMesas()).Methods("GET")
	a.Router.HandleFunc("/putmesa", a.UpdMesa()).Methods("PUT")
	a.Router.HandleFunc("/postmesa", a.AddMesa()).Methods("POST")
	a.Router.HandleFunc("/delmesa", a.DelMesa()).Methods("DELETE")
	a.Router.HandleFunc("/abrirmesa", a.AbrirMesa()).Methods("PUT")
	a.Router.HandleFunc("/limpiarmesa", a.cleanMesa()).Methods("PUT")
	// Mesoneros
	a.Router.HandleFunc("/getmesoneros", a.GetMesoneros()).Methods("GET")
	a.Router.HandleFunc("/postmesonero", a.AddMesonero()).Methods("POST")
	a.Router.HandleFunc("/putmesonero", a.UpdMesonero()).Methods("PUT")
	a.Router.HandleFunc("/delmesonero", a.DelMesonero()).Methods("DELETE")
	// Instrumentos de pagos
	a.Router.HandleFunc("/getinstrumentos", a.GetInstrumentos()).Methods("GET")
	// Facturas
	a.Router.HandleFunc("/postfactura", a.PostFactura()).Methods("POST")
	a.Router.HandleFunc("/getfacturas", a.GetFacturas()).Methods("GET")
	a.Router.HandleFunc("/putanularfac", a.UpdAnularFactura()).Methods("PUT")
	a.Router.HandleFunc("/getfacturaid", a.GetFacturaId()).Methods("POST")
	// Notas de Entrega
	a.Router.HandleFunc("/postentrega", a.PostNotaEntrega()).Methods("POST")
	a.Router.HandleFunc("/getnotasent", a.GetNotasEntrega()).Methods("GET")
	// Ventas
	a.Router.HandleFunc("/postventasdia", a.PostVentasFactura()).Methods("POST")
	a.Router.HandleFunc("/postvtasprod", a.PostVentasProductos()).Methods("POST")
	a.Router.HandleFunc("/gettopventas", a.GetTopVentas()).Methods("GET")
	a.Router.HandleFunc("/getventasmes", a.GetVentasMes()).Methods("GET")
	a.Router.HandleFunc("/sendventas", a.SendVentasMail()).Methods("POST")
	a.Router.HandleFunc("/getventasfecha", a.GetVentas()).Methods("POST")
	// Tasas Y Divisas
	a.Router.HandleFunc("/chgtasa", a.UpdDivisa()).Methods("PUT")
	a.Router.HandleFunc("/divisas", a.GetDivisas()).Methods("GET")
	// Pagos
	a.Router.HandleFunc("/postpagos", a.PostPagos()).Methods("POST")
	a.Router.HandleFunc("/getdetpagosfecha", a.GetDetPagosFecha()).Methods("POST")
	a.Router.HandleFunc("/getresdetpagos", a.GetResumenDetPagos()).Methods("POST")
	// Vendedores
	a.Router.HandleFunc("/getvendedores", a.GetVendedores()).Methods("GET")
	// Usuarios
	a.Router.HandleFunc("/getusuarios", a.GetUsuarios()).Methods("GET")
	a.Router.HandleFunc("/postusuario", a.PostUsuario()).Methods("POST")
	a.Router.HandleFunc("/putusuario", a.PutUsuario()).Methods("PUT")
	a.Router.HandleFunc("/delusuario", a.DelUsuario()).Methods("DELETE")
	a.Router.HandleFunc("/login", a.Login()).Methods("POST")
	a.Router.HandleFunc("/putpassword", a.ChangePassword()).Methods("PUT")
	// Configuración
	a.Router.HandleFunc("/clearlogs", a.ClearLogs()).Methods("POST")
	a.Router.HandleFunc("/getlogs", a.GetLogs()).Methods("POST")
	a.Router.HandleFunc("/getparametros", a.GetParametros()).Methods("GET")
	a.Router.HandleFunc("/postparametro", a.PostParametro()).Methods("POST")
	a.Router.HandleFunc("/putparametro", a.PutParametro()).Methods("PUT")
	// Proveedores
	a.Router.HandleFunc("/getproveedores", a.GetProveedores()).Methods("GET")
	a.Router.HandleFunc("/postproveedor", a.PostProveedor()).Methods("POST")
	a.Router.HandleFunc("/putproveedor", a.PutProveedor()).Methods("PUT")
	a.Router.HandleFunc("/delproveedor", a.DelProveedor()).Methods("DELETE")
	a.Router.HandleFunc("/getcxcresumen", a.GetCxcResumen()).Methods("POST")
	// Compras
	a.Router.HandleFunc("/getcompras", a.GetCompras()).Methods("GET")
	a.Router.HandleFunc("/postcompras", a.PostCompra()).Methods("POST")
	a.Router.HandleFunc("/delcompra", a.DelCompra()).Methods("DELETE")
	// Email Config
	a.Router.HandleFunc("/postemailconfig", a.GetEmailConfig()).Methods("POST")
	a.Router.HandleFunc("/putemailconfig", a.PutEmailConfig()).Methods("PUT")
	// presupuestos
	a.Router.HandleFunc("/getpresupuestos", a.GetPresupuestos()).Methods("GET")
	a.Router.HandleFunc("/postpresupuesto", a.PostPresupuesto()).Methods("POST")
	a.Router.HandleFunc("/getpresupuestoid", a.GetPresupuestoId()).Methods("POST")
	a.Router.HandleFunc("/delpresupuesto", a.DelPresupuesto()).Methods("DELETE")
	// Citas Medicas
	a.Router.HandleFunc("/getcitas", a.GetCitas()).Methods("GET")
	a.Router.HandleFunc("/getcitaspaciente", a.GetCitasPaciente()).Methods("POST")
	a.Router.HandleFunc("/putcitas", a.UpdateCita()).Methods("PUT")
	a.Router.HandleFunc("/postcita", a.AddCita()).Methods("POST")
	a.Router.HandleFunc("/delcita", a.DelCita()).Methods("DELETE")
	a.Router.HandleFunc("/getcitasfecha", a.GetCitasFecha()).Methods("POST")
	a.Router.HandleFunc("/putdiagnostico", a.AddDiagnosis()).Methods("PUT")
	// Doctores
	a.Router.HandleFunc("/getdoctores", a.GetDoctores()).Methods("GET")
	a.Router.HandleFunc("/putdoctores", a.UpdateDoctores()).Methods("PUT")
	a.Router.HandleFunc("/postdoctor", a.PostDoctor()).Methods("POST")
	a.Router.HandleFunc("/deldoctor", a.DelDoctor()).Methods("DELETE")
	// Pacientes
	a.Router.HandleFunc("/getpacientes", a.GetPacientes()).Methods("GET")
	a.Router.HandleFunc("/postpaciente", a.PostPaciente()).Methods("POST")
	a.Router.HandleFunc("/putpaciente", a.UpdPaciente()).Methods("PUT")
	a.Router.HandleFunc("/delpaciente", a.DelPaciente()).Methods("DELETE")
	// Precios por especialidad del paciente
	a.Router.HandleFunc("/pacientes/precios", a.UpsertPrecioEspecialidad()).Methods("POST")
	a.Router.HandleFunc("/pacientes/precios", a.DelPrecioEspecialidad()).Methods("DELETE")
	// Relacion de Pagos
	a.Router.HandleFunc("/getpayments", a.GetPayments()).Methods("POST")
	a.Router.HandleFunc("/postpayments", a.PostPayments()).Methods("POST")
	a.Router.HandleFunc("/getrelpagos", a.GetRelPagos()).Methods("POST")
	a.Router.HandleFunc("/delpayment", a.DelPayment()).Methods("DELETE")
	
	// Historial Clínico
	a.Router.HandleFunc("/clinical_history", a.GetClinicalHistoryHandler()).Methods("GET")
	a.Router.HandleFunc("/clinical_history", a.UpsertClinicalHistoryHandler()).Methods("POST")
	a.Router.HandleFunc("/clinical_records", a.GetClinicalRecordsHandler()).Methods("GET")
	a.Router.HandleFunc("/clinical_records", a.PostClinicalRecordHandler()).Methods("POST")
}
