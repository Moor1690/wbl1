package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	myMap := map[int]int{
		0: 0,
		1: 2,
		2: 4,
		3: 8,
		4: 16,
	}

	for i, _ := range myMap {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			myMap[i] = myMap[i] * myMap[i]
		}(i)
	}

	wg.Wait()

	fmt.Println(myMap)
}
