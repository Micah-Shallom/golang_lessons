package main

import "fmt"

func stringsExercise(){
	var name = "Shallom"
	country := "Nigeria"
	fmt.Printf("Your name: %s\nCountry: %s\n",name,country)
	fmt.Println(`he says: "Hello"`)
	fmt.Println(`C:\Users\a.txt`)
}

func exercise() {
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	newnums := nums[2 : len(nums)-2]
	sums := 0

	for _, v := range newnums {
		sums += v
	}
	fmt.Println(sums)
	//----------------------------------------------------------
	friends := []string{"Marry", "John", "Paul", "Diana"}
	newfriends := make([]string, len(friends))
	copy(newfriends, friends)
	newfriends[1] = "Shallom"
	fmt.Println(friends, newfriends)
	//----------------------------------------------------------
	people := []string{"Marry", "John", "Paul", "Diana"}
	yourpeople := []string{}
	yourpeople = append(yourpeople, people...)

	yourpeople[3] = "Shallom"
	fmt.Println(people, yourpeople)
	//----------------------------------------------------------
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}
	newYears = append(years[:3], years[len(years)-3:]...)
	fmt.Println(newYears)
}

func main() {
	stringsExercise()
	exercise()
}
