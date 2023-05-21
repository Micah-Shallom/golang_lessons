package main

import "fmt"

// func main() {
// 	const (
// 		daysWeek   = 7
// 		lightSpeed = 299792458
// 		pi         = 3.14159
// 	)

// 	const secPerDay = int(60 * 60 * 24)
// 	const daysYear = 365
// 	fmt.Println(secPerDay * daysYear)
// }


// func main() {
//     const x int = 10

//     // declaring a constant of type slice int ([]int)
//     m := []int{1: 3, 4: 5, 6: 8}
//     fmt.Println(m)
// }

// import "math"

// func main() {
// 	const a = 7
// 	const b float64 = 5.6
// 	const c = a * b

// 	const x = 8
// 	const xc int = x

// 	const noIPv6 = math.Pow(2, 128)
// }


// func main(){
// 	// x, y, z := 10, 15.5, "Gophers"
// 	// score := []int{10, 20, 30}
// 	const x float64 = 1.422349587101

// 	fmt.Printf("%.4f",x)
	
// }

// func main() {
// 	shape := "circle"
// 	radius := 3.2
// 	const pi float64 = 3.14159
// 	circumference := 2 * pi * radius

// 	fmt.Printf("Shape: %s\n",shape)
// 	fmt.Printf("Circle's circumference with radius %f is %f\n", radius, circumference)
// 	_ = shape
// }

func main() {
    x, y := 4, 5.1

    z := x * int(y)
    fmt.Println(z)

    a := 5
    b := 6.2 * float64(a)
    fmt.Println(b)

}