//interfaces dynamice type and polymorphism

package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}
type Rectangle struct {
	width  float64
	height float64
}

type Shape interface {
	area() float64
	perimeter() float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rectangle) area() float64 {
	return 2 * (r.width + r.height)
}
func (r Rectangle) perimeter() float64 {
	return r.width * r.height
}

func print(s Shape) {
	fmt.Printf("Shape is: %v\n", s)
	fmt.Printf("Area is %f\n", s.area())
	fmt.Printf("Perimeter is %f\n", s.perimeter())
}

func main() {
	var s Shape
	fmt.Printf("%T\n", s)

	ball := Circle{radius: 2.5}
	s = ball
	fmt.Println(s.area())
	fmt.Printf("Type is: %T\n", ball)
	
	room := Rectangle{width: 25, height: 21}
	s = room
	fmt.Println(s.area())
	fmt.Printf("Type is: %T\n", room)
	print(s)


	

	 
}
