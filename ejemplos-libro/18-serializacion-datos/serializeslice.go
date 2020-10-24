package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	listaJson := []byte(`["hola", "que", "tal"]`)
	var lista []string
	if err := json.Unmarshal(listaJson, &lista); err != nil {
		panic(err)
	}
	fmt.Println("deserializado:", lista)
	lista = append(lista, "amigo")
	listaJson2, err := json.Marshal(lista)
	if err != nil {
		panic(err)
	}
	fmt.Println("serializado:", string(listaJson2))
}
