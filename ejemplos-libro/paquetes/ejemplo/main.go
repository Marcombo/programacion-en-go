package main

import (
	"fmt"

	"github.com/mariomac/analizador"
	"macias.info/go/ejemplo/hola"
)

func main() {
	fmt.Print("Cómo te llamas?: ")
	var nombre string
	fmt.Scanln(&nombre)
	fmt.Println(hola.ConNombre(nombre))

	analizador.PrintEstadistica(nombre)
}
