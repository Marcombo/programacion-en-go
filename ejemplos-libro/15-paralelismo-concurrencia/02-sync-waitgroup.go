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
		// "i" está compartida por todas las gorrutinas, así que
		// la copiamos en una variable de ámbito exclusivo para cada
		// gorrutina
		numTarea := i
		go func() {
			defer wg.Done()
			// no se garantiza el orden de completado
			fmt.Println("Ejecutando tarea", numTarea)
		}()
	}

	wg.Wait()
	fmt.Println("Completadas todas las tareas. Finalizando")
}