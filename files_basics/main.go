package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	var err error
	newFile, err = os.Create("a.txt")
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err) //another way to print errors, recommended
	}

	err = os.Truncate("a.txt", 0) //used to truncate a file based on specified number of bytes
	if err != nil {
		log.Fatal(err)
	}
	newFile.Close()

	//Opening and Closing files
	// file, err := os.Open("a.txt")
	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")
	p := fmt.Println
	p("File Name: ", fileInfo.Name())
	p("Size In Bytes:", fileInfo.Size())
	p("Last Modified:", fileInfo.ModTime())
	p("Is Directory?", fileInfo.IsDir())
	p("Permissions: ", fileInfo.Mode())

	// fileInfo, err = os.Stat("b.txt")
	// if err != nil {
	// 	if os.IsNotExist(err){
	// 		log.Fatal("File does not exist!!!")
	// 	}
	// }

	//renaming a file or move
	oldPath := "a.txt"
	newPath := "aaa.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}

	//removing a file
	err = os.Remove("aaa.txt")
	if err != nil {
		log.Fatal(err)
	}

	file, err = os.OpenFile(
		"b.txt",
		os.O_WRONLY | os.O_TRUNC | os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("I learn Golang!")
	byteWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes Written: %d\n", byteWritten)

	//using ioutils
	bs := []byte("Go programming is cool!")
	err = ioutil.WriteFile("c.txt", bs , 0644)
	if err != nil {
		log.Fatal(err)
	}

	//using bufio
	file, err = os.OpenFile(
		"my_file.txt",
		os.O_WRONLY | os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)
	bs = []byte{97,98,99} // use () for string and {} for actual bytes
	bytesWritten, err := bufferedWriter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file) %d\n", bytesWritten)

	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes available in buffer: %d\n", bytesAvailable)

	bytesWritten, err = bufferedWriter.WriteString("\nMy name is shallom")
	if err != nil {
		log.Fatal(err)
	}
	_ = bytesWritten

	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes Buffered %d\n", unflushedBufferSize)

	bufferedWriter.Flush()

	// bufferedWriter.Reset(bufferedWriter) //used in a situation where we didnt flush the buffered content into a file, we reset it

}
