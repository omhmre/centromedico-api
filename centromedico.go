package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron"
	"omhmre.com/centromedico/app/domain/database"
	app "omhmre.com/centromedico/app/infrastructure"
	"omhmre.com/centromedico/app/websocket"
)

func main() {
	// Configurar un logger estructurado para toda la aplicación
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Cargar variables de entorno (usando tu sistema actual)
	database.FetchVars()

	// Crear el Hub de WebSocket y ejecutarlo en segundo plano
	hub := websocket.NewHub()
	go hub.Run()

	// Inicializar aplicación
	app := app.New(hub)

	// Conexión a la base de datos
	app.DB = &database.DB{}
	if err := app.DB.Open(); err != nil {
		logger.Error("Error al abrir la conexión a la base de datos", "error", err)
		os.Exit(1)
	}
	defer func() {
		if err := app.DB.Close(); err != nil {
			logger.Error("Error al cerrar la conexión a la base de datos", "error", err)
		}
	}()

	// Configurar cron para backups (usando tus variables HORABACK y MINUTOBACK)
	c := cron.New()
	backupSchedule := database.HORABACK + " " + database.MINUTOBACK + " * * *"
	if err := c.AddFunc(backupSchedule, func() {
		logger.Info("Iniciando backup de base de datos...")
		if errDb := app.DB.BackupDatabase(); errDb != nil {
			logger.Error("Error en backup de base de datos", "error", errDb)
		} else {
			logger.Info("Backup de base de datos completado exitosamente")
		}
	}); err != nil {
		logger.Error("Error al programar el backup", "error", err)
	}
	c.Start()
	defer c.Stop()

	// Configurar servidor HTTP con graceful shutdown
	server := &http.Server{
		Addr:    ":" + database.PUERTOAPP,
		Handler: app.WrapWithCORS(app.Router), // Usar el router de la app con CORS
		// Add a timeout for the server to gracefully shut down
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Canal para manejar señales de terminación
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Iniciar servidor en goroutine
	go func() {
		logger.Info("Servidor iniciado", "puerto", database.PUERTOAPP)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Error al iniciar el servidor", "error", err)
			os.Exit(1)
		}
	}()

	// Esperar señal de terminación
	<-stop
	logger.Info("Recibida señal de apagado, iniciando cierre...")

	// Contexto con timeout para el cierre
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Error durante el cierre del servidor", "error", err)
	} else {
		logger.Info("Servidor detenido correctamente")
	}
}
