package main

import (
	"fmt"
	"math/rand"
)

func main() {
	someValue := rand.Int() % 2

	// Classic if-else's work as in other language. However, no parentheses are needed
	if someValue == 0 {
		fmt.Println("Some value is even")
	} else {
		fmt.Println("Some value is odd")
	}

	// you can limit the scope of a variable if you put its initialization in the if,
	// before the condition, separated by a semicolon
	if otherValue := rand.Int(); otherValue%2 == 0 {
		fmt.Printf("Other value (%d) is even\n", otherValue)
	} else {
		fmt.Printf("Other value (%d) is odd\n", otherValue)
	}

	// someValue exists at this point, otherValue doesn't
}
