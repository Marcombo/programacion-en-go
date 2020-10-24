package main

import "fmt"

func Emisor(ch chan<- int) {
	for i := 1; i <= 3; i++ {
		ch <- i
		fmt.Println(i, "enviado correctamente")
	}
}

func Receptor(ch <-chan int) {
	for {
		num := <-ch
		fmt.Println("recibido:", num)
	}
}

func main() {
	ch := make(chan int, 10) //buffered channel

	// With a buffered channel, receiver and Emisor don't strictly need to go in
	// different goroutines
	go Receptor(ch)
	Emisor(ch)
}
