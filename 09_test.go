package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

// Разработать конвейер чисел.
//
// Даны два канала:
// в первый пишутся числа (x) из массива,
// во второй — результат операции x*2,
// данные из второго канала должны выводиться в stdout

func Test9(t *testing.T) {
	sourceSize := 100
	source := make([]int, sourceSize, sourceSize)
	// Рандомно заполняем
	fillSource(source, sourceSize)
	fmt.Println("Source: ", source)

	wg := sync.WaitGroup{}
	wg.Add(2)

	ch1 := make(chan int, sourceSize)
	ch2 := make(chan int, sourceSize)

	// Записать числа массива в канал 1
	go func() {
		defer wg.Done()
		defer close(ch1)
		for _, v := range source {
			ch1 <- v
		}
	}()

	// Перенести данные из канала 1 в канал 2
	go func() {
		defer wg.Done()
		defer close(ch2)
		for val := range ch1 {
			fmt.Printf("Transferred: val=%d\n", val)
			ch2 <- val * 2
		}
	}()

	// Прочитать данные из канала 2 и вывести
	for val := range ch2 {
		fmt.Printf("Final: val=%d\n", val)
	}
	wg.Wait()
}

// Заполнить рандомными данными
func fillSource(source []int, sourceSize int) {
	for i := 0; i < sourceSize; i++ {
		source[i] = rand.Intn(10000*2+1) - 10000
	}
}
