package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Функция-воркер для обработки задач
func worker(id int, tasks <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении работы воркера
	for task := range tasks {
		reversed := reverseString(task)
		fmt.Printf("Рабочий %d: Обработана строка '%s' -> '%s'\n", id, task, reversed)

		results <- reversed
	}
}

func main() {
	// Чтение количества воркеров от пользователя
	fmt.Print("Введите количество воркеров: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	numWorkers, _ := strconv.Atoi(strings.TrimSpace(input))

	// Создание каналов для задач и результатов
	tasks := make(chan string)
	results := make(chan string)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i+1, tasks, results, &wg) // Запуск каждого воркера в отдельной горутине
	}

	// Чтение строк из файла
	file, err := os.Open("C:/Users/Pavel/Downloads/KRIVONOGOV_PP_LABS-main (1)/KRIVONOGOV_PP_LABS-main/lab6/6.6/text.txt") // Замените на ваш путь
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks <- scanner.Text() // Отправка задачи в канал tasks
	}

	// Закрытие канала tasks для сигнализации окончания отправки задач
	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Результат:", result)
	}

	fmt.Println("Все рабочие завершили обработку.")
}
