package main

import "fmt"

// в го нет метода реализации множеств, но можем использовать тим мап
func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(set(data))
}

func set(data []string) map[string]bool {
	mp := make(map[string]bool)
	for _, i := range data {
		mp[i] = true
	}
	return mp
}
