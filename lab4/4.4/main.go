package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Введите какое-нибудь слово, желательно с нижнем регистром)))")
	var str string
	fmt.Scan(&str)
	makeUppercase(str)

}



func makeUppercase(str string) {
	fmt.Println(strings.ToUpper(str))
}

