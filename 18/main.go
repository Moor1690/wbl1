package main

import (
	"fmt"
	"sync"
)

//https://stackoverflow.com/questions/44949467/when-do-you-embed-mutex-in-struct-in-go

type counter struct {
	count int
	sync.Mutex
}

func main() {

	var wg sync.WaitGroup
	c := counter{count: 0}
	count := 150
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Lock()
			c.count++
			c.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(c.count)
}
