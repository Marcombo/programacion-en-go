package main

import "fmt"

func add(n int, np *int) int {
	// si np == nil, lanzará pánico
	return n + *np
}

func main() {
	defer func() {
		if r := recover() ; r != nil {
			fmt.Println("Recuperándonos de pánico:", r)
		}
	}()
	var nilRef *int = nil
	fmt.Println("*nilRef + 3: ", add(3, nilRef))

	fmt.Println("Esto no se mostrará")
}
