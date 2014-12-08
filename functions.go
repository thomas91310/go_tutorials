package main

import "fmt"

func plus (n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	var res int = plus(4, 6)
	fmt.Println(res)
}