package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"

	arr := strings.Fields(str)

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i] // Меняем местами слова
	}
	str = strings.Join(arr, " ") // Соединяем слова обратно в строку

	fmt.Println(str)
}
