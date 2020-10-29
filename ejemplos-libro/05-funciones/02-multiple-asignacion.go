package main

import "fmt"

func unoDos() (int, int) {
	return 1, 2
}

func main() {
	a, b := unoDos()
	fmt.Printf("a = %v, b = %v\n", a, b)

	// intercamabiar dos valores
	a, b = b, a
	fmt.Printf("a = %v, b = %v\n", a, b)
}
