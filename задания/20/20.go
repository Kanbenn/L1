package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"             // строка
	reversed := reverse(str)          //переворачиваем слова в строке
	words := strings.Fields(reversed) // переворачиваем строку
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	fmt.Println(strings.Join(words, " "))
}

func reverse(str string) string {
	words := strings.Split(str, " ")       //  убираем пробелы получаем массив
	reversed := make([]string, len(words)) // массив перевернутых
	for i := 0; i < len(words); i++ {
		reversed[i] = reverseWord([]rune(words[i])) // переворачиваем каждое слово в цикле
	}
	return strings.Join(reversed, " ") // возвращаем результат в строке объединяя пробелом
}

func reverseWord(str []rune) string {
	i := 0
	j := len(str) - 1
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
	return string(str)
}
