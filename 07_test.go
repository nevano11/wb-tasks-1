package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

// Реализовать конкурентную запись данных в map

func Test7(t *testing.T) {
	dict := make(map[int]int)
	mut := sync.Mutex{}
	wg := new(sync.WaitGroup)

	// Метод записи в map. Конкурентность регулируется с помощью mutex
	writeOnMap := func(k, v int) {
		mut.Lock()
		defer mut.Unlock()
		dict[k] = v
	}

	// Запуск конкурентной записи
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			val := rand.Intn(100) + 1
			writeOnMap(val, val*val)
		}()
	}

	wg.Wait()
	fmt.Println(dict)
}
