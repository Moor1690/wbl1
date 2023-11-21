package main

/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
*/
import "fmt"

func reverseString(input string) string {
	// Преобразование строки в срез рун для поддержки Unicode
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// Меняем местами элементы с начала и конца
		runes[i], runes[j] = runes[j], runes[i]
	}
	// Преобразование среза рун обратно в строку
	return string(runes)
}

func main() {
	inputS := "главрыба"
	reversed := reverseString(inputS)
	fmt.Printf("Оригинальная строка: %s\nПеревернутая строка: %s\n", inputS, reversed)
}
