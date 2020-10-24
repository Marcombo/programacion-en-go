package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "hello world"
	fmt.Printf("%q has length %v\n", str1, len(str1))

	str2 := "hello 世界"
	fmt.Printf("%q has length %v\n", str2, len(str2))

	fmt.Printf("(actually, it has %v runes)\n", utf8.RuneCountInString(str2))

	str3 := string("creating a string from a byte array")

	// compilation error! Strings are immutable. (uncomment)
	//str3[0] = 'C'

	// you can convert from string to byte array
	str4 := []byte(str3)
	str4[0] = 'C'

	// but you cannot assign UTF8 runes! (uncomment, constant 12416 overflows byte)
	// str4[1] = 'む'

	fmt.Println(string(str4))

	str5 := []rune(str3)
	str5[0] = 'む'
	fmt.Println(string(str5))

	fmt.Println(`The grave accent delimiter allow writing string literals
which can occupy several lines and ignore escape codes. E.g. \n or \t`)

	fmt.Println(`This text will be printed
in four different lines
and you don't need to escape
characters like the new line, " or \`)
}
