package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//Задание 4
	fmt.Println("Введите что-нибудь: ")
	var str1 string
	fmt.Scanln(&str1)
	fmt.Println(utf8.RuneCountInString(str1))
	
}
