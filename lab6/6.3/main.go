package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	numberChannel := make(chan int)

	wg.Add(2)

	// Горутина для генерации случайного числа
	go func() {
		defer wg.Done()
		for {
			num := rand.Intn(100) // Генерируем случайное число от 0 до 99
			numberChannel <- num
			time.Sleep(1 * time.Second) // Задержка для генерации нового числа
		}
	}()

	// Горутина для проверки четности числа
	go func() {
		defer wg.Done()
		for num := range numberChannel {
			if num%2 == 0 {
				fmt.Printf("Число %d четное\n", num)
			} else {
				fmt.Printf("Число %d нечетное\n", num)
			}
		}
	}()

	// Запускаем горутины на 10 секунд, затем завершаем
	time.Sleep(10 * time.Second)
	close(numberChannel) // Закрываем канал, чтобы завершить вторую горутину
	wg.Wait()            // Ждем завершения всех горутин
}
