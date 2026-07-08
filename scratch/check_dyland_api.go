package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("https://centromedico-api.onrender.com/pacientes")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var patients []map[string]interface{}
	json.Unmarshal(body, &patients)
	for _, p := range patients {
		nombres, ok := p["nombres"].(string)
		if ok && strings.Contains(strings.ToUpper(nombres), "DYLAND") {
			fmt.Printf("ID: %v, Cedula: %v, Nombres: %v\n", p["id"], p["cedula"], p["nombres"])
		}
	}
}
