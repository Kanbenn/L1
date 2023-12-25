package main

// Реализовать конкурентную запись данных в map.
// использую мьютекс для реализации доступа к map

import (
	"fmt"
	"sync"
)

type MyMap struct {
	mu sync.RWMutex
	mp map[int]int
}

func (s *MyMap) Add(key int, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.mp[key] = value
}

func main() {
	m := MyMap{
		mu: sync.RWMutex{},
		mp: make(map[int]int),
	}

	var wg sync.WaitGroup
	numOfGoroutines := 10

	for i := 0; i < numOfGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			m.Add(i, i)
			wg.Done()
		}(i)
	}

	wg.Wait() //ждем выполнения всех воркеров

	fmt.Println(m.mp)
}
