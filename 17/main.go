package main

import "fmt"

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2 // Вычисление среднего индекса

		if arr[mid] == target {
			return mid // Элемент найден
		} else if arr[mid] < target {
			left = mid + 1 // Искать в правой половине
		} else {
			right = mid - 1 // Искать в левой половине
		}
	}

	return -1 // Элемент не найден
}

func main() {
	arr := []int{2, 3, 4, 10, 40} // Отсортированный срез
	target := 10

	result := binarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Элемент найден на индексе %d\n", result)
	} else {
		fmt.Println("Элемент не найден в массиве")
	}
}
