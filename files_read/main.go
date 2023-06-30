package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	//open file for reading
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := make([]byte, 2)
	noBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("No of Bytes read %d\n", noBytesRead)
	log.Printf("Data read: %s\n", byteSlice)

	//reaing all contents of a file using ioutil
	file, err = os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data as string: %s\n", data)
	log.Printf("Number of bytes read: %d\n", len(data))

	// we can use io.ReadFile which reads a file into a byte slice; it provides a quick way of reading the entire contents of a file into a byte slice

	data, err = ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data read: %s\n", data)

}
