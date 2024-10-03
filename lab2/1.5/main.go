package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите ширину и высоту: ")
	var num3 int
	var num4 int
	fmt.Scanln(&num3)
	fmt.Scanln(&num4)
	
	rectangle1 := Rectangle{num3, num4}
	rectangle1.GetArea()
	
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
	