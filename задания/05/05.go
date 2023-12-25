package main

import (
	"fmt"
	"github.com/rs/xid"
	"time"
)

func main() {
	var second time.Duration // переменная для секунд
	fmt.Println("Add second...")
	fmt.Scanln(&second)     // ввод
	ch := make(chan string) // канал
	go read(ch)             // воркер который читает
	//запись в горутину
	go func() {
		for {
			ch <- xid.New().String()
		}
	}()
	//спим столько сколько указали в секундах
	time.Sleep(second * time.Second)
	// закрываем канал и завершаем программу
	close(ch)
}

// функция чтения из канала и вывод значений в консоль
func read(ch chan string) {
	for i := range ch {
		fmt.Println(i)
	}
}
