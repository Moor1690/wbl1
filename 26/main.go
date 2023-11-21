package main

import (
	"fmt"
	"strings"
)

func uniqueString(s string) bool {
	s = strings.ToLower(s)
	m := make(map[rune]struct{}) // Используем мапу для отслеживания уникальных символов
	for _, val := range s {
		if _, exists := m[val]; exists {
			return false // Если символ уже встречался, строка не уникальна
		}
		m[val] = struct{}{} // Добавляем символ в мапу
	}
	return true // Все символы уникальны
}

func main() {
	str := "asdfg"
	fmt.Println(uniqueString(str)) // Выведет false, так как есть повторяющиеся символы
}
