package main

import (
	"fmt"
	"task1/utils"
)

func main() {
	// assigne vars and print them out
	someIntegerNumbers := utils.GetSomeIntegerNumbers()
	fmt.Printf("Got integer numbers:\n  - %v\n", someIntegerNumbers)

	firstNumFloat64, secondNumFloat64, thirdNumFloat64 := utils.GetSomeFloat64Numbers()
	fmt.Printf("Got float64 numbers:\n  - %v\n  - %v\n  - %v\n", firstNumFloat64, secondNumFloat64, thirdNumFloat64)

	someStringOne, someStringTwo := utils.GetSomeStrings()
	fmt.Printf("Got strings:\n  - %v\n  - %v\n", someStringOne, someStringTwo)

	someBooleanOne, someBooleanTwo := utils.GetSomeBooleans()
	fmt.Printf("Got boleans:\n  - %v\n  - %v\n", someBooleanOne, someBooleanTwo)

	someComplex64Numbers := utils.GetSomeComplex64Numbers()
	fmt.Printf("Got complex64 numbers:\n  - %v\n  - %v\n", someComplex64Numbers.First, someComplex64Numbers.Second)

	// get types of vars
	utils.GetVarType(someIntegerNumbers)

	someIntegerNumbersForVariadicUnpacking := make([]interface{}, len(someIntegerNumbers))
	for i, v := range someIntegerNumbers {
		someIntegerNumbersForVariadicUnpacking[i] = v
	}
	utils.GetVarType(someIntegerNumbersForVariadicUnpacking...)
}