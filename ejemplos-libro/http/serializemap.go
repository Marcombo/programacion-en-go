package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	documento := []byte(`{
		"receta": "Huevos Fritos",
		"ingredientes": [
			"Huevos",
			"Aceite",
			"Sal"
		] 
	}`)
	destino := map[string]interface{}{}
	err := json.Unmarshal(documento, &destino)
	if err != nil {
		fmt.Println("Error decodificando:", err)
		return
	}
	fmt.Println("decodificado:")
	for c, v := range destino {
		fmt.Printf("%s --> %v\n", c, v)
	}
}
