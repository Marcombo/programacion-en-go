package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("boom")
	err2 := fmt.Errorf("algo malo sucedió: %w", err)
	fmt.Println(err2)

	causa := errors.Unwrap(err2)
	fmt.Println("la causa es", causa)
	causa = errors.Unwrap(causa)
	fmt.Println("la causa de la causa es", causa)
}
