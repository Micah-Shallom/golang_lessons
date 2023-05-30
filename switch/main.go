package main

import (
	"fmt"
)

func main() {
	language := "golang"

	switch language {
	case "Python":
		fmt.Println("Learning Python! You dont use curly braces but identation")
	case "GO", "golang":
		fmt.Println("Good, go for Go! YOu are using curly braces {}!")
	default:
		fmt.Println("Any programming language is okay")
	}

	n := 5
	switch true{
	case n%2 == 0:
		fmt.Println("Even number")
	case n%2 != 0 :
		fmt.Println("Odd number")
	default:
		fmt.Println("Never Here")
	}

}
