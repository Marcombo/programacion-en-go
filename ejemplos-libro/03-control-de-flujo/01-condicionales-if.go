package main

import (
	"fmt"
	"math/rand"
)

func main() {
	valor := rand.Int() % 2

	// El if-else clásico de otros lenguajes.
	// No se necesita paréntesis alrededor de la condición
	if valor == 0 {
		fmt.Println("valor par")
	} else {
		fmt.Println("valor impoar")
	}

	// se puede limitar el ámbito de una variable si ésta
	// se declara en el mismo if. Delante de la condición,
	// separada por punto y coma
	if otroValor := rand.Int(); otroValor%2 == 0 {
		fmt.Printf("Otro valor (%d) es par\n", otroValor)
	} else {
		fmt.Printf("Otro valor (%d) es impar\n", otroValor)
	}
}
