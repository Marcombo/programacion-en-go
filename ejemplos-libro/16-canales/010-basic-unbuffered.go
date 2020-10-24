package main

import "fmt"

const nums = 3

func Emisor(ch chan<- int) {
	for i := 1; i <= nums; i++ {
		ch <- i
		fmt.Println(i, "enviado correctamente")
	}
}

func Receptor(ch <-chan int) {
	for i := 1; i <= nums; i++ {
		num := <-ch
		fmt.Println("recibido:", num)
	}
}

func main() {
	ch := make(chan int) //unbuffered channel

	go Emisor(ch)
	Receptor(ch) // Receiver and sender need to go in different goroutines
}
