package main

import (
	"fmt"
	"sync"
)

func main() {
	const numTareas = 3

	wg := sync.WaitGroup{}
	wg.Add(numTareas)

	for i := 0 ; i < numTareas; i++ {
		numTarea := i // "i" is in the shared scope for all the goroutines
		go func() {
			defer wg.Done()
			// emphasize that this does not guarantee the order of completion. E.g.:
			// Running parallel task 2
			// Running parallel task 0
			// Running parallel task 1
			fmt.Println("Ejecutando tarea", numTarea)
		}()
	}

	wg.Wait()
	fmt.Println("Completadas todas las tareas. Finalizando")
}