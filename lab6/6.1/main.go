package main

import (
	"fmt"
	"math/rand"
	"time"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func generateRandomNumbers(count int) []int {
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		numbers[i] = rand.Intn(100) // рандом
	}
	return numbers
}

func sumSeries(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	rand.Seed(time.Now().UnixNano())

	go func() { //функция просто обрабатывает события, без упоминания в main
		fmt.Println("Факториал...")
		time.Sleep(2 * time.Second)
		var n = rand.Intn(20)
		result := factorial(n)
		fmt.Print("Факториал ", n, " :", result, "\n")
	}()

	go func() {
		fmt.Println("Создаём случайные числа...")
		time.Sleep(1 * time.Second)
		numbers := generateRandomNumbers(5)
		fmt.Println("Числа:", numbers)
	}()

	go func() {
		fmt.Println("Сумма серии чисел...")
		time.Sleep(3 * time.Second)
		sum := sumSeries(10)
		fmt.Println("сумма:", sum)
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Горутины закрываются...")
}
