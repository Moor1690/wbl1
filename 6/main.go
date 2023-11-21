package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stoppin gorutin1")
				return
			default:
				fmt.Println("some work1")
			}
		}
	}(ctx, &wg)

	time.Sleep(time.Microsecond * 20)
	cancel()
	wg.Wait()
	ctx2, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Microsecond*20))

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stoppin gorutin2")
				return
			default:
				fmt.Println("some work2")
			}
		}
	}(ctx2)
	wg.Wait()

	stopChannel := make(chan struct{})

	wg.Add(1)
	go func(sChannel chan struct{}) {
		defer wg.Done()

		for {
			select {
			case <-sChannel:
				fmt.Println("stoppin gorutin3")
				return
			default:
				fmt.Println("some work3")
			}
		}
	}(stopChannel)

	time.Sleep(time.Microsecond * 20)
	close(stopChannel)
	wg.Wait()

	ctx3, _ := context.WithTimeout(context.Background(), 20)

	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stoppin gorutin4")
				return
			default:
				fmt.Println("some work4")
			}
		}
	}(ctx3, &wg)

	cancel()
	wg.Wait()
}
