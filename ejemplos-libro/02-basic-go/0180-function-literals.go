package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// variable generator may hold any function without arguments returning an int
	var generator func() int

	generator = func() int {
		return 0
	}
	fmt.Println("zero generator:", generator())

	// function literals may access values out of its scope
	count := 0
	generator = func() int {
		count++
		return count
	}
	fmt.Println("counter generator:", generator(), generator(), generator())

	// you can point to any existing function with the same signature
	generator = rand.Int
	fmt.Println("random generator:", generator())

	// weird case: directly invoking a function literal (very useful when combined with goroutines and defer)
	func(message string) {
		fmt.Println(message)
	}("goodbye!")
}
