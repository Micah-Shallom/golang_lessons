package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func checkAndSaveBody(url string) {
	response, error := http.Get(url)

	if error != nil {
		fmt.Printf("%s - site is down", error)
	} else {
		defer response.Body.Close()
		fmt.Printf("%s - status code is - %d \n", url, response.StatusCode)

		if response.StatusCode == 200 {
			bodyBytes, error := io.ReadAll(response.Body)

			_ = error

			file := strings.Split(url, "//")[1]
			file += ".txt"
			fmt.Printf("Writing response Body to %s\n", file)

			error = ioutil.WriteFile(file, bodyBytes, 0664)
			if error != nil {
				log.Fatal(error)
			}

		}
	}
}

func main() {

	urls := []string{"http://www.google.com", "http://www.medium.com", "http://www.facebook.com"}

	for _, url := range urls {
		checkAndSaveBody(url)
		fmt.Println(strings.Repeat("#", 20))
	}
}
