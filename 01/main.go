package main

import "fmt"

// Структура Human с каким то полем
type Human struct {
	someValue int
}

// Метода someFunk для стуктуры Human
func (h Human) someFunk() {
	fmt.Printf("someFunk print value: %d\n", h.someValue)
}

// Определение типа Action, который встраивает структуру Human
// Это значит, что Action наследует все поля и методы Human
type Action struct {
	Human
}

func main() {
	// Создание объекта типа Action, встроенная структура Human инициализируется значением 123
	a := Action{Human{123}}

	// Вызов метода someFunk для объекта Action
	// Поскольку Action встраивает Human, он может вызывать его методы напрямую
	a.someFunk() // Выведет "someFunk print value: 123\n" в консоль

	// Доступ к полю someValue встроенной структуры Human через объект Action
	fmt.Println(a.someValue) // Выведет "123" в консоль
}
