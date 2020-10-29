package main

import "fmt"

func main() {
	// Generalmente usaremos porciones en vez de vectores
	// son tratadas por referencia, y pueden tener un tamaño
	// variable

	s := []int{1, 2, 3, 4}
	fmt.Println("antes =", s)

	s = append(s, 5, 6, 7)
	fmt.Println("despues =", s)

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s1 = append(s1, s2...)
	fmt.Println("concatenación =", s1)

	// 1- crecimiento dinámico
	var sl []int
	fmt.Printf("length: %v. capacity: %v\n", len(sl), cap(sl))
	sl = append(sl, 1, 2, 3, 4)
	fmt.Printf("length: %v. capacity: %v\n", len(sl), cap(sl))
	sl = append(sl, 5)
	fmt.Printf("length: %v. capacity: %v\n", len(sl), cap(sl))

	// 2- puedes definir una capacidad inicial para minimizar costes de redimensionado
	sl = make([]int, 0, 10) // tamaño 0, capacidad 10
	fmt.Printf("length: %v. capacity: %v\n", len(sl), cap(sl))
	sl = append(sl, 1)
	fmt.Printf("length: %v. capacity: %v\n", len(sl), cap(sl))

	// 3- se pueden proveer vistas
	base := []int{1, 0, 3, 4, 5}
	fmt.Println("base:", base)

	vista1 := base[1:3] // incluye elementos 1 y 2
	fmt.Println("vista 1:", vista1)
	vista1[0] = 2 // equivalente a base[1] = 2
	fmt.Println("base:", base)

	vista2 := base[2:] // desde índice 2 hasta el final
	fmt.Println("vista 2:", vista2)

	vista3 := base[:3] // desde el inicio hasta indice 2
	fmt.Println("vista 3:", vista3)

	vista4 := base[:] // vista de inicio a fin
	fmt.Println("vista 4:", vista4)

}
