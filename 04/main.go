package main

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
*/
import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ch := make(chan int)                                    // Создаем канал для передачи данных между горутинами
	var wg sync.WaitGroup                                   // Группа ожидания для синхронизации завершения горутин
	chClose := make(chan os.Signal)                         // Канал для обработки сигналов остановки
	signal.Notify(chClose, syscall.SIGINT, syscall.SIGTERM) // Регистрация сигналов для остановки
	workerCount := 5                                        // Количество воркеров

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(ch <-chan int, id int, chClose chan os.Signal, wg *sync.WaitGroup) {
			defer wg.Done() // Уменьшение счетчика при завершении горутины

			for {
				select {
				case <-chClose: // Прием сигнала остановки
					fmt.Printf("Worker %d stopping\n", id)
					return
				case val, ok := <-ch: // Обработка данных из канала
					if !ok { // Если канал закрыт, завершаем работу
						fmt.Printf("Worker %d stopping because the data channel was closed\n", id)
						return
					}
					fmt.Printf("Worker %d received data: %d\n", id, val)
				}
			}
		}(ch, i, chClose, &wg)
	}
	go func() { // Горутина для отправки данных в канал
		i := 0
		for {
			select {
			case <-chClose: // Остановка при получении сигнала
				fmt.Printf("stopping\n")
				return
			default:
				ch <- i // Отправка данных в канал
				i++
				time.Sleep(time.Second / 4) // Интервал между отправками
			}
		}
	}()
	<-chClose // Ожидание сигнала остановки
	fmt.Println("\nSignal received, stopping program...")
	close(chClose) // Закрытие канала сигналов
	wg.Wait()      // Ожидание завершения всех горутин
}
