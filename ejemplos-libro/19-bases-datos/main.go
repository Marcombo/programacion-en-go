package main

import (
	// driver
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Alumno struct {
	Nombre string
	Id     string
	Nota   float64
}

type AccesoDatos struct {
	bd           *sql.DB
	inserta      *sql.Stmt
	buscaPorNota *sql.Stmt
}

func NuevoAccesoDatos(bd *sql.DB) (AccesoDatos, error) {
	var err error
	a := AccesoDatos{bd: bd}
	a.inserta, err = bd.Prepare(`
				INSERT INTO Alumnos(id, nombre, nota)
				VALUES (?, ?, ?)`)
	if err != nil {
		return a, err
	}
	a.buscaPorNota, err = bd.Prepare(`
				SELECT id, nombre, nota FROM Alumnos
				WHERE nota >= ? AND nota < ?
				ORDER BY nota
  `)
	return a, err
}

func main() {
	bd, err := sql.Open("sqlite3", "./archivo.db")

	if err != nil {
		panic(err)
	}
	_, err = bd.Exec(`CREATE TABLE IF NOT EXISTS
    Alumnos(
			id     VARCHAR PRIMARY KEY,
			nombre VARCHAR,
			nota   FLOAT
		)`)
	if err != nil {
		panic(err)
	}
	defer bd.Close()

	ad, _ := NuevoAccesoDatos(bd)

	aula := []Alumno{
		{Id: "1234A", Nombre: "Jose", Nota: 3.5},
		{Id: "5678B", Nombre: "Lara", Nota: 10},
		{Id: "9101C", Nombre: "Sara", Nota: 7.5},
		{Id: "1121D", Nombre: "Ivan", Nota: 8},
		{Id: "3141E", Nombre: "Juan", Nota: 5.75},
	}
	if err := ad.InsertaTodos(aula); err != nil {
		fmt.Println("error insertando alumnos:", err)
	}
	notables, err := ad.BuscaPorNota(7, 9)
	if err != nil {
		panic(err)
	}
	fmt.Println("Los alumnos con notable son:")
	for _, a := range notables {
		fmt.Printf("%+v\n", a)
	}
}

func (d *AccesoDatos) InsertaTodos(aula []Alumno) error {
	tx, err := d.bd.Begin()
	if err != nil {
		return err
	}
	insertaTx := tx.Stmt(d.inserta)
	for _, a := range aula {
		_, err := insertaTx.Exec(a.Id, a.Nombre, a.Nota)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (d *AccesoDatos) Inserta(a Alumno) error {
	_, err := d.inserta.Exec(a.Id, a.Nombre, a.Nota)
	return err
}

func (d *AccesoDatos) BuscaPorNota(
	min, max float64) ([]Alumno, error) {
	encontrados, err := d.buscaPorNota.Query(min, max)

	if err != nil {
		return nil, err
	}
	var alumnos []Alumno
	for encontrados.Next() {
		a := Alumno{}
		err := encontrados.Scan(&a.Id, &a.Nombre, &a.Nota)
		if err != nil {
			return nil, err
		}
		alumnos = append(alumnos, a)
	}
	return alumnos, nil
}
