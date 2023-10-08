package main

import (
	"testing"
	"time"
)

// Реализовать собственную функцию sleep

func mySleep(d time.Duration) {
	select {
	case <-time.After(d):
		return
	}
}

func Test25(t *testing.T) {
	mySleep(time.Duration(2) * time.Second)
}
