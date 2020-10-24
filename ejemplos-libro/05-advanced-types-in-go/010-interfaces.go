package main

import "fmt"

// As in Java, all the implementors of the Greeter interface implement the
// Greet() function
type Greeter interface {
	Greet() string
}

// Interfaces in Go are implicit: just make a struct to implement methods with
// those signatures
type Dog struct {
	Name string
}

// in the book, we will use pointer receivers for interface implementation,
// since it is the most extended and advised form, despite we can also implement
// them as value receivers
func (d Dog) Greet() string {
	return "Woof!!"
}

// Implementing the fmt.Stringer interface is a standard way of providing a
// Java toString equivalent
func (d *Dog) String() string {
	return "A dog named " + d.Name
}

func main() {
	// as interfaces are implemented through pointer receivers, variables of
	// a given interface type must reference to pointers
	var greeter Greeter = Dog{Name: "Bingo"}
	var str fmt.Stringer = &Dog{Name: "Rufo"}

	fmt.Println(greeter.Greet())
	fmt.Println(str.String())

	f := Foo{Bar: Dog{}}
}

type Foo struct {
	Bar interface{ Greet() string }
}
