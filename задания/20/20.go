package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Введите несколько слов для их разворота:")

	// fmt.Scanln почему-то читает не всю строку, а лишь одно слово
	// до пробела, поэтому сделал через bufio.
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// убираю ненужные символы перевода каретки
	str := strings.ReplaceAll(input, "\r\n", "")
	fmt.Printf("%q \n", str)

	sliceOfStrings := strings.Split(str, " ")

	slices.Reverse(sliceOfStrings)

	fmt.Println(strings.Join(sliceOfStrings, " "))
}
