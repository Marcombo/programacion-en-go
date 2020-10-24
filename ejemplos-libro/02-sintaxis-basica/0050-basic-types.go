package main

import "fmt"

func main() {
	// Go by default uses a "generic", platform-dependent types
	a := 10
	fmt.Printf("a (%T) = %v\n", a, a)

	// explicit type size and sign
	b := int8(10)
	var c uint8 = 10

	fmt.Printf("b (%T) = %v\n", b, b)
	fmt.Printf("c (%T) = %v\n", c, c)

	// explicit type conversion (otherwise, compile error)
	a = int(b)
	fmt.Printf("a (%T) = %v\n", a, a)
}
