package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a string = "initial"
	fmt.Println(a)

	var b, c = 1, 2
	fmt.Println(b, c)
	// reflect.TypeOf allows to print the type of any variable
	// int by default
	fmt.Println(reflect.TypeOf(b))

	var d = true
	fmt.Println(d)

	var e int = 9
	fmt.Println(e)

	f := "short"
	fmt.Println(f)
}