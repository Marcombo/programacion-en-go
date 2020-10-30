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
	ch := make(chan int)

	go Emisor(ch)
	Receptor(ch)
}
