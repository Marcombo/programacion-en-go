package main

import (
	"errors"
	"fmt"
)

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return dividend / divisor, nil
}

func main() {
	div, err := divide(7, 0)
	if err != nil {
		fmt.Println("error!", err)
		return
	}
	fmt.Println("result: ", div)
}
