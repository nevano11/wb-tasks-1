package main

import (
	"fmt"
	"testing"
)

// Удалить i-ый элемент из слайса.

func removeItemByIndex(slice []int, index int) []int {
	// Защита
	if index < 0 || index >= len(slice) {
		return slice
	}

	// Копирование элементов в новый слайс. Старый слайс остается прежним и доступным
	res := make([]int, len(slice)-1)
	for i := 0; i < index; i++ {
		res[i] = slice[i]
	}
	for i := index + 1; i < len(slice); i++ {
		res[i-1] = slice[i]
	}
	return res
}

func Test23(t *testing.T) {
	baseSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(removeItemByIndex(baseSlice, 2))
	fmt.Println(removeItemByIndex(baseSlice, 1))
	fmt.Println(removeItemByIndex(baseSlice, 4))
	fmt.Println(removeItemByIndex(baseSlice, 1000))
	fmt.Println(removeItemByIndex(baseSlice, -1))
}
