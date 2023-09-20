package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	var gr int = 100
	var n int = 0

	wg.Add(gr * 2)

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			n++
			// fmt.Println("Incrementing f1")
			wg.Done()
		}()
		go func() {
			time.Sleep(time.Second / 10)
			n--
			// fmt.Println("Decrementing f2")
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(n)
}
