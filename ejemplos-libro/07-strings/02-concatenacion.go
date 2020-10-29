package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str1 := "El operador"
	str2 := str1 + " de suma es una forma simple y rápida " +
		"de concatenar cadenas. " +
		"Pero solo puedes concatenar cadenas, y pierde" +
		" velocidad si se concatenan varias cadenas en" +
		" distintos ámbitos"
	fmt.Println(str2)

	str3 := fmt.Sprintf(`Las funciones fmt.Sprint o fmt.Sprintf
permiten especificar un formato detallado para la
composición de cadenas (por ejemplo %.2f),
pero generalmente son más lentas que otras opciones.`, 1.1)
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
