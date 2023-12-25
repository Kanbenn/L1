package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

// для этого задания очень подходит Конкурентный тип вычислений,
// который характеризуется использованием мьютексов.

// result struct для обвешивания слайса мьютекс-локерами.
type ConcurrentSlice struct {
	sl []int
	mu sync.RWMutex
}

func main() {
	cs := ConcurrentSlice{
		sl: []int{2, 4, 6, 8, 10},
		mu: sync.RWMutex{},
	}

	var wg sync.WaitGroup

	for i, el := range cs.sl {
		wg.Add(1)
		go func(index, x int) {
			defer wg.Done()
			square := x * x
			cs.sl[index] = square
		}(i, el)
	}

	wg.Wait()

	fmt.Println(cs.sl)
}
