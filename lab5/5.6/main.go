package main

import (
	"fmt"
)

type Book struct {
	id    int
	title string
	price int
}

type Stringer interface {
	getInfoBook()
}

func (anyBook Book) getInfoBook() {
	fmt.Println("Id: ", anyBook.id)
	fmt.Println("Название: ", anyBook.title)
	fmt.Println("Цена: ", anyBook.price)
}

func main() {

	var firstBook Stringer = &Book{
		id:    1,
		title: "Как Red Hot Chili Peppers превратили баг в фичу",
		price: 540}
	firstBook.getInfoBook()
}


