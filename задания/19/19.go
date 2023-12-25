package main

import "fmt"

func main() {
	// переменная для строки
	var str string
	fmt.Println("Add string")
	fmt.Scanln(&str)      //ввод
	newstr := []rune(str) // преобразуем строку в руну для работы с unicode
	i := 0                // определяем начальную точку массива рун
	j := len(newstr) - 1  // определяем конечную точку массива рун

	// простой цикл в котором меняем местами значений массива рун по индексам i,j
	for i < j {
		newstr[i], newstr[j] = newstr[j], newstr[i]
		i++
		j--
	}
	// преобразуем руны в строку и выводим результат
	fmt.Println(string(newstr))
}
