package main

import (
	"fmt"
)

func main() {
	//6 номер
	
	fmt.Println("Введите 3 числа (int)")

	var num1 int
	fmt.Scanln(&num1)

	var num2 int
	fmt.Scanln(&num2)

	var num3 int
	fmt.Scanln(&num3)

	fmt.Println(float64(num1+num2+num3) / 3)
}
