package main

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/
import "fmt"

// Функция для выполнения быстрой сортировки массива arr в диапазоне от start до end.
func myQuicksort(arr *[]int, start int, end int) {
	if start < end {
		// Находим индекс опорного элемента после разделения.
		pivotIndex := partition(arr, start, end)

		// Рекурсивно сортируем левую и правую части массива относительно опорного элемента.
		myQuicksort(arr, start, pivotIndex-1)
		myQuicksort(arr, pivotIndex+1, end)
	}
}

// Функция для разделения массива arr на две части вокруг опорного элемента.
func partition(arr *[]int, start int, end int) int {
	pivot := (*arr)[end] // Опорный элемент.
	i := start - 1

	// Проходим по элементам в диапазоне start до end-1.
	for j := start; j < end; j++ {
		if (*arr)[j] < pivot {
			i++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i] // Обмен элементами, чтобы элементы < pivot были слева.
		}
	}

	// Помещаем опорный элемент на правильное место.
	(*arr)[i+1], (*arr)[end] = (*arr)[end], (*arr)[i+1]

	return i + 1 // Возвращаем индекс опорного элемента.
}

func main() {
	arr := []int{3, 5, 9, 2, 6, 5, 3, 5, 1, 4}
	myQuicksort(&arr, 0, len(arr)-1)

	fmt.Println(arr) // Вывод отсортированного массива.
}
