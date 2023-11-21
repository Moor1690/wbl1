package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now() // Запоминаем время начала выполнения

	var wg sync.WaitGroup
	var mutex sync.Mutex // Мьютекс для защиты операции сложения
	result := 0
	arr := make([]int, 1e6)

	for i := range arr {
		wg.Add(1)
		// Запуск горутины для возведения в квадрат числа
		go func(num int) {
			defer wg.Done()
			square := num * num // Возводим число в квадрат

			mutex.Lock()     // Захват мьютекса перед изменением общей переменной
			result += square // Добавляем квадрат числа к результату
			mutex.Unlock()   // Освобождаем мьютекс
		}(arr[i])
	}

	wg.Wait() // Ожидаем завершения всех горутин

	// Выводим результат
	fmt.Println(result)
	elapsed := time.Since(start) // Вычисляем время выполнения
	fmt.Printf("Execution time: %d ns\n", elapsed.Nanoseconds())
}
