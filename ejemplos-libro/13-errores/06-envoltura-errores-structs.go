package main

import (
	"errors"
	"fmt"
)

type ErrorTV struct {
	Causa error
}

func (e ErrorTV) Error() string {
	return fmt.Sprint("problema con la TV:", e.Causa)
}

func (e ErrorTV) Unwrap() error {
	return e.Causa
}

type ErrorComponente struct {
	Nombre string
}

func (e ErrorComponente) Error() string {
	return "fallo de un componente:" + e.Nombre
}

func main() {
	err := ErrorTV{
		Causa: ErrorComponente{Nombre: "Condensador"},
	}
	var errTV ErrorTV
	if errors.As(err, &errTV) {
		fmt.Println("encontrado en cadena de error:", errTV)
	}
	var errComp ErrorComponente
	if errors.As(err, &errComp) {
		fmt.Println("encontrado en cadena de error:", errComp)
	}
	err = ErrorTV{
		Causa: errors.New("la TV explotó"),
	}
	if errors.As(err, &errComp) {
		fmt.Println("Esto nunca debería mostrarse")
	}
}
