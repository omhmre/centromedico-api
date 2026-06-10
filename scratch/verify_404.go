package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	url := "http://localhost:9000/putdoctores"
	
	// Payload with non-existent ID (9999)
	payload := map[string]interface{}{
		"id": 9999,
		"nombres": "Test Doctor",
		"es_medico": true,
		"monto_cita": 50.0,
		"servicios": []interface{}{},
		"days_of_week": []int{1, 2, 3, 4, 5},
	}
	
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	fmt.Printf("Status Code: %d\n", resp.StatusCode)
}
