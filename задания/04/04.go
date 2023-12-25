package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//канал для коммуникации между воркерами
	ch := make(chan int)

	//кол-во воркеров
	var num int
	// парсинг количества воркеров из аргумента командной строки -w
	flag.IntVar(&num, "w", 10, "Введите число - количество воркеров. По умолчанию 10.")
	flag.Parse()

	// создание переменной context, которая получит сигнал уведомления о нажатии ctrl+c пользователем
	// и сообщит об этом сигнале всем, кто подписан на этот контекст.
	ctx, _ := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	fmt.Println("Эта программа завершится по нажатию клавиш Ctrl+c")

	//Создание указанного пользователем количества воркеров
	//Для каждого воркера запускается функция reader(c) в отдельной горутине
	for i := 0; i < num; i++ {
		go reader(ch)
	}

	//запускает в новой горутине функцию writer, которая
	//будет записывать данные в канал
	writer(ch, ctx)

	// "Graceful Shutdown по нажатию Ctrl+c"
	// <-ctx.Done()

}

func writer(ch chan int, ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\nGraceful Shutdown по нажатию Ctrl+c")
			fmt.Println("writer получил сигнал завершения работы, останавливаю горутину")
			os.Exit(0) // немедленное завершение работы программы с кодом выхода 0
			// return
		default:
			time.Sleep(time.Second) // пауза в 1 секунду
			ch <- i                 // отправка в канал
			i++
		}
	}
}

func reader(ch chan int) {
	for val := range ch {
		fmt.Print(val)
	}
}
