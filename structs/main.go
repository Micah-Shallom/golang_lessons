package main

import "fmt"

func main() {

	//Embedding structs as fields in another struct

	type Contact struct {
		email, address string
		phone int
	}

	type Workers struct{
		name string
		salary int
		contactInfo Contact
	}

	john := Workers{
		name: "shallom micah bawa",
		salary: 4000,
		contactInfo: Contact{
			email: "micahshallom@gmail.com",
			address: "kaduna Nigeria",
			phone: 340008000 ,
		},
	}

	fmt.Printf("%#v\n", john)

	// #############################################
	// #############################################

	// Anonymous Structs
	shallom := struct{
		firstName, lastName string
		age int
	}{
		firstName: "Shallom",
		lastName: "Micah Bawa",
		age: 23,
	}
	fmt.Printf("%#v\n",shallom)

	//Anonymous Fields
	type Book struct{
		string
		float64
		bool
	}

	b1 := Book{"1984 by George", 10.5, false}
	fmt.Printf("%#v\n",b1)
	fmt.Printf("%#v\n",b1.string)

	//mixing named fields and anonymous fields
	type Employee struct{
		name string
		salary int
		bool
	}

	e := Employee{"Shallom", 40000 , false}
	fmt.Printf("%#v\n", e)


	// #############################################
	// #############################################

	type book struct {
		title  string
		author string
		year   int
	}

	lastBook := book{title: "Anna Karenina"}
	fmt.Println(lastBook.title)
	fmt.Printf("%#v\n", lastBook)
	lastBook.author = "Leo Tolstoy"
	lastBook.year = 1878
	fmt.Printf("%+v\n", lastBook)

	//comparing two structs together
	aBook := book{title: "Anna Karenina", author: "", year: 0}
	fmt.Println(lastBook == aBook) //in a struct, we can compare two structs together using the double equals sign operator unlike maps were we can only compare a map to nil or by looping through each element
	
	//making a copy of a struct
	myBook := aBook
	myBook.year = 2020
	fmt.Println(myBook,aBook)


	// #############################################
	// #############################################

	// type book struct {
	// 	title, author string
	// 	year          int
	// }

	// myBook := book{"The Devine Comedy", "Dante Aligheri", 1320}
	// fmt.Println(myBook)

	// bestBook := book{title: "Animal Farm", year: 1945, author: "George Orwell"}
	// _ = bestBook

	// aBook := book{title: "Just a random book"}
	// fmt.Printf("%#v\n", aBook)
	// fmt.Println(aBook)
}
