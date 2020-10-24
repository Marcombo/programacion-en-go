package main

import "fmt"

func add(n int, np *int) int {
	return n + *np
}

func main() {
	defer func() {
		if r := recover() ; r != nil {
			fmt.Println("Recovering panic:", r)
		}
	}()
	var nilRef *int = nil
	fmt.Println("adding 3 to the global ref: ", add(3, nilRef))

	fmt.Println("This won't be printed")
}
