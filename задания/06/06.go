package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Задание:
// Реализовать все возможные способы остановки выполнения горутины.

// Ответ:
// Зная коварность менторов WB, нужно уточнить вопрос: как Завершить горутину либо
// временно Приостановить её выполнение и вернуть управление рантайму?

// Завершить горутину извне в Go возможности нет (на момент версии го 1.22),
// потому и возникают ошибки deadlock или all goroutines are asleep.
// В этом большое отличие легковесных горутин рантайма Go от полновесных процессов ОС,
// которые могут реагировать на SIG INTERRUPT, SIG TERMINATE и тд, а в крайнем
// случае SIGKILL грубо вырубает процесс, не дожидаясь от него ответа.

// Горутина приостанавливается и передаёт управление рантайму (sysmon) при:
// вызове других горутин/функций, ожидание канала, мютекса, таймера, контекста и др.
// Ещё один способ приостановить горутину - запустить сборщик мусора вызовом runtime.gogc()
// (применяется в редких случаях, для тестов)

// Завершиться горутина должна сама инструкцией return.

// Для само-регуляции приостанавливающихся горутин в голанге юзают
// бесконечный цикл с select-switch'ем внутри, ожидающем сигнала
// дальнейших действий из одного или нескольких внешних источников.

// использование таймера
func timerGoroutine() {
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			fmt.Println("Stopping timerGoroutine...")
			return
		default:
			fmt.Println("timerGoroutine working...")
			time.Sleep(time.Second)
		}
	}
}

// с использованием time
func timeGoroutine() {
	fmt.Println("timeGoroutine sleeping...")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	fmt.Println("timeGoroutine woke up.")
}

// с использованием bool переменных
func boolGoroutine(ch chan struct{}) {
	for {
		select {
		case <-ch:
			fmt.Println("Stopping boolGoroutine...")
			return
		default:
			fmt.Println("boolGoroutine working...")
			time.Sleep(time.Second)
		}
	}

}

// с использованием flag
func flagGoroutine() {
	defer wg.Done()
	for {
		if stopFlag {
			fmt.Println("Stopping flagGoroutine...")
			return
		}
		fmt.Println("flagGoroutine working...")
		time.Sleep(time.Second)
	}
}

// с использованием context
func ctxGoroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping ctxGoroutine...")
			return
		default:
			fmt.Println("ctxGoroutine working...")
			time.Sleep(1 * time.Second)
		}
	}
}

var stopFlag bool
var wg sync.WaitGroup

func main() {
	go timerGoroutine()
	time.Sleep(6 * time.Second)

	ch := make(chan struct{})
	go boolGoroutine(ch)
	time.Sleep(5 * time.Second)
	ch <- struct{}{}

	wg.Add(1)
	go flagGoroutine()
	time.Sleep(5 * time.Second)
	stopFlag = true
	wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())
	go ctxGoroutine(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

	go timeGoroutine()
	time.Sleep(5 * time.Second)
}
