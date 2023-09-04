///////////////////////////////////////
//        !!Micah Shallom     !!       //
//        !!23FA       !!       //
//        !!methods  in Go !!        //
//////////////////////////////////////
//type assertions and type switches
//embedding interfaces

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

func (c Circle) volume() float64 {
	return 4 / 3 * math.Pi
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

type Object interface {
	volume ()
}

// embedded interface
//similar to inheritance in python
type Geometry interface{
	Shape
	Object
	getColor() string
}

func main() {
	var s Shape = Circle{radius: 4.5}
	// s.volume() //throws an error as the method isnt a member of the interface
	ball, ok := s.(Circle) //type assertions
	if ok == true {
		fmt.Printf("Ball Volume: %v\n", ball.volume())
	}
}
