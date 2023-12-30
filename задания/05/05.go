package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"
)

type MyChannel struct {
	ch       chan string
	isClosed bool
	once     sync.Once
}

func (mc *MyChannel) SafeClose() {
	log.Println("SafeClose: сэйфово закрываю канал")
	mc.once.Do(func() {
		close(mc.ch)
		mc.isClosed = true
	})
}

func main() {
	var seconds int
	flag.IntVar(&seconds, "s", 10, "количество секунд работы программы. дефолт: 10")
	flag.Parse()
	fmt.Println("количество секунд работы программы:", seconds)

	mc := &MyChannel{ch: make(chan string)}

	wg := &sync.WaitGroup{}
	wg.Add(4)

	go writer(mc, wg) // запуск пишущей горутины
	go writer(mc, wg) // запуск пишущей горутины

	go reader(mc, wg) // запуск читающей горутины
	go reader(mc, wg) // запуск читающей горутины

	// даём программе отработать указанное количество секунд
	time.Sleep(time.Second * time.Duration(seconds))
	log.Println("main: время вышло")

	mc.SafeClose() // закрываем канал

	// ждём когда writer закроет канал и завершится,
	// после чего range в reader'e среагирует и прекратит чтение.
	wg.Wait()

}

func writer(mc *MyChannel, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if mc.isClosed {
			log.Println("writer: канал закрыт, выходим")
			return
		}
		mc.ch <- "."
		time.Sleep(100 * time.Millisecond) // пауза 100 мс чтобы не спамить в консоль.
	}
}

func reader(mc *MyChannel, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("reader: начинаю отправку данных в канал")
	for data := range mc.ch { // range корректно среагирует на раннее закрытие канала
		fmt.Print(data)
	}
	log.Println("reader: читать больше нечего, выходим")
}
