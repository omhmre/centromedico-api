package database

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	"omhmre.com/centromedico/app/domain/utils"
)

// Variables públicas
var (
	DB_SERVER    string
	DB_USER      string
	DB_PASSWORD  string
	DB_NAME      string
	DB_PORT      string
	DB_POOL_MODE string
	DB_SSL_MODE  string
	dbinfo       string
	PUERTOAPP    string
	SECRET_KEY   string
	TIEMPO       string
	HORABACK     string
	MINUTOBACK   string
)

// Constantes para valores por defecto
const (
	DefaultPort   = "9000"
	DefaultBackup = "0 8 * * *" // 8:00 AM por defecto
)

func FetchVars() {
	// Cargar archivo .env
	if err := godotenv.Load("local.env"); err != nil {
		utils.CreateLog("No se encontró el archivo local.env, usando variables de entorno del sistema")
	}

	// Asignar variables con valores por defecto
	// Render y otras plataformas de hosting usan la variable PORT. Le damos prioridad.
	PUERTOAPP = getEnv("PORT", getEnv("PUERTO", DefaultPort))
	HORABACK = getEnv("HORABACK", "8")
	MINUTOBACK = getEnv("MINUTOBACK", "0")

	// Resto de variables
	DB_SERVER = os.Getenv("SERVIDOR")
	DB_USER = os.Getenv("USUARIO")
	DB_PASSWORD = os.Getenv("CLAVE")
	DB_NAME = os.Getenv("BASEDEDATOS")
	DB_PORT = os.Getenv("PUERTOBD")
	DB_POOL_MODE = os.Getenv("pool_mode")
	DB_SSL_MODE = getEnv("SSL_MODE", "disable") // Por defecto 'disable' para desarrollo local
	SECRET_KEY = os.Getenv("SECRET_KEY")
	TIEMPO = os.Getenv("TIEMPO")

	// Resolve hostname to IPv4 to avoid potential IPv6 connection issues
	dbHost := DB_SERVER
	ips, err := net.LookupIP(dbHost)
	if err == nil {
		for _, ip := range ips {
			// Check for IPv4 address
			if ip.To4() != nil {
				dbHost = ip.String() // Found an IPv4 address, use it
				utils.CreateLog(fmt.Sprintf("Resolved database host %s to IPv4 address: %s", DB_SERVER, dbHost))
				break
			}
		}
	} else {
		utils.CreateLog(fmt.Sprintf("Could not resolve host %s: %v. Using original hostname.", DB_SERVER, err))
	}

	// Construir cadena de conexión (sin comillas en la contraseña y con sslmode configurable)
	dbinfo = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT, DB_SSL_MODE)
	if DB_POOL_MODE != "" {
		dbinfo = fmt.Sprintf("%s pool_mode=%s", dbinfo, DB_POOL_MODE)
	}
}

// Función auxiliar para obtener variables con valor por defecto
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// Nuevo: Obtener intervalo de tiempo para tareas programadas
func GetTaskInterval() time.Duration {
	value := os.Getenv("TIMEBACKGROUND")
	unit := os.Getenv("TIMETYPE")

	interval, err := time.ParseDuration(value + unit)
	if err != nil {
		utils.CreateLog(fmt.Sprintf("Intervalo inválido, usando valor por defecto (30m): %v", err))
		return 30 * time.Minute
	}
	return interval
}
