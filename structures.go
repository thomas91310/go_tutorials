package main
import "fmt"

type person struct {
	name string
	age int
}

func change_age(aperson *person) {
	aperson.age = 77
}

func main() { 
	thomas := person{name: "Thomas", age: 24}
	baptiste := person{name: "Baptiste", age: 21}
	estelle := person{name: "Estelle", age: 13}

	fmt.Println("Thomas : ", thomas)
	fmt.Println("Baptiste : ", baptiste)
	fmt.Println("Estelle : ", estelle)

	fmt.Println("Thomas pointer which is not a pointer: ", &thomas)
	fmt.Println("Estelle age : ", estelle.age)
	estelle.age = 100
	fmt.Println("Age Estelle is now :", estelle.age)

	change_age(&estelle)
	fmt.Println("Age Estelle is now :", estelle.age)

	thomaspointer := &thomas
	fmt.Println("thomas pointer : ", thomaspointer)

	fmt.Println("thomas age : ", thomaspointer.age)
	change_age(thomaspointer)
	fmt.Println("thomas age : ", thomaspointer.age)
}