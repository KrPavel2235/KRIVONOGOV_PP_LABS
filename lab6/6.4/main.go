package main

import (
	"fmt"
	"sync"
)

var (
	counterWithMutex    int
	counterWithoutMutex int
	mutex               sync.Mutex
)

func incrementWithMutex(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mutex.Lock()       // Блокируем мьютекс
		counterWithMutex++ // Увеличиваем счетчик
		mutex.Unlock()     // Разблокируем мьютекс
	}
}

func incrementWithoutMutex(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counterWithoutMutex++ // Увеличиваем счетчик без блокировки
	}
}

func main() {
	var wg sync.WaitGroup

	// Запускаем горутины с мьютексами
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go incrementWithMutex(&wg)
	}

	wg.Wait() // Ждем завершения всех горутин
	fmt.Printf("Итоговое значение счетчика с мьютексами: %d\n", counterWithMutex)

	// Сбрасываем счетчик без мьютексов
	counterWithoutMutex = 0

	// Запускаем горутины без мьютексов
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go incrementWithoutMutex(&wg)
	}

	wg.Wait() // Ждем завершения всех горутин
	fmt.Printf("Итоговое значение счетчика без мьютексов: %d\n", counterWithoutMutex)
}
