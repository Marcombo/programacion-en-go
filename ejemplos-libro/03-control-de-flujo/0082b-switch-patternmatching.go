package main

import (
	"fmt"
)

func main()  {

	fmt.Print("Write a char: ")
	var c int8
	fmt.Scanf("%c", &c)

	switch {
	case c >= 'A' && c <= 'Z':
		fmt.Println("Uppercase letter")
	case c>= 'a' && c <= 'z':
		fmt.Println("Lowercase letter")
	case c >= '0' && c <= '9':
		fmt.Println("Digit")
	default:
		fmt.Println("Not a letter nor a digit")
	}
}
