package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := big.NewInt(335544323213213131)
	b := big.NewInt(34359738368)

	// Умножение
	mul := big.NewInt(0)
	mul.Mul(a, b)
	fmt.Println("Умножение:", mul)

	// Деление
	div := big.NewInt(0)
	if b.Cmp(big.NewInt(0)) != 0 { // ограничение деления на 0
		div.Div(a, b)
	}
	fmt.Println("Деление:", div)

	// Сложение
	add := big.NewInt(0)
	add.Add(a, b)
	fmt.Println("Сложение:", add)

	// Вычитание
	sub := big.NewInt(0)
	sub.Sub(a, b)
	fmt.Println("Вычитание:", sub)
}
