package main

/*
Реализовать конкурентную запись данных в map.
*/
import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Исходный мап myMap с некоторыми значениями.
	myMap := map[int]int{
		0: 0,
		1: 2,
		2: 4,
		3: 8,
		4: 16,
	}

	// Итерация по ключам мапа.
	for i, _ := range myMap {
		wg.Add(1)

		// Горутина для обработки каждого элемента мапа.
		go func(i int) {
			defer wg.Done()
			// Блокировка мьютекса для безопасной работы с мапом.
			mu.Lock()
			defer mu.Unlock()
			// Умножение значения элемента на само себя.
			myMap[i] = myMap[i] * myMap[i]
		}(i)
	}

	wg.Wait() // Ожидание завершения всех горутин.

	fmt.Println(myMap) // Вывод обновленного мапа.
}
