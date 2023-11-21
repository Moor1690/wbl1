package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	ch := make(chan int)
	var wg sync.WaitGroup
	chClose := make(chan os.Signal, 1)
	signal.Notify(chClose, syscall.SIGINT, syscall.SIGTERM)
	N := 5

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(ch <-chan int, id int, chClose chan os.Signal, wg *sync.WaitGroup) {
			defer wg.Done()

			for {
				select {
				case <-chClose:
					fmt.Printf("Worker %d stopping\n", id)
					return
				case val, ok := <-ch: // читаем данные из канала и выводим их
					if !ok {
						// канал данных закрыт, завершаем работу воркера
						fmt.Printf("Worker %d stopping because the data channel was closed\n", id)
						return
					}
					// читаем данные из канала и выводим их
					fmt.Printf("Worker %d received data: %d\n", id, val)
				}
			}
		}(ch, i, chClose, &wg)
	}
	go func() {
		i := 0
		for {
			select {
			case <-chClose:
				fmt.Printf("stopping\n")
				return
			default:
				ch <- i
				i++
				time.Sleep(time.Second / 4)
			}
		}
	}()
	<-chClose
	fmt.Println("\nSignal received, stopping program...")
	close(chClose)
	wg.Wait()
}
