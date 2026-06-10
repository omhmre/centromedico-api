package main

import (
	"fmt"
	"log"
	"omhmre.com/centromedico/app/domain/database"
	"omhmre.com/centromedico/app/domain/models"
)

func main() {
	database.FetchVars()
	db := &database.DB{}
	if err := db.Open(); err != nil {
		log.Fatalf("Error al abrir DB: %v", err)
	}
	defer db.Close()

	mailData := models.MailSend{
		To:      "omhmre@gmail.com", // Enviamos a la cuenta configurada como prueba
		Subject: "Prueba de Conexión SMTP",
		Body:    "Este es un correo de prueba desde el backend de Centro Médico.",
	}

	fmt.Println("Intentando enviar correo...")
	err := db.SendMail(mailData)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Println("ÉXITO: Correo enviado correctamente.")
	}
}
