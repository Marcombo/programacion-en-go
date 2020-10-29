package main

// Se puede usar el identificador vacío para evitar errores de
// compilación si queremos importar un paquete que no vayamos
// a usar
import _ "fmt"

func main() {
	// El identificador vacío te permite ignorar algunos valores de retorno
	a, _ := unoDos()

	// El identificador vacío también te permite ignorar errores de compilación
	// para variables sin usar
	_ = a
}