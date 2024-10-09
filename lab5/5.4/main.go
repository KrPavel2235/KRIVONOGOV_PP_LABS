package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

func (anyRectanagle Rectangle) Area() float64 { // площадь треугольника 
	return anyRectanagle.height * anyRectanagle.width
}

func (anyCircle Circle) Area() float64 {
	return anyCircle.radius * anyCircle.radius * math.Pi
}

func callShapeArea(shapes []Shape) {

	for _, v := range shapes {
		fmt.Println(v.Area())
	}
}

func main() {

	//Номера 3-4
	firstRectangle := Rectangle{height: 2, width: 3}
	fmt.Print("Площадь треугольника: ",firstRectangle.Area(),"\n")

	firstCircle := Circle{radius: 3}
	fmt.Println("Площадь круга: ",firstCircle.Area(),"\n")

}
