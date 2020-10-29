package main

import "fmt"

func Suma(a, b int) int {
	return a + b
}

func Multiplica(a, b int) int {
	return a * b
}

func main() {
	var operador func(int, int) int
	operador = Suma
	fmt.Println("suma =", operador(3, 4))
	operador = Multiplica
	fmt.Println("multiplica =", operador(3, 4))
}
