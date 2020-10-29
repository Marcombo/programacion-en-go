package main

import "fmt"

func main() {
	a := 1
	b := 2

	// declaración e inicialización por separado
	var puntero *int
	puntero = &a

	// Lectura del valor de puntero
	fmt.Printf("valor apuntado: %d\n", *puntero)

	// declaración e inicialización con inferencia de tipos
	// forma recomendada
	p := &a
	fmt.Printf("Primero apuntamos a %d\n", *p)

	// apuntando a una nueva variable
	p = &b
	fmt.Printf("Luego apuntamos a %d\n", *p)
}
