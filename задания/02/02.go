package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {

	slice := []int{2, 4, 6, 8, 10}
	fmt.Println("начальный слайс: \n", slice)

	fmt.Println("результат выполнения задания c помощью мьютексов: \n", SquaredSliceWithMutex(slice))
	fmt.Println("результат выполнения задания c помощью каналов: \n", SquaredSliceWithChannels(slice))

}

func SquaredSliceWithMutex(in []int) []int {
	sliceCopy := append(in[:0:0], in...)

	mu := sync.RWMutex{}
	var wg sync.WaitGroup

	for i, el := range sliceCopy {
		wg.Add(1)
		go func(index, n int) {
			defer wg.Done()
			square := n * n
			mu.Lock()
			sliceCopy[index] = square
			mu.Unlock()
		}(i, el)
	}
	wg.Wait()
	return sliceCopy
}

func SquaredSliceWithChannels(in []int) []int {
	sliceCopy := append(in[:0:0], in...)

	var wg sync.WaitGroup

	actionCh := make(chan func())

	go func() {
		for {
			f := <-actionCh
			f()
		}
	}()

	for i, el := range sliceCopy {
		wg.Add(1)
		go func(index, n int) {
			f := func() {
				square := n * n
				sliceCopy[index] = square
				wg.Done()
			}
			actionCh <- f
		}(i, el)
	}
	wg.Wait()

	return sliceCopy
}
