package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Alumno struct {
	Nombre   string `yaml:"nombre"`
	Apellido string `yaml:"apellido"`
}

type Aula struct {
	Alumnos []Alumno `yaml:"alumnos"`
}

func main() {
	a := Aula{Alumnos: []Alumno{
		{Nombre: "Carolina", Apellido: "Martínez"},
		{Nombre: "Juan Francisco", Apellido: "Pérez"},
	}}
	txt, err := yaml.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(txt))
}
