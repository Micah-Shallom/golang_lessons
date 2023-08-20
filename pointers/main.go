package main

import "fmt"

func pointer_functions() {
	x := 8
	p := &x
}

func main() {
	name := "Shallom"
	fmt.Println(&name)

	var x int = 2
	ptr := &x //creates pointer
	fmt.Printf("ptr is of type %T with a value of %v and the address %p\n", ptr, ptr, &ptr)
	fmt.Printf("address of x is %p\n", &x)

	var ptr1 *float64 //declaring a pointer without initializing it with its zero value equal to nil
	_ = ptr1

	p := new(int) //another way of declaring a pointer
	x = 100
	p = &x
	fmt.Printf("p is of type %T with a value of %v\n", p, p)

	*p = 90 //equivalent to x = 90
	fmt.Println(x, *p)

	//Pointer -> Pointer Comparison
	a := 10
	p1 := &a
	pp1 := &p1
	fmt.Printf("Value of p1: %v, Address of p1: %p\n", p1, &p1)
	fmt.Printf("Value of pp1: %v, Address of pp1: %p\n", pp1, &pp1)

	fmt.Printf("*p1 is %v\n", *p1)
	fmt.Printf("*pp1 is %v\n", *pp1)

	fmt.Printf("**pp1 is %v\n", **pp1)

	**pp1++
	fmt.Printf("a is %v\n", a)

	//Comparing Pointers - two pointers are equal if they point to the same value or nil

	var p2 *int
	fmt.Printf("%#v\n", p2) //equals nil

	y := 5
	p2 = &y

	z := 5
	p3 := &z

	fmt.Println(p2 == p3) //false

	p4 := &y
	fmt.Println(p2 == p4) //true

	pointer_functions()

}
