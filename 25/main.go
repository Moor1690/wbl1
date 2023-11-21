package main

import (
	"fmt"
	"time"
)

func mySleep(d time.Duration) {
	select {
	case <-time.After(d):
	}
}

func main() {
	start := time.Now()

	// Declaring durations
	mySleep(time.Duration(time.Second * 3))

	elapsed := time.Since(start)
	fmt.Printf("page took %s", elapsed)
}
