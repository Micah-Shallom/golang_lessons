package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2)

	go func(c chan int) {
		for i := 0; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c <- i
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}
		close(c)
	}(c)

	fmt.Println("Main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c {
		fmt.Println("Main goroutine received value from channel: ", v)
	}
}
