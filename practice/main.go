package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64{
	return math.Pi * math.Pow(c.radius, 2)
}

func main(){
	c1 := Circle{radius: 1.2}
	area := c1.area()
	fmt.Println(area)
}