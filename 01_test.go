package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от
// родительской структуры Human (аналог наследования).

// Структура Human - метод кушать
type Human struct{}

func (h *Human) Eat() string {
	return "eating"
}

// Встраивание Human в Action
type Action struct {
	Human
}

func Test1(t *testing.T) {
	act := &Action{}
	hum := &Human{}

	hEatResult := hum.Eat()
	aEatResult := act.Eat()
	fmt.Printf("human  EatResult: %s\n", hEatResult)
	fmt.Printf("action EatResult: %s\n", aEatResult)
	assert.Equal(t, hum.Eat(), act.Eat())
}
