package concise

import (
	"fmt"
	"math"
	"testing"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// func area(r Rectangle) float64 {
// 	return r.width * r.height
// }

// A method is a function with an implicit first argument, called a receiver.
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

// 调用method通过访问，就像struct里面访问字段一样
func TestMethodByRectangle(t *testing.T) {
	// r1 := Rectangle{12, 2}
	// r2 := Rectangle{9, 4}
	// fmt.Println("Area of r1 is: ", area(r1))
	// fmt.Println("Area of r2 is: ", area(r2))

	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println("Area of r1 is", r1.area())
	fmt.Println("Area of r2 is", r2.area())
	fmt.Println("Area of c1 is", c1.area())
	fmt.Println("Area of c2 is", c2.area())
}
