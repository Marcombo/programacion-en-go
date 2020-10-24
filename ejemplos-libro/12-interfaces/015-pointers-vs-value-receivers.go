package main

import (
	"fmt"
	"reflect"
)

type Val struct{}
func (a Val) String() string {
	return "Implements String with value receiver"
}

type Ptr struct {}
func (b *Ptr) String() string {
	return "Implements String with pointer receiver"
}

func main() {

	var str1 fmt.Stringer = Val{}
	var str2 fmt.Stringer = &Val{}
	var str3 fmt.Stringer = &Ptr{}
	//var str4 fmt.Stringer = Ptr{}
	fmt.Printf("str1 = %#v\n", reflect.TypeOf(str1).Kind().String())
	str1 = &Val{}
	fmt.Printf("str1 = %#v\n", reflect.TypeOf(str2).Kind().String())
	fmt.Printf("str3 = %#v\n", reflect.TypeOf(str3).Kind().String())
	//fmt.Printf("%#v", str1)
}
