package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

// Фигуры реализуют интерфейс, если у них будут метод Area
type Shape interface {
	Area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

// Метод Area для Rectangle
func (anyRectanagle Rectangle) Area() float64 {
	return anyRectanagle.height * anyRectanagle.width
}

// Метод Area для Circle
func (anyCircle Circle) Area() float64 {
	return anyCircle.radius * anyCircle.radius * math.Pi
}

func callShapeArea(shapes []Shape) {

	for _, v := range shapes {
		fmt.Println(v.Area())
	}
}

func main() {
	firstRectangle := Rectangle{height: 2, width: 3}

	firstCircle := Circle{radius: 7}

	Shapes := []Shape{firstCircle, firstRectangle}

	callShapeArea(Shapes)
}

