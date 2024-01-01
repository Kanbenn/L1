package main

import (
	"fmt"
	"time"
)

// Задание: Реализовать собственную функцию sleep.

// Ответ: имплементация функции time.Sleep в го прячется в пакете runtime
// https://github.com/golang/go/blob/master/src/runtime/time.go#L178

// И реализована не через бесконечный цикл (который крутил бы понапрасну процессор),
// а через gopark() - аналог инструкции suspend, await и тд в других языках.
// То есть, time.Sleep возвращает управление выше, в рантайм и просит возобновить
// работу горутины через указанное количество времени.

func MySleepFunc(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Спим 5 секунд")

	sleepDuration := 5
	MySleepFunc(sleepDuration)

	fmt.Println("Проспались")
}
