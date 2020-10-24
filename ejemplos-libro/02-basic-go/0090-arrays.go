package main

import "fmt"

func main() {
	// In Go, arrays are values
	src := [5]int{1, 2, 3, 4, 5}
	var dst [5]int

	// that means, in assignments they are copied
	dst = src
	src[0] = 9

	fmt.Printf("%v != %v", src, dst)
}
