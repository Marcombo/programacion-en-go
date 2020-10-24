package main

import "fmt"

func muestra(cadena *string) {
	fmt.Println(*cadena)
	panic("ha sucedido algo realmente malo")
}

func main() {
	muestra(nil)

	fmt.Println("este mensaje no se mostrará")

	p := recover()
	fmt.Println("recovering from panic: ", p)
}
