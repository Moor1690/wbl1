package main2

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// функция воркера, которая читает данные из канала и выводит их
func worker(ctx context.Context, id int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // если контекст был отменен, завершаем работу воркера
			fmt.Printf("Worker %d stopping\n", id)
			return
		case val := <-data: // читаем данные из канала и выводим их
			fmt.Printf("Worker %d received data: %d\n", id, val)
		}
	}
}

func main() {
	// создаем контекст с возможностью отмены
	ctx, cancel := context.WithCancel(context.Background())

	// вэейт группа для ожидания закрытия всех воркеров в main
	var wg sync.WaitGroup

	// настройка обработчика сигнала для корректного завершения программы
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// создаем канал для передачи данных
	data := make(chan int)

	// количество воркеров
	var workerCount int // Можно изменить на нужное количество воркеров

	fmt.Print("input workerCount:\t")
	fmt.Scan(&workerCount)

	// запускаем воркеров
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(ctx, i, data, &wg)
	}

	// запускаем поток, который будет записывать данные в канал
	go func() {
		counter := 0
		for {
			select {
			case <-ctx.Done(): // если контекст был отменен, завершаем работу
				close(data)
				return
			default:
				data <- counter
				counter++
				time.Sleep(time.Second / 2) // здесь можно настроить частоту отправки данных
			}
		}
	}()

	// ожидание сигнала для завершения программы
	<-signals
	fmt.Println("\nSignal received, stopping program...")
	cancel() // отменяем контекст, что отправляет сигнал воркерам о необходимости завершить работу

	// ожидаем завершения работы воркеров
	wg.Wait()
}
