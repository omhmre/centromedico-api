package utils

import (
	"fmt"
	"log"
	"os"
)

func CreateLog(message string) {
	const file = "log.txt"
	// Open the file in append mode
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a logger that writes to the file
	logger := log.New(f, "", log.LstdFlags)

	// Log the message with a timestamp
	logger.Println(fmt.Sprintf("%v ", message))
}

func ClearLog() {
	file, err := os.OpenFile("log.txt", os.O_RDWR, 0666)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		log.Println(err)
	}
}
