package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

// Задание:
// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

// Идитоматичный подход к горутинам в го такой: во избежание утечек памяти,
// горутины обязательно нужно завершать. Принято это делать через закрытие
// канала. Закрывать канал должен writer, тогда reader'ы закроются автоматически.

// Задание №5 cоставлено так, что закрыть канал отправителем не получится,
// поэтому можно обернуть канал в мьютекс sync.Once и проверять не закрыт ли уже канал
// при каждой отправке writer'ом. Иначе, отправка в закрытый канал это deadlock.
// подробнее: https://go101.org/article/channel-closing.html

type MyChannel struct {
	ch       chan string
	isClosed bool
	once     sync.Once
}

func (mc *MyChannel) SafeClose() {
	fmt.Println("SafeClose: сэйфово закрываю канал")
	mc.once.Do(func() {
		close(mc.ch)
		mc.isClosed = true
	})
}

func main() {
	var seconds int
	flag.IntVar(&seconds, "s", 5, "количество секунд работы программы. дефолт: 10")
	flag.Parse()
	fmt.Println("Программа завершится через", seconds, "секунд.")

	mc := &MyChannel{ch: make(chan string)}

	wg := &sync.WaitGroup{}
	wg.Add(4) // 2 readera + 2 writera

	go writer(mc, wg) // запуск отправителей
	go writer(mc, wg)

	go reader(mc, wg) // запуск получателей
	go reader(mc, wg)

	// даём программе отработать указанное количество секунд
	time.Sleep(time.Second * time.Duration(seconds))
	fmt.Println("\n", "main: время вышло")

	mc.SafeClose() // закрываем канал

	wg.Wait() // ждём завершения работы всех горутин
}

func writer(mc *MyChannel, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("writer: начинаю отправку данных в канал")
	for {
		if mc.isClosed {
			fmt.Println("writer: канал закрыт, выходим")
			return
		}
		mc.ch <- "."
		time.Sleep(100 * time.Millisecond) // пауза 100 мс чтобы не заспамить консоль.
	}
}

func reader(mc *MyChannel, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("reader: начинаю читать канал")

	// range среагирует на закрытие канала и горутина завершится
	for data := range mc.ch {
		fmt.Print(data)
	}
	fmt.Println("reader: читать больше нечего, выходим")
}
