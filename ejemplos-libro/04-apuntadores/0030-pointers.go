package main

import "fmt"

func main() {
	a := 1
	b := 2

	// uninitialized pointer to int
	// for safety reasons, you don't have pointer arithmetics (e.g. pointer++)
	var pointer *int
	pointer = &a

	// C-like formatted output
	fmt.Printf("Pointed value is %d\n", *pointer)

	// initialized pointer (type inference)
	p := &a
	fmt.Printf("We first point to %d", *p)
	p = &b
	fmt.Printf(" and then we point to %d\n", *p)
}
