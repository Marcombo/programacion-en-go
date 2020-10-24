package main

import (
	"fmt"
	"math/rand"
)

func Dado(ch chan<- int, veces int) chan int {
	numeros := make(chan int, veces)
	for i := 0; i < veces; i++ {
		fmt.Println("Agitando el dado...")
		ch <- rand.Intn(6) + 1
	}
	close(numeros)
}

func main() {

}
