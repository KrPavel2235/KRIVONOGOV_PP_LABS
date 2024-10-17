package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter int // Общая переменная-счетчик

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Задействуем все ядра процессора

	var mutex sync.Mutex // Создаем мьютекс

	// Горутина, увеличивающая счетчик с использованием мьютекса
	go func() {
		for i := 0; i < 1000000; i++ {
			mutex.Lock() // Блокируем мьютекс перед изменением счетчика
			counter++
			mutex.Unlock() // Разблокируем мьютекс после изменения счетчика
		}
	}()

	// Горутина, увеличивающая счетчик без использования мьютекса
	go func() {
		for i := 0; i < 1000000; i++ {
			counter++ // Изменяем счетчик без блокировки мьютекса
		}
	}()

	// Ждем завершения горутин
	time.Sleep(time.Second)

	// Выводим результаты
	fmt.Println("Счетчик с мьютексом:", counter)

	// Сбрасываем счетчик
	counter = 0

	// Горутина, увеличивающая счетчик без использования мьютекса
	go func() {
		for i := 0; i < 1000000; i++ {
			counter++ // Изменяем счетчик без блокировки мьютекса
		}
	}()

	// Горутина, увеличивающая счетчик без использования мьютекса
	go func() {
		for i := 0; i < 1000000; i++ {
			counter++ // Изменяем счетчик без блокировки мьютекса
		}
	}()

	// Ждем завершения горутин
	time.Sleep(time.Second)

	// Выводим результаты
	fmt.Println("Счетчик без мьютекса:", counter)
}
