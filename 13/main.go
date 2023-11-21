package main

/*
Поменять местами два числа без создания временной переменной.
*/
import "fmt"

func main() {
	a := -2
	b := 5
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("a:%d\nb:%d\n", a, b)

	a, b = b, a

	fmt.Printf("a:%d\nb:%d\n", a, b)
}
