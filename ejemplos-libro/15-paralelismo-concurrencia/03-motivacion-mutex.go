package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Suma todos los elementos
func Suma(porcion []int) int {
	total := 0
	for _, n := range porcion {
		total += n
	}
	return total
}

func main() {
	tareasParalelas := runtime.GOMAXPROCS(0)

	v := []int{0, 1, 3, 1, 0, 7, 8, 9, 3, 3, 0, 2}

	// experimento: lo repetimos indefinidamente, y el panic
	// final acabará lanzándose por causa de una condición de carrera
	for true {
		wg := sync.WaitGroup{}
		wg.Add(tareasParalelas)

		totalSuma := 0
		for t := 0; t < tareasParalelas; t++ {
			s := t
			go func() {
				defer wg.Done()
				inicio := s * len(v) / tareasParalelas
				fin := (s + 1) * len(v) / tareasParalelas
				suma := Suma(v[inicio:fin])
				totalSuma += suma
			}()
		}

		wg.Wait()
		if totalSuma != 37 {
			panic(fmt.Sprint("totalSuma: ", totalSuma))
		}
	}
	//fmt.Println("total sum: ", totalSuma)
}
