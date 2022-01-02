package main

import "fmt"

type Product struct {
	id string
	title string
	description string
	price float64

}

func (prod *Product) printData(){
	fmt.Printf("ID: %v\n", prod.id)
	fmt.Printf("Title: %v\n", prod.title)
	fmt.Printf("description: %v\n", prod.description)
	fmt.Printf("price: %.2f\n", prod.price)
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

	firstProduct.printData()

	secondProduct.printData()
}