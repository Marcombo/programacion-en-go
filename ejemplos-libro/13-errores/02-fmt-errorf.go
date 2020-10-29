package main

import (
	"fmt"
	"math"
)

func raizCuadrada(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("%f no tiene raiz cuadrada real", num)
	}
	return math.Sqrt(num), nil
}

func main() {
	res, err := raizCuadrada(-3)
	if err != nil {
		fmt.Println("error!", err)
		return
	}
	fmt.Println("resultado: ", res)
}
