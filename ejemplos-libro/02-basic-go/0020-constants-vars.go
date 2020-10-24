package main

import "fmt"

// global constant. Go does not use the UPPER_SNAKE_CASE
const Pi = 3.1415

// uninitialized (default to 0) global variable
var Global int

// initialized global variable (type inference)
var GlobalCounter = 2

func main() {

	GlobalCounter = int32(6)

	local := "Preferred way to define an initialized variable"

	local = "don't confuse short initialization := with assignment = operator"

	fmt.Print("if you define but don't read a local variable,")
	fmt.Println("you will get a compiler error")
	fmt.Println(local)
}
