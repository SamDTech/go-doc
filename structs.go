package main

import "fmt"

type Product struct {
	id string
	title string
	description string
	price float64

}

func NewProduct(id string, t string, d string, p float64) *Product {
		return &Product {id, t, d, p }

}

func main(){
	firstProduct := Product {
		id: "first-product",
		title: "book",
		description: "a new book",
		price: 10.99,
	}

	secondProduct := NewProduct("second product", "pen", "a new pen", 9.99)

	fmt.Println(firstProduct)
	fmt.Println(*secondProduct)
}