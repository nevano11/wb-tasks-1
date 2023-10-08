package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func generateRandomSlice(length int) []int {
	resultSlice := make([]int, length, length)
	for i := 0; i < length; i++ {
		resultSlice[i] = rand.Intn(10001) - 5000
	}
	return resultSlice
}

// Оболочка, которая сама выставит корректные low и high
func myQuickSort[T comparable](arr []T, cmpFunc func(item1, item2 T) bool) {
	length := len(arr)
	qsortbase(arr, 0, length-1, cmpFunc)
}

// Вызывается рекурсивно для подмассивов
func qsortbase[T comparable](arr []T, low, high int, cmpFunc func(item1, item2 T) bool) {
	if high <= low {
		// array part has 1 element
		return
	}
	pivot := arr[high]
	wall := low - 1

	for index := low; index < high; index++ {
		// Число pivot больше числа[index]
		if cmpFunc(pivot, arr[index]) {
			wall++
			arr[index], arr[wall] = arr[wall], arr[index]
		}
	}

	wall++
	arr[high], arr[wall] = arr[wall], pivot
	qsortbase(arr, low, wall-1, cmpFunc)
	qsortbase(arr, wall+1, high, cmpFunc)
}

func Test16(t *testing.T) {
	sliceToSorting := generateRandomSlice(100)
	//sliceToSorting := make([]int, 0, 0)
	fmt.Println(sliceToSorting)

	// Выполняем сортировку
	myQuickSort(sliceToSorting, func(item1, item2 int) bool {
		return item1 > item2
	})
	fmt.Println(sliceToSorting)

	// Проверка на отсортированность с sort.SliceIsSorted
	// (тут используется функция сравнения Less => имеет слегка иной вид)
	assert.True(t, sort.SliceIsSorted(sliceToSorting, func(i, j int) bool {
		return sliceToSorting[j] > sliceToSorting[i]
	}))
}

//[5, 1, 4, 5, 6, 3, 2, 1, 5, 3]
// low = 0, high = 9
// low < high ? yes
// pivot = 3, wall = -1
// [5, 1, 4, 5, 6, 3, 2, 1, 5] 3
// [5, 1, 4, 5, 6, 3, 2, 1, 5] 3
// cmp 5 - nothing
// cmp 1:
// [ | 5, 1, 4, 5, 6, 3, 2, 1, 5] 3
// [1 | 5, 4, 5, 6, 3, 2, 1, 5] swap 1 and wall + 1; wall++ // wall = 0
// cmp 4 - nothing
// cmp 5 - nothing
// cmp 6 - nothing
// cmp 3
// [1 | 5, 4, 5, 6, 3, 2, 1, 5] 3
// [1, 3 | 4, 5, 6, 5, 2, 1, 5] swap 3 and wall + 1; wall++ // wall = 1
// cmp 2
// [1, 3 | 4, 5, 6, 5, 2, 1, 5] 3
// [1, 3, 2 | 5, 6, 5, 4, 1, 5] swap 2 and wall + 1; wall++ // wall = 2
// cmp 1
// [1, 3, 2 | 5, 6, 5, 4, 1, 5] 3
// [1, 3, 2, 1 | 6, 5, 4, 5, 5] swap 1 and wall + 1; wall++ // wall = 3
// cmp 5 - nothing
// at the end :
// swap start 3 with wall+1 // wall = 3
// [1, 3, 2, 1, | 3, 5, 4, 5, 5, 6]
// qsort(low = 0, high = wall)
// qsort(low = wall+1, high = high)

// alg gives O(n^2):
// [1, 2, 3, 4, 5] // pivot 5
// [1, 2, 3, 4] // pivot 4
// [1, 2, 3] // pivot 3
// [1, 2] // pivot 2
// [1] // pivot 1

// Or [5, 4, 3, 2, 1]
// [5, 4, 3, 2 | ]
// ...
