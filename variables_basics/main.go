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

	// var a,b,c int
	// fmt.Println(a,b,c)

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

	age, year, pi := 23, 1999, 3.142
	fmt.Printf("Your age is %d", year)

	fmt.Printf("The value of Pi is %f\n", pi)

	figure, radius := "circle", 3

	fmt.Printf("The  %s has a radius of %d and an area of %f\n",figure,radius,float64(radius)*pi)
	fmt.Printf("The  %v has a radius of %v and an area of %v",figure,radius,float64(radius)*pi)

	// VERBS:
    // %d -> decimal
    // %f -> float
    // %s -> string
    // %q -> double-quoted string
    // %v -> value (any)
    // %#v -> a Go-syntax representation of the value
    // %T -> value Type
    // %t -> bool (true or false)
    // %p -> pointer (address in base 16, with leading 0x)
    // %c -> char (rune) represented by the corresponding Unicode code point

    a, b, c := 10, 15.5, "Gophers"
    grades := []int{10, 20, 30}

    fmt.Printf("a is %d, b is %f, c is %s \n", a, b, c)    // => a is 10, b is 15.500000, c is Gophers
    fmt.Printf("%q\n", c)                      // => "Gophers"
    fmt.Printf("%v\n", grades)                 // => [10 20 30]
    fmt.Printf("%#v\n", grades)                // => b is of type float64 and grades is of type []int
    fmt.Printf("b is of type %T and grades is of type %T\n", b, grades) 
    // => b is of type float64 and grades is of type []int
    fmt.Printf("The address of a: %p\n", &a)    // => The address of a: 0xc000016128
    fmt.Printf("%c and %c\n", 100, 51011)       // =>  d and 읃  (runes for code points 101 and 51011)

    const pie float64 = 3.14159265359
    fmt.Printf("pi is %.4f\n", pie) // => formatting with 4 decimal points

    // %b -> base 2
    // %x -> base 16
    fmt.Printf("255 in base 2 is %b\n", 255)  //  => 255 in base 2 is 11111111
    fmt.Printf("101 in base 16 is %x\n", 101) // => 101 in base 16 is 65


}

