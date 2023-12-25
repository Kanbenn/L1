package main

import "fmt"

func main() {
	a := 5
	b := 10
	a, b = b, a
	fmt.Println(a, b)
	x := 1
	y := 3
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(x *int, y *int) {
	temp := *y // сохраняет значение по адресу *y в переменную
	*y = *x    // присваиваем значение по адресу *x значению по адресу *y
	*x = temp  // присваиваем значению по адресу *x значение temp
}
