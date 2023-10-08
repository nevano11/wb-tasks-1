package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Разработать программу, которая в рантайме способна определить
// тип переменной: int, string, bool, channel из переменной типа interface{}.

func getType(obj interface{}) reflect.Type {
	return reflect.TypeOf(obj)
}

func Test14(t *testing.T) {
	a := 1
	b := "asd"
	c := true
	d := make(chan int)

	fmt.Println(getType(a))
	fmt.Println(getType(b))
	fmt.Println(getType(c))
	fmt.Println(getType(d))
}
