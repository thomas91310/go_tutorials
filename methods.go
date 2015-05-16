package main
import "fmt"

	
type rect struct {
    width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perimeter() int {
	return 2 * (r.width + r.height)
}

func main() { 
	r := rect{width: 10, height: 5}

	pointer_rect := &r

	fmt.Println(pointer_rect.area())
	fmt.Println(pointer_rect.perimeter())
}