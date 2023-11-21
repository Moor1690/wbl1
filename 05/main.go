package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*1))
	chData := make(chan int, 2)
	var wg sync.WaitGroup

	wg.Add(1)
	go func(ctx context.Context, chData chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopping")
				return
			case data := <-chData:
				fmt.Printf("Data: %d\n", data)
			}
		}

	}(ctx, chData, &wg)

	wg.Add(1)
	go func(ctx context.Context, chData chan int) {
		defer wg.Done()
		for i := 0; ; i++ {
			select {

			case <-ctx.Done():
				fmt.Println("stopping")
				return
			case chData <- i:
			}
		}

	}(ctx, chData)

	wg.Wait()
	cancel()
}
