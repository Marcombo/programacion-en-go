package main

import (
	"fmt"

	// despite the package name is short, you need to import the whole path from the GOPATH
	// or the vendor folder (or the modules root)
	"source.datanerd.us/mmacias/go-for-java-programmers/02-basic-go/0200-packages/cube"
)

func main() {
	// to use public functions, constants, variables and types you need to use the name
	// of the package as a prefix, followed by a dot
	fmt.Println("The volume of a cube of edge 3.5 is", cube.Volume(3.5))
}
