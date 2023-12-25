package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// в NewPoint определен конструктор, который создает новую точку с заданными координатами x и y
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// вычисляем расстояние между двумя точками
func Search(p1, p2 Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// создаем точки, вызываем Search, выводим результат
	p1 := NewPoint(1, 3)
	p2 := NewPoint(10, 3)
	fmt.Println(Search(p1, p2))
}
