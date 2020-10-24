package main

import (
	"encoding/xml"
	"fmt"
)

type Animal struct {
	Nombre string `xml:"nombre"`
	Tipo   string `xml:"tipo"`
}

func main() {
	a := Animal{Nombre: "Perro", Tipo: "Mamífero"}
	b, err := xml.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
