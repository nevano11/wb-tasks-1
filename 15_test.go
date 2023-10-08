package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// К каким негативным последствиям может привести данный фрагмент кода,
// и как это исправить? Приведите корректный пример реализации.

//var justString string
//
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}

func getHugeStringLess100(int) string {
	return ""
}

func getHugeString(int) string {
	return "MВCaPRCDm0sUqeP2i4EcWafUEyBrvmwXHJguCq7MВCaPRCDm0sUqeP2i4EcWafUEyBrvmwXHJguCq7MВCaPRCDm0sUqeP2i4EcWafUEy"
}

func someFunc(createHugeString func(int) string) (justString string) {
	v := createHugeString(1 << 10)
	justString = v[:100]
	return
}

func someFuncGood(createHugeString func(int) string) string {
	v := createHugeString(1 << 10)
	// Защита от паники по поводу недостаточной длины страки
	if len(v) > 100 {
		//во избежание утечки памяти
		return strings.Clone(v[:100])
	}
	return "" // or add error on signature
}

func Test15(t *testing.T) {
	// 1 - глобальная переменная плохо, к ней все имеют доступ все функции,
	// бизнес-логика становится непрозрачной
	fmt.Println("1 - глобальная переменная плохо. Бизнес-логика становится непрозрачной")

	// 2 - createHugeString точно вернет строку длинной больше 100?
	// Получаем панику
	fmt.Println("2 - createHugeString точно вернет строку длинной больше 100?. Паника")
	assert.Panics(t, func() {
		someFunc(getHugeStringLess100)
	})

	// 3 - утечка памяти. Слайс полученный операцией [1:100] продолжает тянуть его за собой повсеместно
	fmt.Println("3 - утечка памяти")

	big := getHugeString(1)
	_ = someFunc(func(i int) string {
		return big
	})

	// Корректный пример реализации
	res := someFuncGood(getHugeStringLess100)
	res2 := someFuncGood(getHugeString)
	fmt.Printf("Good func result: %s, %s\n", res, res2)
}
