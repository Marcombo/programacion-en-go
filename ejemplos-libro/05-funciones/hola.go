package main

import "fmt"

func Hola(nombre string, apellido string) {
	fmt.Printf("Hola, %s %s!\n", nombre, apellido)
}

func main() {
	Hola("Marta", "García")
	Hola("Juan", "Martínez")
}
