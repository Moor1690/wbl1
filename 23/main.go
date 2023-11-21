package main

import "fmt"

func removeElement(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	arr := []int{3, 6, 786, 1, 47, 4, 9, 1, 9}
	fmt.Println(arr)

	fmt.Println(removeElement(arr, 2))
}
