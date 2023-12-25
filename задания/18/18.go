package main

import (
	"fmt"
	"sync"
)

// структура счетчика
type Counter struct {
	val   int
	valMu sync.Mutex // чтобы предотвратить гонки данных
}

// метод структуры Counter
func (c *Counter) Count() {
	c.valMu.Lock()   // лочим
	c.val++          // увеличиваем счетчик
	c.valMu.Unlock() // анлочим
}

func main() {
	count := Counter{
		val: 0, // инициализируем val
	}
	var wg sync.WaitGroup // вейт группа для обеспечения количества горутинк

	wg.Add(10)
	// цикл фор
	for i := 0; i < 10; i++ {
		// запуск воркера
		go func() {
			count.Count() // запуск метода
			wg.Done()     // подтверждение что воркер отработал
		}()
	}

	wg.Wait()              // ожидание завершения всех воркеров
	fmt.Println(count.val) // вывод получившегося числа
}
