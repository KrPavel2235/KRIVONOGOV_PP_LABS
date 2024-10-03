package main

import (
	"fmt"
	"time"
)

func main() {
	//Текущая дата
	fmt.Println("Текущая дата:")
	fmt.Println(time.Time.Date(time.Now()))

	//Текущее время
	fmt.Println("Текущее время:")
	fmt.Println(time.Time.Clock(time.Now()))

	//Дата и время
	fmt.Println("Текущие дата и время:")
	fmt.Println(time.Now().Local().Format(time.RFC1123))
}
