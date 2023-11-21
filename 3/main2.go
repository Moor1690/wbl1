package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now() // Запоминаем время начала выполнения

	var wg sync.WaitGroup
	result := 0
	arr := make([]int, 1e6)

	for i := range arr {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			arr[i] = arr[i] * arr[i]
		}(i)
	}

	wg.Wait()

	for i := range arr {
		result += arr[i]
	}
	fmt.Println(result)
	elapsed := time.Since(start) // Вычисляем время выполнения
	fmt.Printf("Execution time: %d ns\n", elapsed.Nanoseconds())
}
