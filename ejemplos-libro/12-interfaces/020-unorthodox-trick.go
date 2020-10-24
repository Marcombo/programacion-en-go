package main

import "fmt"

// Sometimes it's difficult to realize in compile time that a struct is not
// implementing an interface, as it is referenced from another file in the code
type Jumper interface {
	// jump returns the distance, in meters, after jumping
	Jump() float32
}

type Frog struct{}

func (f *Frog) Jump() float32 {
	fmt.Println("Croak!")
	return 0.75
}

type Kangaroo struct{}

func (f *Kangaroo) Jump() int {
	fmt.Println("Boing!")
	return 2
}

// By mistake (e.g. after a refactoring) the Kangaroo struct is not implementing
// the Jumper interface. We probably won't notice that until we try to compile
// the whole code, and the error won't appear in the definition of the struct
// (because Go doesn't provide the "implements" keyword) but in other part of
// the code, where we try to assign a Kangaroo to a Jumper variable or argument.
// Also, if the interface is defined in other part of the code, it's difficult
// to know all the interfaces a type is implementing.

// Writing this below the type definition will allow us:
// - Know the interfaces a type is implementing
// - Immediately get an error if the type does not implement an interface
var _ Jumper = (*Frog)(nil)
var _ Jumper = (*Kangaroo)(nil)

// WARNING!! Despite this formula is used in a lot of top-visibility projects,
// purist Go community doesn't like this practice

func main() {
	fmt.Println(Kangaroo{}.Jump())
}

func SomeFunctionInADistantPackage() {
	jumpers := []Jumper{&Kangaroo{}, &Frog{}}
	fmt.Println("Jumping all together!")
	for _, j := range jumpers {
		j.Jump()
	}
}
