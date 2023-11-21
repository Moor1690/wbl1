package main

/*
Реализовать пересечение двух неупорядоченных множеств.
*/
import "fmt"

// Функция intersection принимает два слайса arr1 и arr2, и возвращает мапу, содержащую пересекающиеся элементы.
func intersection(arr1 []int, arr2 []int) map[int]struct{} {
	// Создаем пустую мапу, в которой ключами будут пересекающиеся элементы.
	var arrResult = make(map[int]struct{})

	// Итерируемся по элементам arr1.
	for i := 0; i < len(arr1); i++ {
		// Для каждого элемента arr1, итерируемся по элементам arr2.
		for j := 0; j < len(arr2); j++ {
			// Если элементы arr1[i] и arr2[j] равны, то добавляем его в мапу arrResult.
			if arr1[i] == arr2[j] {
				arrResult[arr1[i]] = struct{}{}
			}
		}
	}
	return arrResult
}

func main() {
	arr1 := []int{2, 4, 7, 1, 9, 3, 8}
	arr2 := []int{9, 2, 6, 1, 7, 5, 0, 8, 8}

	// Вызываем функцию intersection и выводим результат на стандартный вывод.
	if len(arr1) > len(arr2) {
		fmt.Println(intersection(arr1, arr2))
	} else {
		fmt.Println(intersection(arr2, arr1))
	}
}
