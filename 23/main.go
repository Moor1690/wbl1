package main

/*
Удалить i-ый элемент из слайса.
*/
import "fmt"

func removeElement(slice []int, i int) []int {
	// Срез разбивается на две части: до индекса i (включительно) и после него.
	// Затем эти две части объединяются с помощью операции append.
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	arr := []int{3, 6, 786, 1, 47, 4, 9, 1, 9}
	fmt.Println(arr)

	fmt.Println(removeElement(arr, 2)) // Выводим измененный срез после удаления элемента.
}
