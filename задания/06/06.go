package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//В Go есть несколько способов остановить выполнение горутины:
//1. Возврат из функции: Возвращение из функции, вызывающей горутину, автоматически останавливает выполнение горутины.
//2. Использование канала: Вы можете использовать канал для уведомления горутины о необходимости остановки. Горутина может регулярно проверять значение в канале и завершить выполнение, если получит указанный сигнал.
//3. Использование флага: Вы можете использовать булеву переменную в качестве флага, которая будет указывать, когда горутина должна быть остановлена. Горутина может регулярно проверять значение этого флага и завершить выполнение, если значение true.
//4. Использование функции context: Пакет context в Go предоставляет возможность передавать сигналы о прекращении выполнения через контекст. Горутина может проверять контекст и завершить выполнение, если получит сигнал отмены.
//5. Использование таймера: Горутина может останавливаться после определенного времени с помощью использования таймера. Вы можете использовать пакет time для этого.

// использование таймера
func timerGoroutine() {
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			fmt.Println("Stopping goroutine...")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(time.Second)
		}
	}
}

// с использованием time
func timeGoroutine() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

// с использованием bool переменных
func boolGoroutine(ch chan bool) {
	for {
		select {
		case msg := <-ch:
			if msg == true {
				fmt.Println("Stopping goroutine...")
				return
			}
		default:
			fmt.Println("Working...")
			time.Sleep(time.Second)
		}
	}

}

// с использованием flag
func flagGoroutine() {
	defer wg.Done()
	for {
		if stopFlag {
			fmt.Println("Stopping goroutine...")
			return
		}
		fmt.Println("Working...")
		time.Sleep(time.Second)
	}
}

// с использованием context
func ctxGoroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping goroutine...")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

var stopFlag bool
var wg sync.WaitGroup

func main() {
	go timerGoroutine()
	time.Sleep(6 * time.Second)
	fmt.Println("Exiting... timerGoroutine")

	ch := make(chan bool)
	go boolGoroutine(ch)
	time.Sleep(5 * time.Second)
	ch <- true
	fmt.Println("Exiting... boolGoroutine")

	wg.Add(1)
	go flagGoroutine()
	time.Sleep(5 * time.Second)
	fmt.Println("Exiting... flagGoroutine")
	stopFlag = true
	wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	go ctxGoroutine(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("Exiting... ctxGoroutine")
	cancel()
	time.Sleep(1 * time.Second)

	go timeGoroutine()
	time.Sleep(5 * time.Second)
	fmt.Println("Exiting... timeGoroutine")
}
