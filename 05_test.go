package main

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func Test5(t *testing.T) {
	// Количество сообщений,
	seconds := 2
	messages := make(chan int, 1)
	mut := sync.RWMutex{}
	stopper := false

	// Методы для дальнейшей остановки записи в канал
	canSend := func() bool {
		mut.RLock()
		defer mut.RUnlock()
		return !stopper
	}
	stop := func() {
		mut.Lock()
		defer mut.Unlock()
		stopper = true
	}

	// reader - постоянно пытается читать из канала
	go func() {
		for message := range messages {
			fmt.Printf("Readed message %d\n", message)
		}
	}()

	// writer - если можно писать - пишу, нельзя - закрываю канал
	go func() {
		for {
			if canSend() {
				message := rand.Intn(999)
				fmt.Printf("Write message %d\n", message)
				messages <- message
			} else {
				close(messages)
				break
			}
		}
	}()

	// Основной поток ждет установленное время
	select {
	case <-time.After(time.Duration(seconds) * time.Second):
		stop()
		fmt.Printf("Time is out\n")
	}
}
