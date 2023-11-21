package main

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.


Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/
import (
	"fmt"
)

func main() {
	// Создаем слайс с числами с плавающей точкой.
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем пустую мапу, где ключами будут диапазоны чисел, а значениями - сами числа.
	result := make(map[int][]float64)

	// Итерируемся по элементам слайса.
	for _, temp := range arr {
		// Вычисляем ключ, который представляет диапазон для текущего числа.
		key := int(temp/10.0) * 10

		// Добавляем текущее число в соответствующий диапазон в мапе.
		result[key] = append(result[key], temp)
	}

	// Выводим результат группировки.
	for key, temps := range result {
		fmt.Printf("%d: %v\n", key, temps)
	}
}
