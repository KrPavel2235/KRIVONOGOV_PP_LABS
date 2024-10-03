package main

import (
	"fmt"
)

func main() {
	//Задание 1
	fmt.Println("Введите число: ")
	var num1 int
	fmt.Scanln(&num1)
	
	if num1 % 2 == 0 {
		fmt.Println("Число чётное")
	} else {
		fmt.Println("Число не чётное")
	}
	
	
}

