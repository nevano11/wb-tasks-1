package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

// 1 Без редактирования тела горутины
// 1.1 Дождаться выполнения горутины (с использованием WaitGroup)
// 1.2 Завершить родительскую горутину

// 2 С редактированием тела горутины
// 2.1 Дождаться в теле горутины сигнала в виде чтения из канала
// 2.2 Флажок

// .
// 1.1 Дождаться выполнения горутины (с использованием WaitGroup)
func Test6_1(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer fmt.Println("Goroutine end")
		fmt.Println("Goroutine start")
		time.Sleep(time.Second)
	}()

	fmt.Println("main")
	wg.Wait()
}

// 1.2 Завершить родительскую горутину
func Test6_2(t *testing.T) {
	go func() {
		defer fmt.Println("Goroutine end")
		fmt.Println("Goroutine start")
		time.Sleep(time.Second)
	}()

	fmt.Println("main")
}

// 2.1 Дождаться в теле горутины сигнала в виде чтения из канала
func Test6_3(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	intCh := make(chan int, 1)

	go func(intCh <-chan int) {
		defer wg.Done()

		for {
			select {
			case <-intCh:
				fmt.Printf("Stopping\n")
				return
			default:
				time.Sleep(time.Duration(2) * time.Second)
			}
		}
	}(intCh)

	fmt.Println("main")
	time.Sleep(time.Second)
	fmt.Println("send 1000")
	intCh <- 1000

	wg.Wait()
}

// 2.2 Флажок
func Test6_4(t *testing.T) {
	canWork := true

	go func() {
		for canWork {
			fmt.Println("Working...")
			time.Sleep(time.Duration(200) * time.Millisecond)
		}
	}()

	time.Sleep(time.Second)
	canWork = false
}
