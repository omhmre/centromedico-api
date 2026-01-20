package app

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/websocket"
	"omhmre.com/centromedico/app/domain/database"
	"omhmre.com/centromedico/app/domain/utils"
	ws "omhmre.com/centromedico/app/websocket"
)

// App holds the application's dependencies
type App struct {
	Router *http.ServeMux
	DB     database.PostDB
	Hub    *ws.Hub
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections for development.
		// For production, you should implement a proper origin check.
		// e.g., return r.Header.Get("Origin") == "http://your-frontend-domain.com"
		return true
	},
}

// New creates a new App instance and initializes routes
func New(hub *ws.Hub) *App {
	a := &App{
		Router: http.NewServeMux(),
		Hub:    hub,
	}
	a.initializeRoutes()
	return a
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		// Exponer headers personalizados para que el cliente los pueda leer
		w.Header().Set("Access-Control-Expose-Headers", "X-Total-Count")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (a *App) WrapWithCORS(handler http.Handler) http.Handler {
	return enableCORS(handler)
}

// ServeWs handles websocket requests from the peer.
func (a *App) ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.CreateLog("Error upgrading to websocket: " + err.Error())
		return
	}
	client := ws.NewClient(a.Hub, conn)
	a.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}

// initializeRoutes registers all the application's routes.
func (a *App) initializeRoutes() {
	// Static files
	a.Router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(".", "app", "templates")))))

	// WebSocket
	a.Router.HandleFunc("/ws", a.ServeWs)

	// General
	a.Router.HandleFunc("/", a.IndexHandler())
	a.Router.HandleFunc("/health", a.HealthHandler())
	a.Router.HandleFunc("/menuweb", a.MenuWeb())
	a.Router.HandleFunc("/configvar", a.ConfigVar())

	// Inventario
	a.Router.HandleFunc("/getinventario", a.GetInventario())
	a.Router.HandleFunc("/getinventarioformal", a.GetInventarioFormal())
	a.Router.HandleFunc("/getinvcompacto", a.GetInventarioCompacto())
	a.Router.HandleFunc("/getinventariomenu", a.GetInventarioMenu())
	a.Router.HandleFunc("/postinventario", a.AddInventario())
	a.Router.HandleFunc("/postitemsinventario", a.AddItemsInventario())
	a.Router.HandleFunc("/postpresentaciones", a.AddPresenInventario())
	a.Router.HandleFunc("/putinventario", a.UpdInventario())
	a.Router.HandleFunc("/delinventario", a.DelInventario())

	// Menu
	a.Router.HandleFunc("/getmenu", a.GetMenu())
	a.Router.HandleFunc("/getmenuclases", a.GetMenuClases())
	a.Router.HandleFunc("/postmenu", a.AddMenu())
	a.Router.HandleFunc("/postallmenu", a.AddAllMenu())
	a.Router.HandleFunc("/putmenu", a.UpdMenu())
	a.Router.HandleFunc("/delmenu", a.DelMenu())
	a.Router.HandleFunc("/delallmenu", a.DelAllMenu())
	a.Router.HandleFunc("/getmenucompleto", a.GetMenuCompleto())

	// Clases
	a.Router.HandleFunc("/getclases", a.GetClases())
	a.Router.HandleFunc("/postclases", a.AddClase())
	a.Router.HandleFunc("/putclases", a.UpdClase())
	a.Router.HandleFunc("/delclases", a.DelClase())

	// Empresas
	a.Router.HandleFunc("/getempre", a.GetEmpre())
	a.Router.HandleFunc("/postempre", a.AddEmpresa())
	a.Router.HandleFunc("/updempre", a.UpdEmpresa())
	a.Router.HandleFunc("/delempre", a.DelEmpresa())

	// Images
	a.Router.HandleFunc("/getimages", a.Images())

	// PreFacturas
	a.Router.HandleFunc("/getprefacturas", a.GetPrefacturas())
	a.Router.HandleFunc("/getprefactura", a.GetPrefactura())
	a.Router.HandleFunc("/postprefactura", a.PostPreFactura())
	a.Router.HandleFunc("/putprefactura", a.UpdPreFactura())

	// Clientes
	a.Router.HandleFunc("/getclientes", a.GetClientes())
	a.Router.HandleFunc("/postcliente", a.AddCliente())
	a.Router.HandleFunc("/putclientes", a.UpdCliente())
	a.Router.HandleFunc("/delclientes", a.DelCliente())
	a.Router.HandleFunc("/getcxcs", a.GetCxcVencida())

	// Mesas
	a.Router.HandleFunc("/getmesas", a.GetMesas())
	a.Router.HandleFunc("/putmesa", a.UpdMesa())
	a.Router.HandleFunc("/postmesa", a.AddMesa())
	a.Router.HandleFunc("/delmesa", a.DelMesa())
	a.Router.HandleFunc("/abrirmesa", a.AbrirMesa())
	a.Router.HandleFunc("/limpiarmesa", a.cleanMesa())

	// Mesoneros
	a.Router.HandleFunc("/getmesoneros", a.GetMesoneros())
	a.Router.HandleFunc("/postmesonero", a.AddMesonero())
	a.Router.HandleFunc("/putmesonero", a.UpdMesonero())
	a.Router.HandleFunc("/delmesonero", a.DelMesonero())

	// Instrumentos de pagos
	a.Router.HandleFunc("/getinstrumentos", a.GetInstrumentos())

	// Facturas
	a.Router.HandleFunc("/postfactura", a.PostFactura())
	a.Router.HandleFunc("/getfacturas", a.GetFacturas())
	a.Router.HandleFunc("/putanularfac", a.UpdAnularFactura())
	a.Router.HandleFunc("/getfacturaid", a.GetFacturaId())

	// Notas de Entrega
	a.Router.HandleFunc("/postentrega", a.PostNotaEntrega())
	a.Router.HandleFunc("/getnotasent", a.GetNotasEntrega())

	// Ventas
	a.Router.HandleFunc("/postventasdia", a.PostVentasFactura())
	a.Router.HandleFunc("/postvtasprod", a.PostVentasProductos())
	a.Router.HandleFunc("/gettopventas", a.GetTopVentas())
	a.Router.HandleFunc("/getventasmes", a.GetVentasMes())
	a.Router.HandleFunc("/sendventas", a.SendVentasMail())
	a.Router.HandleFunc("/getventasfecha", a.GetVentas())

	// Tasas Y Divisas
	a.Router.HandleFunc("/chgtasa", a.UpdDivisa())
	a.Router.HandleFunc("/divisas", a.GetDivisas())

	// Pagos
	a.Router.HandleFunc("/postpagos", a.PostPagos())
	a.Router.HandleFunc("/getdetpagosfecha", a.GetDetPagosFecha())
	a.Router.HandleFunc("/getresdetpagos", a.GetResumenDetPagos())

	// Vendedores
	a.Router.HandleFunc("/getvendedores", a.GetVendedores())

	// Usuarios
	a.Router.HandleFunc("/getusuarios", a.GetUsuarios())
	a.Router.HandleFunc("/postusuario", a.PostUsuario())
	a.Router.HandleFunc("/putusuario", a.PutUsuario())
	a.Router.HandleFunc("/delusuario", a.DelUsuario())
	a.Router.HandleFunc("/login", a.Login())
	a.Router.HandleFunc("/putpassword", a.ChangePassword())

	// Configuración
	a.Router.HandleFunc("/clearlogs", a.ClearLogs())
	a.Router.HandleFunc("/getlogs", a.GetLogs())
	a.Router.HandleFunc("/getparametros", a.GetParametros())
	a.Router.HandleFunc("/postparametro", a.PostParametro())
	a.Router.HandleFunc("/putparametro", a.PutParametro())

	// Proveedores
	a.Router.HandleFunc("/getproveedores", a.GetProveedores())
	a.Router.HandleFunc("/postproveedor", a.PostProveedor())
	a.Router.HandleFunc("/putproveedor", a.PutProveedor())
	a.Router.HandleFunc("/delproveedor", a.DelProveedor())
	a.Router.HandleFunc("/getcxcresumen", a.GetCxcResumen())

	// Compras
	a.Router.HandleFunc("/getcompras", a.GetCompras())
	a.Router.HandleFunc("/postcompras", a.PostCompra())
	a.Router.HandleFunc("/delcompra", a.DelCompra())

	// Email Config
	a.Router.HandleFunc("/postemailconfig", a.GetEmailConfig())
	a.Router.HandleFunc("/putemailconfig", a.PutEmailConfig())

	// presupuestos
	a.Router.HandleFunc("/getpresupuestos", a.GetPresupuestos())
	a.Router.HandleFunc("/postpresupuesto", a.PostPresupuesto())
	a.Router.HandleFunc("/getpresupuestoid", a.GetPresupuestoId())
	a.Router.HandleFunc("/delpresupuesto", a.DelPresupuesto())

	// Citas Medicas
	a.Router.HandleFunc("/getcitas", a.GetCitas())
	a.Router.HandleFunc("/getcitaspaciente", a.GetCitasPaciente())
	a.Router.HandleFunc("/putcitas", a.UpdateCita())
	a.Router.HandleFunc("/postcita", a.AddCita())
	a.Router.HandleFunc("/delcita", a.DelCita())
	a.Router.HandleFunc("/getcitasfecha", a.GetCitasFecha())
	a.Router.HandleFunc("/putdiagnostico", a.AddDiagnosis())
	a.Router.HandleFunc("/update-exchange-rate", a.UpdateExchangeRateAndAppointments())

	// Doctores
	a.Router.HandleFunc("/getdoctores", a.GetDoctores())
	a.Router.HandleFunc("/putdoctores", a.UpdateDoctores())
	a.Router.HandleFunc("/postdoctor", a.PostDoctor())
	a.Router.HandleFunc("/deldoctor", a.DelDoctor())

	// Pacientes
	a.Router.HandleFunc("/getpacientes", a.GetPacientes())
	a.Router.HandleFunc("/postpaciente", a.PostPaciente())
	a.Router.HandleFunc("/putpaciente", a.UpdPaciente())
	a.Router.HandleFunc("/delpaciente", a.DelPaciente())

	// Precios por especialidad del paciente
	a.Router.HandleFunc("/pacientes/precios", a.HandlePreciosEspecialidad())

	// Relacion de Pagos
	a.Router.HandleFunc("/getpayments", a.GetPayments())
	a.Router.HandleFunc("/postpayments", a.PostPayments())
	a.Router.HandleFunc("/getrelpagos", a.GetRelPagos())
	a.Router.HandleFunc("/delpayment", a.DelPayment())
}
