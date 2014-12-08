package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func add(n1 int, n2 int) (int, int, int) {
	return n1, n2, n1 + n2
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() //we don't care about the first value returned
	fmt.Println(c)	

	_, _, e := add(a, b)
	fmt.Println(e)

}