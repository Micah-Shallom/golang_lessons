package main

import "fmt"

type Book struct {
	price float64
	title string
}

func (b Book) vat() float64 {
	return b.price * 0.09
}

func (b *Book) discount() float64 {
	(*b).price = b.price * 0.9
	return b.price
}

func (b *Book) changePrice(c float64) {
	b.price = c
}
func main(){
	pH := Book{1.2, "Purple Hibiscus"}
	pH.changePrice(100)
	fmt.Println(pH.vat())
	fmt.Println(pH.discount())
}