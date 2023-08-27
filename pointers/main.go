package main

import "fmt"

func change(a *int) *float64 {
	*a = 100

	b := 5.5
	return &b
}

func changeVar(a int) {
	a = 66
}

func pointer_functions() {
	x := 8
	p := &x

	fmt.Println("Value of x before calling change():", x)
	change(p)
	fmt.Println("Value of x after calling change():", x)

	fmt.Println("Value of x before calling changeVar():", x)
	changeVar(x)
	fmt.Println("Value of x after calling changeVar():", x)

}

func changeValues(quantity int, price float64, title string, sold bool) {
	quantity = 3
	price = 500.4
	title = "Mobile Phone"
	sold = false
}

func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool){
	*quantity = 3
	*price = 500.5
	*name = "Mobile Phone"
	*sold = false
}

type Product struct {
	price float64
	productName string
}

func changeProduct(p Product){
	p.price = 300
	p.productName = "Bicycle"
}
func changeProductByPointer(p *Product){
	(*p).price = 300 //equivalent to p.price = 300
	(*p).productName = "Bicycle"
}


//testing the effect of pointers on maps and slice
func changeSlice(s []int){
	for i := range s{
		s[i]++
	}
}

func changeMap(m map[string]int){
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30

} //we dont use pointers to maps cus a map value is already a pointer

func main() {
	quantity, price, title, sold := 5, 300.4, "Laptop", true
	fmt.Println("BEFORE calling the changevalues():", quantity, price, title, sold)
	changeValues(quantity, price, title, sold)
	fmt.Println("AFTER calling the changevalues():", quantity, price, title, sold)
	changeValuesByPointer(&quantity, &price, &title, &sold)
	fmt.Println("AFTER calling the changevaluesByPointer():", quantity, price, title, sold)
	

	gift := Product{
		price : 100,
		productName: "Watch",
	}
	changeProduct(gift)
	fmt.Println(gift)
	changeProductByPointer(&gift)
	fmt.Println(gift)

	prices := []int{1,2,3}
	changeSlice(prices)
	fmt.Println("Prices slice after calling changeSlice():", prices)//when we check the output, we see that the function increments each value of the slice without the need of passing pointers to it, this is because a slice already contains a pointer to its backing array(main array)

	myMap := map[string]int{"a": 1, "b": 2}
	changeMap(myMap)
	fmt.Println("myMap after calling changeMap():", myMap)


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


//Note: In a nutshell int,float,strings,bools and structs are only modified by functions only if they are passed as pointers