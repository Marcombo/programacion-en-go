package main

import (
	"errors"
	"fmt"
)

var (
	ErrYaExiste      = errors.New("clave ya existe")
	ErrNoCapacidad   = errors.New("no hay capacidad")
	ErrClaveInvalida = errors.New("clave inválida")
)

type Almacen struct {
	capacidad int
	elementos map[string]interface{}
}

func NewAlmacen(capacidad int) Almacen {
	return Almacen{
		capacidad: capacidad,
		elementos: map[string]interface{}{},
	}
}

func (s *Almacen) Guarda(
	clave string, valor interface{}) error {

	// verifica que la clave sea legal
	if clave == "" {
		return ErrClaveInvalida
	}
	// verifica que no se haya llegado al límite de elementos
	if len(s.elementos) >= s.capacidad {
		return ErrNoCapacidad
	}
	// verifica que la clave no exista ya
	if _, ok := s.elementos[clave]; ok {
		return ErrYaExiste
	}
	// todo OK: guardar el valor
	s.elementos[clave] = valor
	return nil
}

func main() {
	s := NewAlmacen(30)
	err := s.Guarda("un_numero", 12345)
	switch err {
	case ErrYaExiste:
		fmt.Println(err, ": prueba con otra clave única")
	case ErrNoCapacidad:
		fmt.Println(err, ": no se pueden guardar más elementos")
	case ErrClaveInvalida:
		fmt.Println(err, ": prueba con otra clave bien formateada")
	case nil:
		fmt.Println("operación llevada a cabo con éxito!")
	default:
		fmt.Println("error desconocido:", err)
	}
}
