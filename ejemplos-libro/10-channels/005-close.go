package main

import "fmt"

func main() {
	ch := make(chan string)
	close(ch)

	x := <-ch
	fmt.Println("x:", x)

	ch <- "hooola"
}
