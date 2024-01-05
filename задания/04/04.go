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

// Задание:
// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
// которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C.
// Выбрать и обосновать способ завершения работы всех воркеров.

// На устных экзаменах L1/L2 встречается коварный вопрос по
// этой теме: "Сколько горутин могут одновременно писать в stdout?"
// Правильный ответ: пытаться писать могут хоть все сразу,
// но внутри мьютексы и в моменте пишет только одна горутина.

// Под капотом, os.Stdout в голанге это файловый дескриптор
// в исходниках: Go\src\os\file.go NewFile(uintptr(syscall.Stdout), "/dev/stdout")
// а значит го-рантайм сделает синхронный syscall к операционной системе,
// в отличии от а-синхронного syscall'a к netpoller'y при сетевых запросах.
// Операционная система или Докер могут ставить лимит на количество доступных
// файловых дескриптэров.

// struct{} в каналах это пустая заглушка.
// смотри Устный Вопрос №4.
type myChan chan struct{}

func main() {

	ch := make(myChan, 1)

	var num int
	// парсинг количества воркеров из аргумента командной строки -w
	flag.IntVar(&num, "w", 10, "Введите число - количество воркеров.")
	flag.Parse()

	// переменная context получит уведомление от операционной системы
	// о нажатии ctrl+c пользователем. Этот сигнал далее передаётся всем,
	// кто подписан на данный контекст.
	ctx, _ := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	fmt.Println("Эта программа завершится по нажатию клавиш Ctrl+c")

	//Создание указанного пользователем количества воркеров
	//Для каждого воркера запускается функция reader(c) в отдельной горутине
	for i := 0; i < num; i++ {
		go reader(ch, i)
	}

	//запускает в новой горутине функцию writer, которая
	//будет записывать данные в канал
	writer(ch, ctx)
}

func writer(ch myChan, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\n writer: Graceful Shutdown по нажатию Ctrl+c")

			// закрытие канала остановит циклы for-range во всех reader'ах,
			// благодаря чему все читающие горутины сэйфово завершатся без утечек памяти.
			close(ch)
			return
		default:
			time.Sleep(100 * time.Millisecond) // пауза чтобы не заспамить консоль
			ch <- struct{}{}                   // отправка буддийской пустоты в канал
		}
	}
}

func reader(ch myChan, id int) {
	// for i := range ch {
	for range ch {
		fmt.Print(id, " ") // принт в консоль id'шника горутины
	}
}
