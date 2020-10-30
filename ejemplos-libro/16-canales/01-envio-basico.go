package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hola"
	}()

	recibido := <-ch
	fmt.Println("He recibido:", recibido)
	// no es necesario cerrar channel
}
