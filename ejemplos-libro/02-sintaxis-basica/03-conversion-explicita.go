package main

import "fmt"

func main() {
	var numeroGrande int64
	var numeroPequeño int8 = 21

	// Go no soporta conversiones implícitas de tipos
	// debe, realizarse explicitamente
	numeroGrande = int64(numeroPequeño)

	fmt.Println("numeroGrande, smallnum:", numeroGrande, numeroPequeño)

}
