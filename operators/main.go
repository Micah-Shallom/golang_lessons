package main

import "fmt"

type km float64
type mile float64

func main(){
	var a = "ops"
	fmt.Println(a)

	var x = 3
	var y = 3.1
	fmt.Println(x * int(y))

	// converting strings to integers/floats and vice versa
	s := string(99)
	fmt.Println(s)

	var sfl = fmt.Sprintf("%f", 3.14)
	fmt.Printf("%T",sfl)


	type age int //int is the underlying type
	type oldage age //int is the underlying type                       
	type veryoldage age //int is the underlying type

	type speed uint

	var s1 speed = 10
	var s2 speed = 20

	_,_ = s1,s2

	var z = 30

	// z = uint(z)

	fmt.Println(z)

	var s3 = speed(z)
	fmt.Println(s3 )

	var parisToLondon km = 465
	var distanceInMiles mile

	distanceInMiles = mile(parisToLondon) / 0.621

	fmt.Print(distanceInMiles)

	type second = uint

	var hour second = 3600
	fmt.Printf("\nminutes in an hour: %d \n", hour/60)
}
