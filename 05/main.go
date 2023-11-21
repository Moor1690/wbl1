package main

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/
import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Создаем контекст с ограничением времени выполнения в 1 секунду.
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*1))

	// Создаем канал для передачи данных.
	chData := make(chan int, 2)

	// WaitGroup для ожидания завершения горутин.
	var wg sync.WaitGroup

	// Горутина для чтения данных из канала.
	wg.Add(1)
	go func(ctx context.Context, chData chan int, wg *sync.WaitGroup) {
		defer wg.Done() // Уменьшает счетчик WaitGroup при завершении горутины
		for {
			select {
			case <-ctx.Done(): // Выполняется, когда контекст завершается (по таймауту)
				fmt.Println("stopping")
				return
			case data := <-chData: // Чтение данных из канала
				fmt.Printf("Data: %d\n", data)
			}
		}
	}(ctx, chData, &wg)

	// Горутина для записи данных в канал.
	wg.Add(1)
	go func(ctx context.Context, chData chan int) {
		defer wg.Done() // Уменьшает счетчик WaitGroup при завершении горутины
		for i := 0; ; i++ {
			select {
			case <-ctx.Done(): // Выполняется, когда контекст завершается (по таймауту)
				fmt.Println("stopping")
				return
			case chData <- i: // Запись данных в канал
			}
		}
	}(ctx, chData)

	// Ожидание завершения всех горутин.
	wg.Wait()

	// Отмена контекста для завершения всех горутин.
	cancel()
}
