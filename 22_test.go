package main

import (
	"fmt"
	"math/big"
	"testing"
)

// Разработать программу, которая
// перемножает, делит, складывает, вычитает две числовых переменных
// a,b, значение которых > 2^20.

func Test22(t *testing.T) {
	a := big.NewInt(522_211_122_556)
	b := big.NewInt(174_070_374_186)

	add := big.NewInt(0) // 696,281,496,742
	sub := big.NewInt(0) // 348,140,748,370
	div := big.NewInt(0) // 2.99999999998851039409
	mod := big.NewInt(0) // 174,070,374,184
	mul := big.NewInt(0) // 90,901,485,507,414,024,739,416

	var div2, _ = div.DivMod(a, b, mod)
	fmt.Printf("ints: %d %d\n+: %d,\n-: %d,\n/, %%: %d, %d,\n*: %d\n", a, b, add.Add(a, b), sub.Sub(a, b), div2, mod, mul.Mul(a, b))
}
