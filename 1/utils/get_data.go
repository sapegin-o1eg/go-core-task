package utils

import (
	"task1/models"
)

func GetSomeIntegerNumbers() []int {
	var firstNumIntInDecimalNotation int = 1
	var secondNumIntInDecimalNotation int = 2
	var thirdNumIntInDecimalNotation int = 3

	return []int{
		firstNumIntInDecimalNotation,
		secondNumIntInDecimalNotation,
		thirdNumIntInDecimalNotation,
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

	// explicit retuen
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
