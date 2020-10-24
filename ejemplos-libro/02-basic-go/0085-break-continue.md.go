package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for true {
		fmt.Println("Infinite loop! ... ?")
		if rand.Intn(10) == 0 {
			fmt.Println("Randomly breaking loop")
			break
		}
	}

	for i := 0 ; i < 10 ; i++ {
		if rand.Intn(5) == 0 {
			fmt.Println("Randomly skipping count")
			continue
		}
		fmt.Println("Counting: ", i)
	}
}
