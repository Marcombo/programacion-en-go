package main

import "fmt"

func TareaAsincrona() <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		fmt.Println("haciendo alguna cosa en paralelo...")
		for i := 0; i < 3; i++ {
			fmt.Println(i, "...")
		}
		fmt.Println("finalizada tarea en paralelo")

		close(ch) // any receive operation will be unblocked
	}()
	return ch
}

func main() {

	espera := TareaAsincrona()

	<-espera // channel output will be ignored

	fmt.Println("programa finalizado")
}
