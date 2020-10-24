package main

import "fmt"

type Name struct {
	FirstName string
	Surname   string
}

func main() {
	// the empty interface may be used as an equivalent to the Java Object class
	var anything interface{}

	anything = 1
	anything = true
	anything = "hello"
	anything = Name{FirstName: "Alan", Surname: "Turing"}

	fmt.Println(anything)
}
