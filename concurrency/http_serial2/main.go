package main

import (
	"fmt"
	"io"
	"io/ioutil"
	// "log"
	"net/http"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, c chan string) {
	response, error := http.Get(url)

	if error != nil {
		s := fmt.Sprintf("%s - site is down!\n", url)
		s += fmt.Sprintf("Error: %v \n", error)
		c <- s //sending error into channel

	} else {
		defer response.Body.Close()
		s := fmt.Sprintf("%s - status code is - %d \n", url, response.StatusCode)

		if response.StatusCode == 200 {
			bodyBytes, error := io.ReadAll(response.Body)

			_ = error

			file := strings.Split(url, "//")[1]
			file += ".txt"
			s := fmt.Sprintf("Writing response Body to %s\n", file)

			error = ioutil.WriteFile(file, bodyBytes, 0664)
			if error != nil {
				// log.Fatal(error)
				s += "Error writing file \n"
				c <- s
			}
		}
		s += fmt.Sprintf("%s - site is up", url)
		c <- s

	}
}

func main() {

	ch := make(chan string)
	defer close(ch)

	urls := []string{"http://www.google.com", "http://www.medium.com", "http://www.facebook.com"}

	for _, url := range urls {
		go checkAndSaveBody(url, ch)
		fmt.Println(strings.Repeat("#", 20))
	}

	fmt.Println("No of Goroutines:", runtime.NumGoroutine())

	for i := 0; i <= len(urls); i++ {
		fmt.Println(<- ch)
	}

}
