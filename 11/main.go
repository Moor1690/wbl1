package main

/*
Реализовать пересечение двух неупорядоченных множеств.
*/
// Работает только с множествами хранящими уникальные значения
import (
	"fmt"
)

// Функция intersection принимает два слайса arr1 и arr2 и возвращает слайс arrResult,
// содержащий пересекающиеся элементы между arr1 и arr2.
func intersection(arr1 []int, arr2 []int) []int {
	// Создаем пустой слайс arrResult для хранения пересекающихся элементов.
	var arrResult []int

	// Итерируемся по элементам слайса arr1.
	for i := 0; i < len(arr1); i++ {
		// Вложенный цикл для сравнения элементов arr1 с элементами arr2.
		for j := 0; j < len(arr2); j++ {
			// Если элементы arr1[i] и arr2[j] равны, то добавляем его в arrResult с использованием append.
			if arr1[i] == arr2[j] {
				arrResult = append(arrResult, arr1[i])
				// Используем continue, чтобы продолжить поиск других совпадений.
			}
		}
	}

	// Возвращаем слайс arrResult, содержащий пересекающиеся элементы.
	return arrResult
}

func main() {
	// Создаем два слайса arr1 и arr2 с целочисленными значениями.
	arr1 := []int{2, 4, 7, 1, 9, 3, 8}
	arr2 := []int{9, 2, 6, 1, 7, 5, 0}

	// Проверяем, из какого из двух слайсов arr1 и arr2 больше элементов,
	// и вызываем функцию intersection с этими слайсами в качестве аргументов.
	if len(arr1) > len(arr2) {
		fmt.Println(intersection(arr1, arr2))
	} else {
		fmt.Println(intersection(arr2, arr1))
	}
}
