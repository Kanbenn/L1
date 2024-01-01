package main

import (
	"fmt"
	"math"
)

// Задание: Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными
// параметрами x,y и конструктором.

type Point struct {
	x float64
	y float64
}

// NewPoint это функция-фабрика, реализующая конструктор структуры Point.
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// формулу взял здесь: https://en.wikipedia.org/wiki/Distance
func Distance(p1, p2 Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	p1 := NewPoint(5, 3)
	p2 := NewPoint(10, 20)
	fmt.Printf("%.2f \n", Distance(p1, p2))
}
