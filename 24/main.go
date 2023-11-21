package main

/*
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/
import (
	"fmt"
	"math"
)

// Определение структуры Point, представляющей координаты точки на плоскости.
type Point struct {
	x float64
	y float64
}

// Функция NewPoint создает новую точку с заданными координатами x и y и возвращает ее.
func NewPoint(x float64, y float64) Point {
	return Point{x, y}
}

// Метод distanceBetween вычисляет расстояние между текущей точкой (p) и переданной точкой (p2)
func (p Point) distanceBetween(p2 Point) float64 {
	return math.Pow(math.Pow((p.x-p2.x), 2)+math.Pow((p.y-p2.y), 2), 0.5)
}

func main() {
	p1 := NewPoint(3, 5)
	p2 := NewPoint(0, 0)

	// Вызов метода distanceBetween для точки p1 с передачей точки p2 в качестве аргумента.
	// Этот метод вычисляет расстояние между точками p1 и p2 и возвращает результат.
	distance := p1.distanceBetween(p2)

	fmt.Println(distance) // Выводим расстояние между точками на экран.
}
