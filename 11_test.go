package main

import (
	"fmt"
	"testing"
)

// Реализовать пересечение двух неупорядоченных множеств

func Test11(t *testing.T) {
	set1 := []int{1, 2, 3, 4, 5, 6}
	set2 := []int{10, 2, 4, 1, 49, 20, 7, 8}

	// Изначально имеем дело с множеством set1
	// Запоминаем, какие числа уже встречались
	stats := make(map[int]struct{})
	for _, v := range set1 {
		stats[v] = struct{}{}
	}

	// Добавляем в 1 множество если ранее не встречалось
	for _, v := range set2 {
		// Поиск по мапе шустрый, тк под капотом хеш-функция
		if _, exists := stats[v]; !exists {
			stats[v] = struct{}{}
			set1 = append(set1, v)
		}
	}

	fmt.Println(set1)
}
