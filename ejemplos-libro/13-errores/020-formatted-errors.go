package main

import (
	"fmt"
	"math"
)

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("%f does not have a real square root", num)
	}
	return math.Sqrt(num), nil
}

func main() {
	res, err := sqrt(-3)
	if err != nil {
		fmt.Println("error!", err)
		return
	}
	fmt.Println("result: ", res)
}
