package main

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
*/
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
	str1 := "asdfg"
	str2 := "abcd"
	str3 := "abCdefAaf"
	str4 := "aabcd"

	fmt.Println(uniqueString(str1))
	fmt.Println(uniqueString(str2))
	fmt.Println(uniqueString(str3))
	fmt.Println(uniqueString(str4))

}
