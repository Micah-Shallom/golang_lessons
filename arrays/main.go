package main

import "fmt"

func main() {
	// var numbers [4]int

	// fmt.Printf("%d\n", numbers)
	// fmt.Printf("%#v\n", numbers)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [4]int{-10, 1, 100}
	fmt.Printf("%#v\n", a2)

	a3 := [4]string{"Dan", "Diana", "Paul", "John"}
	fmt.Printf("%#v\n", a3)

	a4 := [4]string{"x", "y"}
	fmt.Printf("%#v\n", a4)

	//ellipsis operator

	a5 := [...]int{1, 2, 5, 1, -10, 66}
	fmt.Printf("%#v\n", a5)
	fmt.Printf("The length of a5 is %d\n", len(a5))

	for i, v := range a2 {
		fmt.Println("Index: ", i, "Values: ", v)
	}

	//multi-dimensional arrays

	balances := [2][3]int{
		{5, 6, 7},
		{8, 9, 10}, //this comma is mandatory
	}
	fmt.Println(balances)

	// keyed elements array
	grades := [3]int{
		1: 10,
		0: 3,
		2: 43,
	}
	fmt.Println(grades)

	accounts := [3]int{2: 50}
	fmt.Println(accounts)

	names := [...]string{5: "Shallom"}
	fmt.Println(names)

	cities := [...]string{
		5: "Kaduna",
		"London",
		1: "France",
	}
	fmt.Printf("%#v\n", cities)

	numbers := []int{2, 3, 4, 5}

	fmt.Println(numbers)
}
