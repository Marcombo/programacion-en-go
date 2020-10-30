package funciones

import "fmt"

func Factorial(n uint64) uint64 {
	f := uint64(1)
	for i := uint64(2); i <= n; i++ {
		f *= i
	}
	return f
}

func NoProbado() {
	fmt.Println("hola!")
}
