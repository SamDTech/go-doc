package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	id string
	title string
	description string
	price float64

}

func (prod *Product) store(){
	file, _ := os.Create(prod.id + ".txt")

	content := fmt.Sprintf("ID: %v\nTitle: %v\ndescription: %v\nprice: %.2f\n", prod.id, prod.title, prod.description, prod.price)

	file.WriteString(content)

	file.Close()
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
	createdProduct := getUserInput()

	createdProduct.printData()

	createdProduct.store()
}

func getUserInput() *Product{
	fmt.Println("Please Enter the product details")

	fmt.Println("---------------------------")

	reader := bufio.NewReader(os.Stdin)

	idInput := readUserInput(reader, "Product ID: ")

	titleInput := readUserInput(reader, "Product title: ")

	descInput := readUserInput(reader, "Product description: ")

	priceInput := readUserInput(reader, "Product price: ")

	priceValue, _ := strconv.ParseFloat(priceInput, 64)

	product := NewProduct(idInput, titleInput, descInput, priceValue)

	return product
}

func readUserInput(reader *bufio.Reader, promptText string) string{
	fmt.Print(promptText)

	userInput, _ := reader.ReadString('\n')

	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput
}