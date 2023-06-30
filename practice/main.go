package main

import (
	// "fmt"
	"io/ioutil"
	"log"
	"os"
)

func main(){
	file, err := os.OpenFile(
		"z.txt",
		os.O_APPEND | os.O_CREATE | os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("My name is Shallom")
	byteWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("The number of written bytes are: %d\n", byteWritten)

	bs := []byte("Go programming is cool i swear")
	err = ioutil.WriteFile("y.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}
}