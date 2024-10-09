package main

import (
	"fmt"
	"lab3/3.1/mathutils"
)

func main() {
	fmt.Println("Введите число: ")

	var someNum int
	fmt.Scan(&someNum)
	fmt.Println("Факториал введённого числа: ", mathutils.Factorial(someNum))

}