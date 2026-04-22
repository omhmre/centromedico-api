//go:build ignore
package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"omhmre.com/centromedico/app/domain/models"
)

func main() {
	jsonStr := `{
  "id_paciente": 491,
  "fecha": "2026-04-21T07:33:35.746803",
  "id_doctor": 0,
  "diagnostico": "DIAGNOSTICO FICTICIO PARA PRUEBAS",
  "evolucion": "MOTIVO DE CONSULTA: Cefalea hemicraneana persistente y parestesias en extremidad superior derecha.\r\n\r\nENFERMEDAD ACTUAL: Paciente refiere cuadro clínico de 3 meses de evolución caracterizado por cefalea pulsátil, escala EVA 7/10, asociada a fotofobia y episodios de adormecimiento transitorio en mano derecha. No refiere pérdida de conciencia.\r...",
  "plan": "",
  "recomendaciones": "",
  "contenido": "...",
  "entregado": false,
  "fecha_entrega": null,
  "modificado_post_entrega": false,
  "usuario_operacion": "Peggy Colina"
}`

	var i models.InformeMedico
	err := json.NewDecoder(strings.NewReader(jsonStr)).Decode(&i)
	if err != nil {
		fmt.Printf("DECODE ERROR: %v\n", err)
		return
	}
	fmt.Println("DECODE SUCCESS")
}
