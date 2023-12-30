package main

import (
	"fmt"
	"slices"
)

// в голанге нет типа "множество" как set в питоне. можно сымитировать через тип Map.
// но конкретно для этой задачи более эффективный алгоритм - сортировка слайса.

func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println("начальный слайс:", input)

	fmt.Println("setViaMap:", setViaMap(input))

	fmt.Println("setViaSortedSlice:", setViaSortedSlice(input))
}

func setViaMap(in []string) map[string]struct{} {
	// пустой struct занимает 0 байт, в отличии от bool который 1 байт.
	// в голанге это юзают как лайфхак для заглушки в Мапах или
	// пустого синала в Каналах. Смотри также: Устный Вопрос №4.
	m := make(map[string]struct{}, len(in))
	for _, i := range in {
		m[i] = struct{}{}
	}
	return m
}

func setViaSortedSlice(in []string) []string {
	slices.Sort(in)

	out := make([]string, 0, len(in))
	var prev string
	for _, cur := range in {
		if cur != prev {
			out = append(out, cur)
			prev = cur
		}
	}
	// обрезаю capacity до получившегося количества элементов.
	return out[:len(out):len(out)]
}
