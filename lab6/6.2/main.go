package main

import (
	"fmt"
	"sync"
)

func fibonacci(c chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении функции
	x, y := 0, 1
	for i := 0; i < 15; i++ {
		c <- x // Отправка числа в канал
		x, y = y, x+y
	}
	close(c) // Закрытие канала после отправки всех чисел
}

func printNumbers(c chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении функции
	fmt.Println("числа фибоначчи:")
	for number := range c {
		fmt.Print(" ", number)
	}
}

func main() {
	c := make(chan int)   // Создание канала для передачи целых чисел
	var wg sync.WaitGroup // Создание WaitGroup

	wg.Add(2) // Увеличиваем счетчик на 2 для двух горутин

	// Запуск горутины для генерации чисел Фибоначчи
	go fibonacci(c, &wg)

	// Запуск горутины для вывода чисел
	go printNumbers(c, &wg)

	wg.Wait() // Ожидание завершения всех горутин
}
