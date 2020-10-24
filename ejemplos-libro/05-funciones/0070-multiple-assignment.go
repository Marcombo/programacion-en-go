package main

import "fmt"

func oneTwo() (int, int) {
	return 1, 2
}

func main() {
	a, b := oneTwo()
	fmt.Printf("a = %v, b = %v\n", a, b)

	// a handy way to swap values
	a, b = b, a
	fmt.Printf("a = %v, b = %v\n", a, b)
}