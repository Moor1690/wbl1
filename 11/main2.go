package main

import "fmt"

func intersection(arr1 []int, arr2 []int) map[int]int {
	var arrResult = make(map[int]int)
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				arrResult[arr1[i]] = arr1[i]
			}
		}
	}
	return arrResult
}

func main() {
	arr1 := []int{2, 4, 7, 1, 9, 3, 8}
	arr2 := []int{9, 2, 6, 1, 7, 5, 0, 8, 8}

	//arrResult = make(map[int]int)

	if len(arr1) > len(arr2) {
		fmt.Println(intersection(arr1, arr2))
	} else {
		fmt.Println(intersection(arr2, arr1))
	}
}
