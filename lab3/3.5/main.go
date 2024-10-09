package main

import (
	"fmt"
)

func main() {

	array := [...]int{5, 4, 6, 7, 3}
	fmt.Println(array)

	newSlice := array[:]

	newSlice[2] = 901
	
	newSlice[1] = 89

	newSlice = append(newSlice, 20)
	fmt.Println(array, newSlice)

}
