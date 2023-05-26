package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	price, instock := 100, true

	if price > 800 {
		fmt.Println("Too Expensive")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("It is cheap\n")
	}

	if price <= 100 && instock {
		fmt.Println(("buy it"))
	}

	// fmt.Println("os.Args", os.Args)
	// var result, err = strconv.ParseFloat(os.Args[1], 64)
	// fmt.Printf("%T\n", os.Args[1])
	// fmt.Printf("%T\n", result)
	// _ = err

	i, err := strconv.Atoi("45")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	if i, err := strconv.Atoi("50"); err == nil {
		fmt.Println(i)
	}else{
		fmt.Println(err)
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("one Argument is required!")
	}else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("The argument must be an integer! Error:", err)
	}else{
		fmt.Printf("%d km in miles is %v\n",km, float64(km)*0.621)
	}
}
