package main

import (
	"fmt"
	"sync"
)

func Engullidor(nombre string, dulces <-chan string) {
	for dulce := range dulces {
		fmt.Println(nombre, "come", dulce)
	}
	fmt.Println(nombre, ": No más dulces? Adiós!")
}

func main() {
	dulces := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(3)

	for _, nom := range []string{"Marcos", "Aina", "Judit"} {
		n := nom
		go func() {
			defer wg.Done()
			Engullidor(n, dulces)
		}()
	}

	// There is no guarantee of a fair delivery
	dulces <- "Donut"
	dulces <- "Crusán"
	dulces <- "Ensaimada"
	dulces <- "Pestiño"

	close(dulces) // Important! or Eater goroutines won't end

	// don't need to sleep... just wait for completions
	wg.Wait()
	fmt.Println("Todos los dulces repartidos y engullidos")
}
