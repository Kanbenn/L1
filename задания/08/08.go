package main

import "fmt"

func main() {
	var num int64 = 33   // Исходное число
	position := uint(5)  // Позиция бита, который хотим изменить
	bitValue := int64(0) // Желаемое значение бита (0 или 1)

	result := setBit(num, position, bitValue)
	fmt.Printf("Исходное число: %d\n", num)
	fmt.Printf("Измененное число: %d\n", result)
}

func setBit(num int64, i uint, bitValue int64) int64 {
	mask := int64(1 << i) // Создаем маску с помощью побитового сдвига
	num &= ^mask          // Обнуляем i-й бит в числе, применяя побитовое отрицание на маску
	num |= bitValue << i  // Устанавливаем i-й бит в заданное значение
	return num
}
