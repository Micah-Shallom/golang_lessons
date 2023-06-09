package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := `I am \n a boy!` //raw string
	fmt.Println(s1)

	fmt.Println("Price:100\nBrand:Nike")
	fmt.Println(`
Price:100
Brand:Nike
	`)
	fmt.Println(`C:\Users\Shallom`)
	fmt.Println("C:\\Users\\Shallom")

	var s3 = "I Love " + "Go " + "Programming "
	fmt.Println(s3 + "!")
	fmt.Println("Element at index 0: ", s3[0]) //returns the number 73 as strings are slices of bytes in go

	var1, var2 := 'a', 'b'
	fmt.Printf("Type: %T, Value: %d\n", var1, var2)

	str := "tara"
	fmt.Println(len(str))
	fmt.Println("Byte (not rune) at position 1:", str[1])
	fmt.Println("\n" + strings.Repeat("#", 20))

	s4 := "Golang"
	fmt.Println(len(s4))

	name := "Codruta"
	fmt.Println(len(name))

	p := fmt.Println
	result := strings.Contains("I love Go Programming!", "love")
	p(result)
	
	result = strings.ContainsAny("success", "xys")
	p(result)

	result = strings.ContainsRune("golang", 'g')
	p(result)

	n := strings.Count("cheese","e")
	p(n)

	p(strings.ToLower("GO PyTHON jaVA"))
	p(strings.ToUpper("GO PyTHON jaVA"))
	p("go" == "go")
	p(strings.EqualFold("GO", "go")) //optimized way of comparing strings leaving out the case sensitivity

	mystr := strings.Repeat("ab", 10)
	p(mystr)

	mystr = strings.Replace("192.168.0.1",".",":",2) //replaces first 2 occurences
	mystr = strings.Replace("192.168.0.1",".",":",-1) //replaces all
	mystr = strings.ReplaceAll("192.168.0.1",".",":") //replaces all
	p(mystr)
	s := strings.Split("a,b,c", ",")
	fmt.Printf("%T\n", s)
	fmt.Printf("%#v\n", s)

	s = []string{"I", "learn", "Golang"}
	mystr = strings.Join(s, "-")
	p(mystr)

	s5 := strings.TrimSpace("\t   Goodbye Windows, Welcome Linux!")
	fmt.Printf("%q\n", s5)

	s6 := strings.Trim("....Hello Gophers!!!!!?", ".!?")
	p(s6)
}
