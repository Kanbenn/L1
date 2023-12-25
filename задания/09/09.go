package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //массив для перебора
	ch1 := make(chan int)                       //канал который принимает значения массива
	ch2 := make(chan int)                       //канал который принимает квадраты значений массива
	go Stepen(ch1, ch2)                         //запуск функции возведения в квадрат и записи во второй канал
	go Read(ch2)                                // запуск функции вывода квадратов занчений массива из канала
	for _, i := range arr {
		ch1 <- i
	}
}

func Stepen(ch1 chan int, ch2 chan int) {
	for {
		select {
		case val := <-ch1:
			ch2 <- val * val
		default:
			fmt.Println("No data in ch1")
			time.Sleep(time.Second)
		}
	}
}

func Read(ch2 chan int) {
	for {
		select {
		case val := <-ch2:
			fmt.Print(val, "\t")
		default:
			fmt.Println("No data in ch2")
			time.Sleep(time.Second)
		}
	}

}
