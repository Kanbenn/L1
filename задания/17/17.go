package main

import "fmt"

func main() {
	target := 2
	arr := []int{1, 2, 3, 4, 5, 6, 10}
	fmt.Println(binarysearch(arr, target))
}

// работает в остсортированном массиве
func binarysearch(arr []int, target int) int {
	left := 0             // определяем наименьшее значение
	right := len(arr) - 1 // наибольшее значение

	for left <= right {
		val := (left + right) / 2 //среднее двух этих значений
		if arr[val] == target {   // если совпадает то возвращаем индекс числа
			return val
		} else if arr[val] > target { // если больше искомого числа уменьшаем большую часть на val-1
			right = val - 1
		} else if arr[val] < target { // если меньше искомого числа увеличиваем наименьшую часть val+1
			left = val + 1
		}
	}
	return 0 // если не нашли возвращаем 0
}
