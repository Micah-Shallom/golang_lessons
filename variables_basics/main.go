package main

import "fmt"

func main() {
	// var age int = 30
	var age = 30
	fmt.Println("Age: ",age)

	var name = "Dan"
	fmt.Println("Your name is: ", name)

	s := "Learning golang!"
	fmt.Println(s)

	name = "Shallom"
	_  =  name

	//multiple declarations
	car, cost := "Audi" , 20000
	fmt.Println(car,cost)

	var opened = false
	opened, file := true, "a.txt"
	_,_ = opened, file


	var (
		salary float64
		firstName string
		gender bool
	)
	fmt.Println(salary,gender,firstName)

	var a,b,c int
	fmt.Println(a,b,c)

	var i,j int
	i, j = 5,8

	j, i = i, j //swapping variables

	fmt.Println(i,j)

	sum := 5 + 2.3
	fmt.Println(sum)

	//commenting 
	/*
	age := 18
	_ = age
	*/
}

