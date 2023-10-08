package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// Разработать программу, которая переворачивает подаваемую
// на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

func reverseString(input string) string {
	// for reverse + string builder
	sb := strings.Builder{}
	inputRunes := []rune(input)
	for i := len(inputRunes) - 1; i >= 0; i-- {
		sb.WriteRune(inputRunes[i])
	}
	return sb.String()
}

func Test19(t *testing.T) {
	input := "главрыба"
	assert.Equal(t, "абырвалг", reverseString(input))
}
