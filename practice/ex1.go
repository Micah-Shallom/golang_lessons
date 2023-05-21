// package main

// import "fmt"

// func main(){
// 	var (
// 		a int
// 		b float64
// 		c bool
// 		d string
// 	)

// 	x , y, z := 20, 15.5, "Gopher"

// 	fmt.Println(a,b,c,d,x,y,z)
// }

package main

import "fmt"

// ERROR -> := is not allowed in package scope (only local scope)
// version := "3.1"

var version = "3.1"

func main() {
	// ERROR -> A string is initialized only using double quotes ("")
	// name = 'Golang'

	name := "Golang"
	fmt.Println(name)
}
