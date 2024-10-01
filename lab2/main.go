package main

import (
	"fmt"
	"unicode/utf8"
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
	
	//Задание 3
	fmt.Println("Числа от 1 до 10 ")
	var i = 0 
	for i < 10 {
		i++
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Println("")
	
	//Задание 4
	fmt.Println("Введите что-нибудь: ")
	var str1 string
	fmt.Scanln(&str1)
	fmt.Println(utf8.RuneCountInString(str1))
	
	fmt.Println("Введите ширину и высоту: ")
	var num3 int
	var num4 int
	fmt.Scanln(&num3)
	fmt.Scanln(&num4)
	
	rectangle1 := Rectangle{num3, num4}
	rectangle1.GetArea()
	
	//Задание 6
	fmt.Println("Тут результаты Задания №6")
	AVG(6,8,1)
	
	
}

type Rectangle struct {
		width  int
		height int
	}
	
	func (q Rectangle) GetArea() {
		fmt.Print("Площадь равна: ")
		fmt.Println(q.width * q.height)
		fmt.Println(" ")
	}
	
	func AVG(num...int){
		sum := 0
		count := 0
		res := 0.0
		for _, v := range num{
			count++
			sum += v
			
		}
		res = float64(sum/count)
		fmt.Println(res)
	}