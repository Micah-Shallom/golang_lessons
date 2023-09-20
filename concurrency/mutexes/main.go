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

	var m sync.Mutex

	wg.Add(gr * 2)

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			n++
			m.Unlock()
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			defer m.Unlock()
			n--
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(n)
}




