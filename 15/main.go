package main

//https://ru.stackoverflow.com/questions/1446734/%D0%98%D0%B7%D0%BC%D0%B5%D0%BD%D0%B5%D0%BD%D0%B8%D0%B5-%D0%B3%D0%BB%D0%BE%D0%B1%D0%B0%D0%BB%D1%8C%D0%BD%D0%BE%D0%B9-%D0%BF%D0%B5%D1%80%D0%B5%D0%BC%D0%B5%D0%BD%D0%BD%D0%BE%D0%B9

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100])) // Копирование необходимых данных в новую строку
}

func main() {
	someFunc()
}
