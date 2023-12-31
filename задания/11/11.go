package main

import "fmt"

func main() {
	data1 := []int{1, 2, 5, 7, 9}
	data2 := []int{4, 6, 8, 1, 2}
	fmt.Println(check(data1, data2))
}

// проверяем Set через Map
func check(data1, data2 []int) []int {
	mp := make(map[int]bool)
	var res []int

	for _, i := range data1 {
		mp[i] = true
	}

	for _, i := range data2 {
		if mp[i] == true {
			res = append(res, i)
		}
	}
	return res
}
