package main

import "fmt"

func main() {
	// Go programs usually use slices instead of arrays,
	// which are treated as references, and have many advantages:

	s := []int{1, 2, 3, 4}
	fmt.Println("before =", s)

	s = append(s, 5, 6, 7)
	fmt.Println("after =", s)

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s1 = append(s1, s2...)
	fmt.Println("concat =", s1)

	// 1- can grow dynamically
	var sl []int
	fmt.Printf("length: %v. capacity: %v\n", len(sl), cap(sl))
	sl = append(sl, 1, 2, 3, 4)
	fmt.Printf("length: %v. capacity: %v\n", len(sl), cap(sl))
	sl = append(sl, 5)
	fmt.Printf("length: %v. capacity: %v\n", len(sl), cap(sl))

	// 2- you can define an initial capacity to minimize resizing costs
	sl = make([]int, 0, 10) // declare a slice of size 0 and capacity 10
	fmt.Printf("length: %v. capacity: %v\n", len(sl), cap(sl))
	sl = append(sl, 1)
	fmt.Printf("length: %v. capacity: %v\n", len(sl), cap(sl))

	// 3- you can provide views
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
