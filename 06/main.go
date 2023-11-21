package main

/*
Реализовать все возможные способы остановки выполнения горутины.
*/
import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Создаем корневой контекст с возможностью отмены.
	ctx, cancel := context.WithCancel(context.Background())

	// WaitGroup для ожидания завершения горутин.
	var wg sync.WaitGroup

	// Горутина 1: Бесконечный цикл, который выполняет "some work1".
	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopping goroutine 1")
				return
			default:
				fmt.Println("some work1")
			}
		}
	}(ctx, &wg)

	// Пауза для выполнения горутины 1 и отмены контекста.
	time.Sleep(time.Microsecond * 20)
	cancel()  // Отмена контекста для завершения горутины 1.
	wg.Wait() // Ожидание завершения горутины 1.

	// Горутина 2: Аналогично горутине 1, но использует контекст с ограничением по времени.
	ctx2, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Microsecond*20))
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopping goroutine 2")
				return
			default:
				fmt.Println("some work2")
			}
		}
	}(ctx2)
	wg.Wait()

	// Горутина 3: Использует канал для остановки.
	stopChannel := make(chan struct{})
	wg.Add(1)
	go func(sChannel chan struct{}) {
		defer wg.Done()
		for {
			select {
			case <-sChannel:
				fmt.Println("stopping goroutine 3")
				return
			default:
				fmt.Println("some work3")
			}
		}
	}(stopChannel)

	time.Sleep(time.Microsecond * 20)
	close(stopChannel) // Закрытие канала для завершения горутины 3.
	wg.Wait()

	// Горутина 4: Использует контекст с таймаутом.
	ctx3, _ := context.WithTimeout(context.Background(), 20*time.Millisecond)
	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopping goroutine 4")
				return
			default:
				fmt.Println("some work4")
			}
		}
	}(ctx3, &wg)

	cancel() // Отмена контекста для завершения горутины 4.
	wg.Wait()
}
