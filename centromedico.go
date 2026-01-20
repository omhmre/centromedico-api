package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron"
	"omhmre.com/centromedico/app/domain/database"
	"omhmre.com/centromedico/app/domain/utils"
	app "omhmre.com/centromedico/app/infrastructure"
	"omhmre.com/centromedico/app/websocket"
)

func main() {
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
		utils.CreateLog(fmt.Sprintf("Error al abrir la conexión a la base de datos: %v", err))
		os.Exit(1)
	}
	defer func() {
		if err := app.DB.Close(); err != nil {
			utils.CreateLog(fmt.Sprintf("Error al cerrar la conexión a la base de datos: %v", err))
		}
	}()

	// Configurar cron para backups (usando tus variables HORABACK y MINUTOBACK)
	c := cron.New()
	backupSchedule := database.HORABACK + " " + database.MINUTOBACK + " * * *"
	if err := c.AddFunc(backupSchedule, func() {
		if err := app.DB.BackupDatabase(); err != nil {
			utils.CreateLog(fmt.Sprintf("Error en backup de base de datos: %v", err))
		} else {
			utils.CreateLog("Backup de base de datos completado exitosamente")
		}
	}); err != nil {
		utils.CreateLog(fmt.Sprintf("Error al programar el backup: %v", err))
	}
	c.Start()
	defer c.Stop()

	// Configurar servidor HTTP con graceful shutdown
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", database.PUERTOAPP),
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
		utils.CreateLog(fmt.Sprintf("Servidor iniciado en el puerto %s", database.PUERTOAPP))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.CreateLog(fmt.Sprintf("Error al iniciar el servidor: %v", err))
			os.Exit(1)
		}
	}()

	// Esperar señal de terminación
	<-stop
	utils.CreateLog("Recibida señal de apagado, iniciando cierre...")

	// Contexto con timeout para el cierre
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		utils.CreateLog(fmt.Sprintf("Error durante el cierre del servidor: %v", err))
	} else {
		utils.CreateLog("Servidor detenido correctamente")
	}
}
