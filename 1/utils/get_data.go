package utils

import (
	"task1/models"
)

func GetSomeIntegerNumbers() []int {
	var firstNumIntInDecimalNotation = 1
	var secondNumIntInDecimalNotation = 2
	var thirdNumIntInDecimalNotation = 3
	var firstNumIntInOctalNotation = 0o4
	var secondNumIntIOctalNotation = 0o5
	var thirdNumIntInOctalNotation = 0o6
	var firstNumIntInHexNotation = 0x7
	var secondNumIntIHexNotation = 0x8
	var thirdNumIntInHexNotation = 0x9

	return []int{
		firstNumIntInDecimalNotation,
		secondNumIntInDecimalNotation,
		thirdNumIntInDecimalNotation,
		firstNumIntInOctalNotation,
		secondNumIntIOctalNotation,
		thirdNumIntInOctalNotation,
		firstNumIntInHexNotation,
		secondNumIntIHexNotation,
		thirdNumIntInHexNotation,
	}
}

func GetSomeFloat64Numbers() (float64, float64, float64) {
	// let's use another declaration style
	// and return the values as a "tuple"
	firstNumFloat64 := 1.01
	secondNumFloat64 := 2.02
	thirdNumFloat64 := 3.03

	return firstNumFloat64, secondNumFloat64, thirdNumFloat64
}

func GetSomeStrings() (FirstString, SecondString string) {
	FirstString = "It's the first string"
	SecondString = "It's the second string"

	// naked return
	return
}

func GetSomeBooleans() (FirstBoolean, SecondBoolean bool) {
	FirstBoolean = true
	SecondBoolean = false

	// explicit return
	return FirstBoolean, SecondBoolean
}

func GetSomeComplex64Numbers() models.Complex64Numbers {
	firstComplex64 := complex64(1 + 1i)
	secondComplex64 := complex64(2 + 2i)

	return models.Complex64Numbers{
		First:  firstComplex64,
		Second: secondComplex64,
	}
}
