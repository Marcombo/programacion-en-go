package main

import "fmt"

func main() {
	// En Go, los vectores se tratan por valor
	src := [5]int{1, 2, 3, 4, 5}
	var dst [5]int

	// Eso significa que, en asignaciones, se copian
	dst = src
	src[0] = 9

	fmt.Printf("%v != %v", src, dst)
}
