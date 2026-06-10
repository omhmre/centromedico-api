package main

import (
	"fmt"
	"log"
	"omhmre.com/centromedico/app/domain/database"
)

func main() {
	database.FetchVars()
	db := &database.DB{}
	if err := db.Open(); err != nil {
		log.Fatalf("Error al abrir DB: %v", err)
	}
	defer db.Close()

	config, resp := db.GetEmailConfig()
	fmt.Printf("Status: %d, Mensaje: %s\n", resp.Status, resp.Mensaje)
	fmt.Printf("SMTP: %s, Puerto: %d, Usuario: %s, TLS: %v\n", config.Smtp, config.Puerto, config.Usuario, config.Tls)
    
    // También verificar si hay usuarios con correo
    rows, err := db.Conn.Query(`SELECT codigo, nombre, correo FROM seguridad.usuarios WHERE correo IS NOT NULL AND correo != '';`)
    if err != nil {
        fmt.Printf("Error al consultar usuarios: %v\n", err)
    } else {
        defer rows.Close()
        fmt.Println("Usuarios con correo:")
        for rows.Next() {
            var codigo, nombre, correo string
            rows.Scan(&codigo, &nombre, &correo)
            fmt.Printf("- %s: %s (%s)\n", codigo, nombre, correo)
        }
    }
}
