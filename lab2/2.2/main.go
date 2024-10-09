package main

import (
	"fmt"
)

func main() {
	//Задание 2
	fmt.Println("Введите число: ")
	var num2 int
	fmt.Scanln(&num2)
	
	if num2 > 0 {
		fmt.Println("Positive")
	} else if num2 < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}
	
}