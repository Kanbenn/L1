package main

import (
	"fmt"
)

func main() {
	sl := []int{1, 2, 3, 4, 5}

	fmt.Println(sl)
	sl = deleteAtIndex(sl, 2)
	fmt.Println(sl)
	// fmt.Println(deleteAtIndex(sl, 1))
	// fmt.Println(sl)
}

func deleteAtIndex[T any](in []T, i int) []T {
	// if i < 0 || i > len(in)-1 {
	// 	return in
	// }
	// out := append(in[:i], in[i+1:]...)
	// copy(in[i:], in[i+1:])
	// return in[:len(in)-1]
	return append(in[:i], in[i+1:]...)
	// slices.Delete[]()
	// a = append(a[:i], a[i+1:]...)
}
