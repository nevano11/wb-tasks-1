package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Поменять местами два числа без создания временной переменной.

func Test13(t *testing.T) {
	// Способ 1 - средствами языка Go
	a, b := 1, 2
	a, b = b, a
	assert.Equal(t, 2, a)
	assert.Equal(t, 1, b)

	// Способ 2, с помощью операций + -
	// Опасно на больших числах (math.MaxInt, math.MaxInt)
	c, d := 1, 2
	c = c + d // 1 + 2 = 3  - c
	d = c - d // 3 - 2 = 1  - d
	c = c - d // 3 - 1 = 2  - c
	assert.Equal(t, 2, c)
	assert.Equal(t, 1, d)
}
