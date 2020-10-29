package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// la variable generador puede apuntar a cualquier función sin argumentos
	// que retorne un int
	var generador func() int

	generador = func() int {
		return 0
	}
	fmt.Println("generador ceros:", generador())

	// los literales de función pueden acceder a valores fuera de su ámbito
	cuenta := 0
	generador = func() int {
		cuenta++
		return cuenta
	}
	fmt.Println("contador generador:", generador(), generador(), generador())

	// generador puede apuntar a cualquier función que ya exista, si ésta
	// comparte la misma signatura
	generador = rand.Int
	fmt.Println("generador aleatorio:", generador())

	// caso que en determinados momentos podrá ser útil:
	// invocar a una función inmediatamente después de definirla
	func(message string) {
		fmt.Println(message)
	}("goodbye!")
}
