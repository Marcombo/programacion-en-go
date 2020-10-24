package main

import "fmt"

/*
func CambiaRegistro(clave, valor string) error {
	cnx := estableceConexion()
	datos, err := cnx.ObtenDatos()
	if err != nil {
		cnx.Cerrar()
		return fmt.Errorf("obteniendo datos: %w", err)
	}
	if err := validar(datos); err != nil {
		cnx.Cerrar()
		return fmt.Errorf("validando datos: %w", err)
	}
	if err := cnx.GuardaDatos(clave, valor); err != nil {
		cnx.Cerrar()
		return fmt.Errorf("guardando datos: %w", err)
	}
	cnx.Cerrar()
	return nil
}*/
func main() {
	fmt.Println("hola")
	defer fmt.Println("ejecución aplazada")
	fmt.Println("adiós!")
}
