package main

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sync"
	"testing"
)

// Реализовать структуру-счетчик, которая будет
// инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

// Конкуррентный счетчик - защищает свой внутренний счетчик мьютексом
type ConcurrentCounter struct {
	counter int
	mut     sync.Mutex
}

func NewConcurrentCounter(initialValue int) *ConcurrentCounter {
	return &ConcurrentCounter{
		counter: initialValue,
		mut:     sync.Mutex{},
	}
}
func (cc *ConcurrentCounter) Increment() {
	cc.mut.Lock()
	defer cc.mut.Unlock()
	cc.counter++
}
func (cc *ConcurrentCounter) Counter() int {
	cc.mut.Lock()
	defer cc.mut.Unlock()
	return cc.counter
}

func Test18(t *testing.T) {
	goriutineCount := rand.Intn(90) + 10
	wg := sync.WaitGroup{}
	wg.Add(goriutineCount)

	counter := NewConcurrentCounter(0)

	// Рандомное количество горутин инкрементирует счетчик
	for i := 0; i < goriutineCount; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()
	assert.Equal(t, goriutineCount, counter.Counter())
}
