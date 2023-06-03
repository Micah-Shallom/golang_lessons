package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// numbers := []int{2, 3, 4, 5}
	// fmt.Println(numbers)

	// nums := make([]int, 2)
	// fmt.Println("%#v\n", nums)

	// type names []string

	// friends := names{"Dan", "Maria"}
	// fmt.Println(friends)

	// myFriend := friends[0]
	// fmt.Println("My best friend is ", myFriend)

	// for index, value := range numbers {
	// 	fmt.Printf("index %v, value %v\n", index, value)
	// }

	////////////comparing slices/////////////////////////

	var n []int //empty and unintialized
	fmt.Println(n == nil)

	m := []int{} //empty but initialized
	fmt.Println(m == nil)

	//slices cannot be compared to one another with the equal operator only to nil

	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	_, _ = a, b
	// fmt.Println(a == b) //throws an error

	// to compare 2 slices together, we use a for loop to iterate through then compare both element by element
	var eq bool = true

	for i, valueA := range a {
		if valueA != b[i] {
			eq = false
			break
		}
	}
	if len(a) != len(b) {
		eq = false
	}
	if eq {
		fmt.Println("a and b equal")
	} else {
		fmt.Println("a and b not equal")
	}

	// Append to a slice and/or slice copying

	numbers := []int{2, 3}
	numbers = append(numbers, 10)
	fmt.Println(numbers)

	numbers = append(numbers, 20, 30, 40)
	fmt.Println(numbers)

	o := []int{2, 3, 4}
	numbers = append(numbers, o...)
	fmt.Println(numbers)

	src := []int{10, 20, 30}
	dst := make([]int, 2)
	nn := copy(dst, src)
	fmt.Println(src, dst, nn)

	//slice expression
	//notice that when we slice an array, it returns a slice and not an array

	var narr = [5]int{1, 2, 3, 4, 5}

	barr := narr[1:4]

	fmt.Printf("%v, %T\n", barr, barr)

	// slice internals backing arrays and slice headers
	// when a slice is created by slicing an array, the array becomes the backing array of the new slice
	//this shows that slices from an array are always connected and not referred to as a complete new slice
	s1 := []int{10, 20, 30, 40, 50, 60}
	s3, s4 := s1[0:2], s1[1:3]

	s3[1] = 600
	fmt.Println(s1)
	fmt.Println(s4)

	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]

	arr1[1] = 600
	fmt.Println(arr1, slice1, slice2)

	//to create a completely new slice, one that isnt connected to its backing array, we use the append function

	cars := []string{"Ford", "Nissan", "Honda", "RangeRover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:3]...)

	cars[0] = "Chevrolet"

	fmt.Println(cars, newCars)

	// now observe something crazy
	ss1 := []int{10, 20, 30, 40, 50}
	newSlice := ss1[0:3]
	fmt.Println(len(newSlice), cap(newSlice)) // 3 5
 
	newSlice = ss1[2:5]
	fmt.Println(len(newSlice), cap(newSlice)) // 3 3

	fmt.Printf("%p, %p\n", &ss1, &newSlice)

	ar := [5]int{1,2,3,4,5}
	sl := []int{1,2,3,4,5}
	fmt.Printf("array size in byte %d \n", unsafe.Sizeof(ar))
	fmt.Printf("slice size in byte %d \n", unsafe.Sizeof(sl))

}
