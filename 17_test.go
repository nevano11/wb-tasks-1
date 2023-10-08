package main

import (
	"fmt"
	"testing"
)

//Реализовать бинарный поиск встроенными методами языка.

// Оболочка, которая сама выставит корректные low и high
// Пришлось использовать generic на ограниченный набор
// тк comparable не дает > и <
func myBinarySearch[T int | float64 | float32](arr []T, item T) int {
	return binarySearchInner(arr, item, 0, len(arr)-1)
}

func binarySearchInner[T int | float64 | float32](arr []T, item T, low, high int) int {
	if high <= low {
		// подмассив состоит из 1 элемента
		if arr[low] == item {
			return low
		}
		// Если этот элемент != item, то -1 - не найден
		return -1
	}

	// Делим пополам подмассив. Попали - возвращаем индекс
	// Иначе идем в ту часть, где ожидается поиск элементов
	midIndex := (high + low) / 2
	if arr[midIndex] == item {
		return midIndex
	} else if arr[midIndex] > item {
		return binarySearchInner(arr, item, low, midIndex-1)
	}
	return binarySearchInner(arr, item, midIndex+1, high)
}

func Test17(t *testing.T) {
	baseArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// array must be sorted asc

	for i := 0; i < len(baseArr); i++ {
		fmt.Println(myBinarySearch(baseArr, baseArr[i]))
	}
	fmt.Println(myBinarySearch(baseArr, 0))
	fmt.Println(myBinarySearch(baseArr, -1))
	fmt.Println(myBinarySearch(baseArr, 10))
}
