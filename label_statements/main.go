package main

import "fmt"

func main() {
	people := [5]string{"Mary", "Mark", "Joy", "Theo", "Shallom"}
	friends := [2]string{"Joy", "Shallom"}

outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %s at index %d\n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("Next instruction after the break")
}
