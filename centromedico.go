package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron"
	"omhmre.com/centromedico/app/domain/database"
	"omhmre.com/centromedico/app/infrastructure"
	"omhmre.com/centromedico/app/websocket"
)

func main() {
	// Configurar un logger estructurado para toda la aplicación
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if err := run(logger); err != nil {
		logger.Error("La aplicación terminó con un error fatal", "error", err)
		os.Exit(1)
	}
}

func run(logger *slog.Logger) error {
	// Contexto base para controlar la propagación del apagado
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 1. Cargar variables de entorno
	database.FetchVars()

	// 2. Inicializar base de datos
	db := &database.DB{}
	if err := db.Open(); err != nil {
		return fmt.Errorf("error al abrir la conexión a la base de datos: %w", err)
	}
	db.InitAbonosTables()
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("Error al cerrar la conexión a la base de datos", "error", err)
		}
	}()

	// 3. Crear el Hub de WebSocket y ejecutarlo en segundo plano
	hub := websocket.NewHub()
	go hub.Run()

	// 4. Inicializar aplicación usando inyección de dependencias
	application := infrastructure.New(hub)
	application.DB = db

	// 5. Configurar cron para backups
	c := cron.New()
	// v1.2.0 de robfig/cron requiere 6 campos (Segundos, Minutos, Horas, DiaMes, Mes, DiaSemana)
	backupSchedule := fmt.Sprintf("0 %s %s * * *", database.MINUTOBACK, database.HORABACK)

	if err := c.AddFunc(backupSchedule, func() {
		logger.Info("Iniciando backup de base de datos...")
		if errDb := application.DB.BackupDatabase(); errDb != nil {
			logger.Error("Error en backup de base de datos", "error", errDb)
			return
		}
		logger.Info("Backup de base de datos completado exitosamente")
	}); err != nil {
		return fmt.Errorf("error al programar el backup: %w", err)
	}
	c.Start()
	defer c.Stop()

	// 6. Configurar servidor HTTP con graceful shutdown y time-outs más robustos contra ataques
	server := &http.Server{
		Addr:              ":" + database.PUERTOAPP,
		Handler:           application.WrapWithCORS(application.Router),
		ReadHeaderTimeout: 5 * time.Second,   // CRÍTICO evitar ataques Slowloris
		ReadTimeout:       15 * time.Second,  // Tiempo máximo para leer la petición
		WriteTimeout:      15 * time.Second,  // Tiempo máximo para escribir la respuesta
		IdleTimeout:       120 * time.Second, // Tiempo máximo de Keep-Alive inactivo
	}

	// 7. Manejo Concurrente de Señales y Servidor
	serverErrors := make(chan error, 1)

	// Goroutine que inicia el servidor HTTP
	go func() {
		logger.Info("Servidor iniciado", "puerto", database.PUERTOAPP)
		serverErrors <- server.ListenAndServe()
	}()

	// Canal para recibir señales de terminación del sistema operativo
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// 8. Sincronización de eventos (Select Pattern)
	select {
	case err := <-serverErrors:
		// Se disparó un error (ej. puerto ocupado o cerrado abrúptamente)
		if err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("servidor falló al escuchar: %w", err)
		}

	case sig := <-shutdown:
		// Se capturó una señal del OS
		logger.Info("Recibida señal de apagado, iniciando cierre...", "señal", sig.String())

		// Le damos un máximo de 15 segundos al server para terminar las peticiones activas
		shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 15*time.Second)
		defer shutdownCancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			// Si hubo timeout o falló el shutdown ordenado forzamos el cierre
			if errClose := server.Close(); errClose != nil {
				return fmt.Errorf("fallo forzando el cierre del servidor: %w", errClose)
			}
			return fmt.Errorf("servidor detenido forzosamente por error en el apagado: %w", err)
		}
	}

	logger.Info("Servidor detenido correctamente")
	return nil
}
