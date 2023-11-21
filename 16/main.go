package main

import "fmt"

func myQuicksort(arr *[]int, start int, end int) {
	if start < end {
		pivotIndex := partition(arr, start, end)
		myQuicksort(arr, start, pivotIndex-1)
		myQuicksort(arr, pivotIndex+1, end)
	}
}

func partition(arr *[]int, start int, end int) int {
	pivot := (*arr)[end]
	i := start - 1
	for j := start; j < end; j++ {
		if (*arr)[j] < pivot {
			i++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i] // обмен элементами
		}
	}
	(*arr)[i+1], (*arr)[end] = (*arr)[end], (*arr)[i+1] // помещаем опорный элемент на место
	return i + 1
}

func main() {

	arr := []int{3, 5, 9, 2, 6, 5, 3, 5, 1, 4}
	myQuicksort(&arr, 0, len(arr)-1)

	fmt.Println(arr)

}
