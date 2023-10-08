package main

import (
	"fmt"
	"testing"
)

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func Test12(t *testing.T) {
	objs := []string{"cat", "cat", "dog", "cat", "tree"}
	var set []string
	setController := make(map[string]struct{})

	// С помощью мапы отслеживаем уникальность элементов
	for _, v := range objs {
		if _, exists := setController[v]; !exists {
			setController[v] = struct{}{}
			set = append(set, v)
		}
	}

	fmt.Println(set)
}
