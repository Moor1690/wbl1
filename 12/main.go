package main

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/
import (
	"fmt"
)

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{}) // Создание множества как карты

	for _, value := range sequence {
		set[value] = struct{}{} // Добавление элемента в множество
	}

	for key := range set {
		fmt.Println(key) // Вывод элементов множества
	}
	fmt.Println(set)
}
