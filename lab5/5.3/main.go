package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (anyCircle Circle) Area() float64 {
	return anyCircle.radius * anyCircle.radius * math.Pi
}


func main() {

	firstCircle := Circle{radius: 5}
	fmt.Println(firstCircle.Area())

}

