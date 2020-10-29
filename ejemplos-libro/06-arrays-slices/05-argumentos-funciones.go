package main

import "fmt"

func Cero(vec [3]int, porc []int) {
	vec[0] = 0
	if len(porc) > 0 {
		porc[0] = 0
	}
}

func main() {
	v := [3]int{1, 2, 3}
	p := []int{1, 2, 3}
	Cero(v, p)
	fmt.Println("vector:", v)
	fmt.Println("porción:", p)
}
