package main

import (
	"fmt"
	"testing"
)

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножноствах не важна.

//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func nearestInt(num float32) int {
	helper := int(num) / 10
	if num < 0 {
		helper -= 1
	}
	return helper * 10
}

func Test10(t *testing.T) {
	tempStatistics := []float32{24.5, -21.0, 32.5, -25.4, -27.0, 13.0, 19.0, 15.5, 20, -1, 1}
	groupedData := make(map[int][]float32)

	// Для каждого числа берем близжайшее число слева, кратное 10
	for _, v := range tempStatistics {
		groupedData[nearestInt(v)] = append(groupedData[nearestInt(v)], v)
	}

	fmt.Println(groupedData)
}
