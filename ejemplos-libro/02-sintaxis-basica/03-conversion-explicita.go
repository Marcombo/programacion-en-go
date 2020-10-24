package main

import "fmt"

func main() {
	var bigNum int64
	var smallNum int8 = 21

	// golang does not do implicit type conversions
	bigNum = int64(smallNum)

	fmt.Println("bignum, smallnum:", bigNum, smallNum)

}
