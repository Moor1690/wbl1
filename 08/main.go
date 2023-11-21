package main

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/
import (
	"fmt"
)

func main() {
	var bit int64
	var someNumber int64 = 16

	fmt.Print("Введите номер бита:\t")
	fmt.Scan(&bit) // Запрашиваем у пользователя ввод числа bit.

	bit = 1 << bit // Выполняем битовый сдвиг влево на значение bit.

	someNumber = someNumber | bit // Выполняем побитовое ИЛИ с someNumber и bit.

	fmt.Printf("%b", someNumber) // Выводим результат в двоичном формате.
}
