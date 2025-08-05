package utils

import (
	"fmt"
)

func PrintSomeIntegerNumbers() []int {
	var firstNumIntInDecimalNotation int = 1
	var secondNumIntInDecimalNotation int = 2
	var thirdNumIntInDecimalNotation int = 3

	return []int{
		firstNumIntInDecimalNotation,
		secondNumIntInDecimalNotation,
		thirdNumIntInDecimalNotation,
	}
}

func PrintSomeFloat64Numbers() {
	// let's use another declaration style
	firstNumFloat64 := 1.01
	secondNumFloat64 := 2.02
	thirdNumFloat64 := 3.03

	fmt.Println(firstNumFloat64)
	fmt.Println(secondNumFloat64)
	fmt.Println(thirdNumFloat64)
}

func PrintSomeStrings() {
	firstString := "It's the first string"
	secondString := "It's the second string"

	fmt.Println(firstString)
	fmt.Println(secondString)
}

func PrintSomeBooleans() {
	firstBoolean := true
	secondBoolean := false

	fmt.Println(firstBoolean)
	fmt.Println(secondBoolean)
}

func PrintSomeComplex64Numbers() {
	firstComplex64 := 1 + 1i
	secondComplex64 := 2 + 2i

	fmt.Println(firstComplex64)
	fmt.Println(secondComplex64)
}
