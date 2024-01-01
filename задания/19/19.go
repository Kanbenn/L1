package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {

	fmt.Println("Введите строку для разворота:")

	// fmt.Scanln почему-то читает не всю строку, а лишь одно слово
	// до пробела, поэтому сделал через bufio.
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// убираю ненужные символы перевода каретки
	str := strings.ReplaceAll(input, "\r\n", "")
	fmt.Printf("%q \n", str)

	runes := []rune(str)
	fmt.Printf("%q \n", runes)
	slices.Reverse(runes)

	fmt.Println(string(runes))
}
