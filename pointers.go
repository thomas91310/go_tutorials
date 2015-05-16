package main
import "fmt"

// if you plan on modifying a value in a function, pass a reference to the object. 
// Passing to value itself won't work. you have to change the value of the memory location where the pointer points to.
func zeroval(ival int) {
    ival = 0
    fmt.Println("ival : ", ival)
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
	test := 15
	zeroval(test)
	fmt.Println("test : ", test)
	zeroptr(&test)
	fmt.Println("test : ", test)
	fmt.Println("test pointer : ", &test)
}