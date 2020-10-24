package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println(
			"Calculando la respuesta a la Gran Pregunta" +
				"de la Vida, el Universo, y Todo lo Demás")
		time.Sleep(15 * time.Second)
		ch <- 42
	}()

	fmt.Println("Esperando...")
	select {
	case ret := <-ch:
		fmt.Println("Recibido:", ret)
	case <-time.After(2 * time.Second):
		fmt.Println("Error: tiempo de espera agotado")
	}

}
