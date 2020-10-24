package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str1 := "The add"
	str2 := str1 + " operator is a quick an simple " +
		"way to concatenate strings. " +
		"But you can't concatenate other value different than a String."
	fmt.Println(str2)

	str3 := fmt.Sprintf(`The C-like fmt.Sprint or fmt.Sprintf functions allow
you providing a more detailed description of the data you
concatenate (e.g. %.2f), but they are usually slower.`, 1.1)
	fmt.Println(str3)

	var sb strings.Builder
	sb.WriteString("strings.Builder\n")
	l := sb.Len()
	for i := 0; i < l; i++ {
		sb.WriteRune('=')
	}
	sb.WriteString("\nEs la forma más eficiente de construir ")
	sb.WriteString("cadenas, de manera condicional o iterativa")
	sb.WriteString("\nSin embargo, solo se pueden agregar ")
	sb.Write([]byte("cadenas o porciones de bytes o runas"))
	sb.WriteString("Para añadir otros tipos, primero debes")
	sb.WriteString("transformarlos (p. ej. con strconv.Itoa(")
	sb.WriteString(strconv.Itoa(1234))
	sb.WriteByte(')')

	fmt.Println(sb.String())
}
