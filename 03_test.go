package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func Test3(t *testing.T) {
	sum := 0
	mut := sync.Mutex{}
	arr := []int{2, 4, 6, 8, 10}

	// Исходные числа загоняются в массив
	channel := make(chan int, len(arr))
	for i := 0; i < len(arr); i++ {
		channel <- arr[i]
	}
	close(channel)

	// Для предотвращения race condition блокируем мьютексом
	appendCounter := func(a int) {
		mut.Lock()
		defer mut.Unlock()
		sum += a
	}

	// Параллельное вычисление квадратов чисел
	wg := sync.WaitGroup{}
	wg.Add(len(arr))
	for i := 0; i < len(arr); i++ {
		go func() {
			defer wg.Done()
			for i2 := range channel {
				appendCounter(i2 * i2)
			}
		}()
	}
	wg.Wait()

	assert.Equal(t, 220, sum)
	fmt.Println(sum)
}
