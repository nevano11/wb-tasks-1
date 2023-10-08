package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
//
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func hasUniqueLetters(input string) bool {
	runes := []rune(input)
	validator := make(map[rune]struct{})

	// Пользуемя map для контраля за уникальностью элементов

	for _, r := range runes {
		if _, exists := validator[r]; exists {
			return false
		}
		validator[r] = struct{}{}
	}
	return true
}

func Test26(t *testing.T) {
	assert.Equal(t, true, hasUniqueLetters("abcd"))
	assert.Equal(t, false, hasUniqueLetters("abCdefAaf"))
	assert.Equal(t, false, hasUniqueLetters("aabcd"))
}
