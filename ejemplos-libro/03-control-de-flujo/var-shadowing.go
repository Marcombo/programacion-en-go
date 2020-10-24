package main

import "fmt"

func main() {
	a := 0
	b := 0

	if true {
		a := 1
		b = 1
		a++
		b++
	}

	fmt.Printf("a = %d, b = %d\n", a, b)

}
