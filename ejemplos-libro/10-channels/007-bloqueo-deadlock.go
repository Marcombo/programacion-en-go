package main

import "fmt"

func main() {
	ch := make(chan string, 5)
	ch <- "hola"
	recibido := <-ch
	fmt.Println(recibido)
}
