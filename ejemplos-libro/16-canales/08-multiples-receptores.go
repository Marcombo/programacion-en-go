package main

import (
	"fmt"
	"time"
)

func Engullidor(nombre string, dulces <-chan string) {
	for {
		dulce := <-dulces
		fmt.Println(nombre, "come", dulce)
	}
}

func main() {
	dulces := make(chan string, 10)

	go Engullidor("Marcos", dulces)
	go Engullidor("Aina", dulces)
	go Engullidor("Judit", dulces)

	dulces <- "Donut"
	time.Sleep(time.Second)
	dulces <- "Crusán"
	time.Sleep(time.Second)
	dulces <- "Ensaimada"
	time.Sleep(time.Second)

	time.Sleep(2 * time.Second)
}
