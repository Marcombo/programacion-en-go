package main

import "fmt"

// Constante global
const Pi = 3.1415

// variable global no inicializada (tomará el valor 0)
var Global int

// variable global inicializada
// (por inferencia de tipos, será un int)
var GlobalCounter = 2

// Constantes globales agrupadas
const (
	TimeoutMS  = 1000
	MaxRetries = 4
	FailOnErr  = true
)

func main() {

	local := "Forma preferida de declarar una variable inicializada"

	local = "no confundir el operador := de declaración con el = de asignación"

	fmt.Print("si defines pero no usas una variable local,")
	fmt.Println("obtendrías un error de compilación")
	fmt.Println(local)
}
