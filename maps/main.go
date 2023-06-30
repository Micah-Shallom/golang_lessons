package main

import "fmt"

func main() {
	var employees map[string]string

	fmt.Printf("%#v\n", employees)

	fmt.Printf("No of pairs: %d", len(employees))

	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["0x323"])

	// var m1 map[[]int]string //throws a compile time error as slices cant be used as keys

	var m1 map[[5]int]string //this is viable as a comparable datatype hence arrays can be used as keys

	_ = m1

	// employees["Shallom"] = "Programmer" //this throws an error as it is illegal to populate an uninitialized map, we have to initialize the map using the make function or a map literal before we can add any elements

	people := map[string]float64{} //initialized map (not empty)

	people["John"] = 12.35
	people["Mary"] = 23.45

	fmt.Println(people)

	map1 := make(map[int]int) //initializing map (empty)
	map1[4] = 8

	balances := map[string]float64{
		"USD": 323.11,
		"EUR": 4356.23, //this final comma is compulsory
	}
	fmt.Println(balances)

	m := map[int]int{1:10, 2:20, 3:30} //no final comma needed for one liner maps
	_ = m

	balances["USD"] = 500.0
	balances["GBP"] = 123.76

	fmt.Println(balances["RON"]) //returns 0 as it is not a key in the map

	v, ok := balances["RON"] //2 variables , 1 returns the value while the other returns bool typee output whether true or false

	if ok{
		fmt.Println("The RON balance is: ", v)
	}else{
		fmt.Println("The RON key does not exist in the map")
	}

	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}

	delete(balances, "USD") //to remove a key and its value
	fmt.Println(balances)


	//COMPARING MAPS
	a := map[string]string{"A": "X"}
	b := map[string]string{"B": "Y"}



	// fmt.Println(a == b) //invalid operation: a == b (map can only be compared to nil)

	 //two maps can only be compared only if they are stringified into a string 
	 s1 := fmt.Sprintf("%s", a)
	 s2 := fmt.Sprintf("%s", b)
	 fmt.Println(s1, s2)

	 fmt.Println(s1 == s2)


	friends := map[string]int{"Dan": 40, "Maria":25}
	neighbors := friends

	friends["Dan"] = 600
	fmt.Println(neighbors)


	// to make a copy of a map without referencing
	peoples := make(map[string]int)

	for k , v := range friends {
		peoples[k] = v
	}
	fmt.Println(peoples)
}