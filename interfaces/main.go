package main

import (
	"fmt"
	"math"
)


type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}


//using an interface
type shape interface {
	area() float64
	perimeter() float64
}

// func printCircle(c Circle) {
// 	fmt.Println("Shape:", c)
// 	fmt.Println("Area: ", c.area())
// 	fmt.Println("Perimeter: ", c.perimeter())
// }

// func printRectangle(r Rectangle) {
// 	fmt.Println("Shape:", r)
// 	fmt.Println("Area: ", r.area())
// 	fmt.Println("Perimeter: ", r.perimeter())
// }

func print(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %#v\n", s.area())
	fmt.Printf("Perimeter: %#v\n", s.perimeter())
}

func main() {
	c1 := Circle{radius: 5.}
	r1 := Rectangle{width: 3., height: 4.3}

	print(c1)
	print(r1)

}
