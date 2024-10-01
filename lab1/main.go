package main

import (
	"fmt"
	"time"
)

func main() {
	//Текущая дата
	fmt.Println("Текущая дата:")
	fmt.Println()

	//Текущее время
	fmt.Println("Текущее время:")
	fmt.Println(time.Time.Clock(time.Now()))

	//Дата и время
	fmt.Println("Текущие дата и время:")
	fmt.Println(time.Now().Local().Format(time.RFC1123))

	//Номера 2 и 3
	intExample, float64Example, stringExample, boolExample := 1, 3.77, "А справа от меня bool ->", false
	fmt.Println(intExample, float64Example, stringExample, boolExample)

	//Номер 4
	fmt.Println("Операции над целочисленными")
	
	fmt.Println(7+12)
	
	fmt.Println(7-12)
	
	fmt.Println(6*8)
	
	fmt.Println(14/30)

	//номер 5
	fmt.Println("Операции над числами с плавоющей запятой")
	
	fmt.Println(7.67 + 6.89)
	
	fmt.Println(7.67 - 6.89)

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
