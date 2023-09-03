package main

import "fmt"

type car struct{
	brand string
	price int
}

type names []string

type Rectangle struct {
	width  float64
	height float64
}

func (n names) print(){
	for i, name := range n{
		fmt.Println(i, name)
	}
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func changeCar(c car, newBrand string, newPrice int){
	c.price = newPrice
	c.brand = newBrand
}

func (c car) changeCar1(newBrand string, newPrice int){
	c.brand = newBrand
	c.price = newPrice
}

//using pointer receiver
func (c *car) changeCar2(newBrand string, newPrice int){
	(*c).brand = newBrand
	(*c).price = newPrice
}


func main() {
	myCar := car{brand: "Audi", price: 40000}
	changeCar(myCar, "Renault", 20000)
	fmt.Println(myCar)

	myCar.changeCar1("Opel", 21000)
	fmt.Println(myCar)

	(&myCar).changeCar2("Seat", 25000)
	fmt.Println(myCar)

	var yourCar *car
	yourCar = &myCar
	fmt.Println(*yourCar)
	yourCar.changeCar2("VM", 17000)
	fmt.Println(*yourCar)

	rect := Rectangle{1.2, 2.2}
	area := rect.Area()
	fmt.Println(area)
	names := names{"Shallom", "Micah"}
	names.print()
}
