package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup){

	fmt.Println("f1 goroutine started")
	for i := 0; i < 5; i++ {
		fmt.Println("f1, i=", i)
		time.Sleep(time.Second)
	}

	wg.Done()
}
func f2(){
	fmt.Println("f2 goroutine started")
	for i := 0; i < 3; i++ {
		fmt.Println("f2, i=", i)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	// NumCPU returns the number of logical CPUs or cores usable by the current process.
	fmt.Println("No. of CPUs:", runtime.NumCPU()) // => No. of CPUs: 4

	// NumGoroutine returns the number of goroutines that currently exists.
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine()) // => No. of Goroutines: 1

	// GOOS and GOARCH are environment variables
	fmt.Println("OS:", runtime.GOOS)     // => OS: linux
	fmt.Println("Arch:", runtime.GOARCH) // => Arch: amd64

	//  GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns
	//  the previous setting.
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0)) // => GOMAXPROCS: 4
	// If n < 1, it does not change the current setting.

	go f1(&wg)
	fmt.Println("Number of goroutines after f1", runtime.NumGoroutine())

	f2()
	time.Sleep(time.Second * 2)

	wg.Wait()
	fmt.Println("Main execution stopped")
}
