package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) Point {
	return Point{x, y}
}

func (p Point) distanceBetween(p2 Point) float64 {
	return math.Pow(math.Pow((p.x-p2.x), 2)+math.Pow((p.y+p2.y), 2), 0.5)
}

func main() {
	p1 := NewPoint(3, 5)
	p2 := NewPoint(0, 0)

	fmt.Println(p1.distanceBetween(p2))
}
