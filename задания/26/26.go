package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"

	fmt.Println(hasUniqueChars(s1), s1)
	fmt.Println(hasUniqueChars(s2), s2)
	fmt.Println(hasUniqueChars(s3), s3)
}

func hasUniqueChars(s string) bool {
	runes := []rune(strings.ToLower(s))
	slices.Sort(runes)

	for i := 1; i < len(runes); i++ {
		if runes[i] == runes[i-1] {
			return false
		}
	}
	return true
}
