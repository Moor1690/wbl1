package main

import (
	"fmt"
	"sync"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	var wg sync.WaitGroup

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, val := range arr {
			chan1 <- val
		}
		close(chan1) // Закрыть канал после отправки всех значений
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range chan1 { // Использовать range для чтения из канала
			chan2 <- val * 2
		}
		close(chan2) // Закрыть второй канал после обработки
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range chan2 { // Использовать range для чтения из канала
			fmt.Println(val)
		}
	}()

	wg.Wait() // Дождаться завершения всех горутин
}
