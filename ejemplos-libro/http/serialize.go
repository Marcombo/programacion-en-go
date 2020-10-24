package main

import (
	"encoding/json"
	"fmt"
)

type Futbolista struct {
	Nombre     string   `json:"nombre"`
	Nacimiento int      `json:"nacimiento"`
	Equipos    []string `json:"equipos"`
}

var cr = Futbolista{
	Nombre:     "Carles Reixach",
	Nacimiento: 1947,
	Equipos: []string{
		"F.C. Barcelona",
		"Selección Española de Fútbol",
	},
}

func main() {
	documento := []byte(`{
		"nombre": "Paco Garcia",
		"nacimiento": "un dia cualquiera",
		"liga": "Liga Peruana"
	}`)
	var f Futbolista
	err := json.Unmarshal(documento, &f)
	if err != nil {
		fmt.Println("Error decodificando:", err)
		return
	}
	fmt.Printf("decodificado: %#v\n", f)
}
