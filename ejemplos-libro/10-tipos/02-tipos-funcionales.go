package main

import (
	"fmt"
	"math/rand"
)

type Generador func() int

func GenerarTodo(gens ...Generador) []int {
	nums := make([]int, 0, len(gens))
	for _, g := range gens {
		nums = append(nums, g())
	}
	return nums
}

func Cero() int {
	return 0
}

func Incremento() Generador {
	cuenta := 0
	return func() int {
		cuenta++
		return cuenta
	}
}

func Aleatorio(semilla int64) Generador {
	rnd := rand.NewSource(semilla)
	return func() int {
		return int(rnd.Int63())
	}
}

func main() {
	cnt := Incremento()
	rnd := Aleatorio(456)
	for i := 0; i < 5; i++ {
		fmt.Println(GenerarTodo(Cero, cnt, rnd))
	}
}
