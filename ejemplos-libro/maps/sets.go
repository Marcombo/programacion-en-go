package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Los números ganadores de la lotería son:")
	numeros := map[int]struct{}{}
	for len(numeros) < 6 {
		n := rand.Intn(49) + 1
		if _, ok := numeros[n]; !ok {
			// si el numero está repetido, no se muestra
			numeros[n] = struct{}{}
			fmt.Println("El", n, "...")
		}
	}
	fmt.Println("Felicidades a los premiados!")
}
