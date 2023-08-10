package main

import (
	"fmt"
	"math"
)

func f1() {
	fmt.Println("This is function 1")
}

func f2(a int, b int) {
	fmt.Println("Sum:", a+b)
}
func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}
func f4(a float64) float64 {
	return math.Pow(a, a)
}
func f5(a int, b int)(int, int){
	return a+b, a*b
}


func sum(a,b int)(s int){
	s = a + b
	return s
}

//variadic function
func f6(a ...int){
	fmt.Printf("%T\n",a)
	fmt.Printf("%#v\n",a)
}

func f7(a ...int){
	a[0] = 50
}

func sumAndProduct(a... float64)(float64, float64){
	sum := 0.
	product := 1.

	for _ , v := range a {
		sum += v
		product *= v
	}
	return sum, product
}

func main() {
	f1()
	f2(5, 7)
	f3(1, 2, 3, 4.5, 3.3, "Love")
	p := f4(2.3)
	fmt.Println(p)
	a,b := f5(1,2)
	fmt.Println(a,b)
	s := sum(1,2)
	fmt.Println(s)
	fmt.Println("---------------------------------")
	f6(1,2,3,4,5)


	//another example of a variadic function is append
	nums := []int{1,2}
	nums = append(nums, 3,4,5,6,7)
	f6(nums...)
	f7(nums...) 
	fmt.Println("Nums: ",nums) //output [50,2,3,4,5,6,7]

	y, z := sumAndProduct(1.,2.,3.,4.,5.)
	fmt.Println(y,z)

}
