package main

import (
	"fmt"
	"sync"
	"testing"
)

// Написать программу, которая конкурентно рассчитает значение квадратов
// чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func Test2(t *testing.T) {
	arr := [5]int{2, 4, 6, 8, 10}

	// Исходный массив заливается в канал
	intChannel := make(chan int, len(arr))
	for i := 0; i < len(arr); i++ {
		intChannel <- arr[i]
	}
	close(intChannel)

	// Создается len(arr) горутин, вычисляющих квадраты чисел
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for i := 0; i < len(arr); i++ {
		go func() {
			defer wg.Done()
			for v := range intChannel {
				fmt.Printf("%d\n", v*v)
			}
		}()
	}
	wg.Wait()
}
