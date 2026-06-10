package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"omhmre.com/centromedico/app/domain/database"
	"omhmre.com/centromedico/app/domain/models"
	"omhmre.com/centromedico/app/infrastructure/services"
)

func main() {
	// Load env
	err := godotenv.Load("local.env")
	if err != nil {
		log.Fatalf("Error loading env: %v", err)
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatalf("GEMINI_API_KEY is empty in env")
	}

	// 1. Fetch DB vars from env
	database.FetchVars()

	// 2. Open DB connection
	db := &database.DB{}
	if err := db.Open(); err != nil {
		log.Fatalf("Error opening connection to database: %v", err)
	}
	defer db.Close()

	fmt.Println("Conectado a la base de datos!")

	// 3. Setup test request
	chatReq := models.AIChatRequest{
		Message: "agendame una cita para el paciente nilio mañanas lunes con el dr montaner a las 9:00",
		History: []models.AIChatMessage{},
		Cedula:  "",
	}

	fmt.Println("Procesando chat con Gemini...")
	ctx := context.Background()
	response, err := services.ProcessAIChat(ctx, apiKey, db, chatReq)
	if err != nil {
		log.Fatalf("ProcessAIChat ERROR: %v", err)
	}

	fmt.Printf("GEMINI RESPONSE SUCCESS!\n")
	fmt.Printf("Message: %s\n", response.Response)
	if response.Widget != nil {
		fmt.Printf("Widget type: %s\n", response.Widget.Type)
		fmt.Printf("Widget data: %+v\n", response.Widget.Data)
	} else {
		fmt.Println("No widget returned.")
	}
}
