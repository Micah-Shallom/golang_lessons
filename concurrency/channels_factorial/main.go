package main

import "fmt"

func factorial(n int, ch chan int) {
	f := 1

	for i := 2; i <= n; i++ {
		f *= i
	}

	ch <- f
}

func main() {
	ch := make(chan int)
	defer close(ch)

	go factorial(5, ch)

	f := <- ch
	fmt.Println(f)

	for i := 1; i <= 20; i++ {
		go factorial(i, ch)
		f = <- ch
		fmt.Println(f)
	}

}
