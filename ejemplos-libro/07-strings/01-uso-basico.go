package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "hola mundo"
	fmt.Printf("%q tiene longitud %v\n", str1, len(str1))

	str2 := "hola 世界"
	fmt.Printf("%q tiene longitud %v\n", str2, len(str2))

	fmt.Printf("(de hecho, tiene %v runas)\n", utf8.RuneCountInString(str2))

	barr := []byte("creando una cadena desde un array de bytes")
	str3 := string(barr)

	// Esto daría error de compilación! Las cadenas son inmutables
	//str3[0] = 'C'

	// Puedes convertir de string a un vector de bytes
	str4 := []byte(str3)
	str4[0] = 'C'

	// Pero no podrás asignar runas!! (constante 12416 no cabe en un byte)
	// str4[1] = 'む'

	fmt.Println(string(str4))

	str5 := []rune(str3)
	str5[0] = 'む'
	fmt.Println(string(str5))

	fmt.Println(`El uso del acento grave como delimitador
permite escribir literales de string que pueden
ocupar varias líneas e ignorar los códigos de
elusión, como por ejemplo \n o \t`)

	fmt.Println(`Este texto se imprimirá
en varias líneas
y no necesitarás eludir algunos
carácteres, como " or \`)
}
