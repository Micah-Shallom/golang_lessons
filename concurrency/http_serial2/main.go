package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"
)

func checkURL(url string, c chan string) {
	response, error := http.Get(url)

	if error != nil {
		s := fmt.Sprintf("%s - site is down!\n", url)
		s += fmt.Sprintf("Error: %v \n", error)
		fmt.Println(s)
		c <- url //sending error into channel
		
		} else {
			defer response.Body.Close()
			s := fmt.Sprintf("%s - status code is - %d \n", url, response.StatusCode)
			s += fmt.Sprintf("%s - site is up", url)
			fmt.Println(s)
			c <- url

	}
}

func main() {

	ch := make(chan string)
	defer close(ch)

	urls := []string{"http://www.google.com", "http://www.medium.com", "http://www.facebook.com"}

	for _, url := range urls {
		go checkURL(url, ch)
	}

	fmt.Println("No of Goroutines:", runtime.NumGoroutine())

	// for i := 0; i <= len(urls); i++ {
	// 	fmt.Println(<- ch)
	// }

	// for {
	// 	go checkURL(<- ch, ch)
	// 	fmt.Println(strings.Repeat("#", 30))
	// 	time.Sleep(time.Second * 2)
	// }
	
	// for url := range ch {
	// 	go checkURL(url, ch)
	// 	fmt.Println(strings.Repeat("#", 30))
	// 	time.Sleep(time.Second * 2)
	// }
	
	for url := range ch {
		go func(u string){
			checkURL(u, ch)
			fmt.Println(strings.Repeat("#", 30))
			time.Sleep(time.Second * 2)
	
		}(url)
	}

}
