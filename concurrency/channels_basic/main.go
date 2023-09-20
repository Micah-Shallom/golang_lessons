package main

import "fmt"

func main() {
	var ch chan int
	fmt.Println(ch) // zero value of a channel is <nil>

	ch = make(chan int) //initializing a channel
	fmt.Println(ch) //returns a pointer value

	c := make(chan int) //the above channel creation and initialization can be achieved with a one liner using this short declaration
	fmt.Println(c)

	// <- channel operator

	//SEND Operation
	c <- 10 //this sends the value 10 into the channel

	//RECEIVE operation
	num := <- c

	fmt.Println(num) //or
	fmt.Println(<- c)

	close(c) //this closes a channel

	
}
