package main

import (
	"time"
)

type Estudiante struct {
	nombre     string
	nacimiento time.Time
	descuento  bool
}

type Opcion func(*Estudiante)

func Nombre(nombre string) Opcion {
	return func(stud *Estudiante) {
		stud.nombre = nombre
	}
}

func Nacimiento(nacimiento time.Time) Opcion {
	return func(stud *Estudiante) {
		stud.nacimiento = nacimiento
	}
}

func Descuento() Opcion {
	return func(stud *Estudiante) {
		stud.descuento = true
	}
}

func NewEstudiante(opciones ...Opcion) Estudiante {
	stud := Estudiante{
		nombre: "desconocido",
	}
	for _, opt := range opciones {
		opt(&stud)
	}
	return stud
}

func main() {
	estu1 := NewEstudiante()
	estu2 := NewEstudiante(Nombre("Pedro"), Descuento())
	estu3 := NewEstudiante(Nombre("Juan"), Nacimiento(
		time.Date(2001, 10, 12, 0, 0, 0, 0, time.UTC)))

	// evitamos error de compilación por no usar las
	// variables generadas como ejemplo
	_, _, _ = estu1, estu2, estu3
}
