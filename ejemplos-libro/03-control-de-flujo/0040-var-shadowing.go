package main

import "fmt"

func main() {
	value := 3

	if true {
		value := 100
		fmt.Printf("Cool! our first true condition! ")
		fmt.Printf("I'm changing the value to %d\n", value)
	}

	if value != 100 {
		fmt.Println("Wait! Should'nt value be 100?. I must" +
			" have mixed declaration and assignment operators again")
	}
}
