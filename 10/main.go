package main

import (
	"fmt"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	result := make(map[int][]float64)

	for _, temp := range arr {

		key := int(temp/10.0) * 10
		result[key] = append(result[key], temp)
	}

	for key, temps := range result {
		fmt.Printf("%d: %v\n", key, temps)
	}
}
