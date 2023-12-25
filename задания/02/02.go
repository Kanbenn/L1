package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

// result struct для сохранения порядка элементов исходного слайса
type result struct {
	i int
	n int
}

func main() {
	in := []int{2, 4, 6, 8, 10}

	resultCh := make(chan result)
	var wg sync.WaitGroup

	for i := 0; i < len(in); i++ {
		wg.Add(1)
		go func(index, n int) {
			defer wg.Done()
			square := n * n
			resultCh <- result{i: index, n: square}
		}(i, in[i])
	}

	go func() {
		for res := range resultCh {
			in[res.i] = res.n
		}
	}()

	wg.Wait()
	close(resultCh)

	fmt.Println(in)
}
