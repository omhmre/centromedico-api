package utils

import (
	"log"
	"os"
)

// Logger is the instance of the logger
var Logger *log.Logger

// InitLogger initializes the logger
func InitLogger() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	Logger = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
}

// LogInfo logs informational messages
func LogInfo(message string) {
	Logger.Println("INFO: " + message)
}

// LogError logs error messages
func LogError(message string) {
	Logger.Println("ERROR: " + message)
}
