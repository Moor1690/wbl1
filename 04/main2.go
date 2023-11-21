package main

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
*/
import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// worker - функция воркера, которая читает данные из канала и выводит их.
func worker(ctx context.Context, id int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): // Если контекст был отменен, завершаем работу воркера.
			fmt.Printf("Worker %d stopping\n", id)
			return
		case val := <-data: // Читаем данные из канала и выводим их.
			fmt.Printf("Worker %d received data: %d\n", id, val)
		}
	}
}

func main() {
	// Создаем контекст с возможностью отмены.
	ctx, cancel := context.WithCancel(context.Background())

	// WaitGroup для ожидания закрытия всех воркеров в main.
	var wg sync.WaitGroup

	// Настройка обработчика сигнала для корректного завершения программы.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Создаем канал для передачи данных.
	data := make(chan int)

	// Количество воркеров.
	var workerCount int // Можно изменить на нужное количество воркеров

	fmt.Print("input workerCount:\t")
	fmt.Scan(&workerCount)

	// Запускаем воркеров.
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(ctx, i, data, &wg)
	}

	// Запускаем поток, который будет записывать данные в канал.
	go func() {
		counter := 0
		for {
			select {
			case <-ctx.Done(): // Если контекст был отменен, завершаем работу.
				close(data)
				return
			default:
				data <- counter
				counter++
				time.Sleep(time.Second / 4) // Здесь можно настроить частоту отправки данных.
			}
		}
	}()

	// Ожидание сигнала для завершения программы.
	<-signals
	fmt.Println("\nSignal received, stopping program...")
	cancel() // Отменяем контекст, что отправляет сигнал воркерам о необходимости завершить работу.

	// Ожидаем завершения работы воркеров.
	wg.Wait()
}
