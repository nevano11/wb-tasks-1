package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func Test20(t *testing.T) {
	input := "snow dog sun"
	splitted := strings.Fields(input)

	// for reverse + string builder
	sb := strings.Builder{}
	for i := len(splitted) - 1; i >= 0; i-- {
		sb.WriteString(splitted[i] + " ")
	}
	res := strings.TrimRight(sb.String(), " ")

	assert.Equal(t, "sun dog snow", res)
}
