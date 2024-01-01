package main

import "fmt"

func main() {
	a, b := 5, 10
	fmt.Println("a and b before swap:", a, b)
	genericSwap(&a, &b)
	fmt.Println("a and b after  swap:", a, b)

	s1, s2 := "left", "right"
	fmt.Println("s1 and s2 before swap:", s1, s2)
	genericSwap(&s1, &s2)
	fmt.Println("s1 and s2 after  swap:", s1, s2)
}

func genericSwap[T any](left, right *T) {
	*left, *right = *right, *left
}
