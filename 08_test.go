package main

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

// Функция, устанавливающая бит в 1 или 0
func changeBit(src int64, bitNum int, bitVal rune) (int64, error) {
	// Защита от дурака
	if bitVal < 0 || bitVal > 1 {
		return 0, errors.New("Wrong attributes")
	}

	if bitVal == rune(1) {
		src |= 1 << bitNum
		return src, nil
	}
	src &= ^(1 << bitNum)
	return src, nil
}

func Test8(t *testing.T) {
	var num int64 = 15412

	// Изначально
	fmt.Println("Before")
	fmt.Println(num)
	fmt.Println(strconv.FormatInt(num, 2))

	// Ставит 1 на 7 место
	bitVal := rune(1)
	bitNum := 7
	fmt.Printf("\nSet %d on %d position\n", bitVal, bitNum)

	newNum, _ := changeBit(num, bitNum, bitVal)
	fmt.Println(newNum)
	fmt.Println(strconv.FormatInt(newNum, 2))

	// Ставит 0 на 2 место
	bitVal = rune(0)
	bitNum = 2
	fmt.Printf("\nSet %d on %d position\n", bitVal, bitNum)

	newNum2, _ := changeBit(newNum, bitNum, bitVal)
	fmt.Println(newNum2)
	fmt.Println(strconv.FormatInt(newNum2, 2))
}
