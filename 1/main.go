package main

import (
	"fmt"
	"task1/utils"
)

func main() {
	// assigne vars and print them out
	fmt.Printf("#  Print assigned variables  =======================================>\n")
	someIntegerNumbers := utils.GetSomeIntegerNumbers()
	fmt.Printf("Got array of integer numbers:\n  - %v\n", someIntegerNumbers)

	firstNumFloat64, secondNumFloat64, thirdNumFloat64 := utils.GetSomeFloat64Numbers()
	fmt.Printf("Got float64 numbers:\n  - %v\n  - %v\n  - %v\n", firstNumFloat64, secondNumFloat64, thirdNumFloat64)

	someStringOne, someStringTwo := utils.GetSomeStrings()
	fmt.Printf("Got strings:\n  - %v\n  - %v\n", someStringOne, someStringTwo)

	someBooleanOne, someBooleanTwo := utils.GetSomeBooleans()
	fmt.Printf("Got boleans:\n  - %v\n  - %v\n", someBooleanOne, someBooleanTwo)

	someComplex64Numbers := utils.GetSomeComplex64Numbers()
	fmt.Printf("Got complex64 numbers:\n  - %v\n  - %v\n", someComplex64Numbers.First, someComplex64Numbers.Second)
	fmt.Printf("#  Print assigned variables  =======================================<\n\n")

	// get types of vars
	fmt.Printf("#  Print variables' types   ========================================>\n")
	utils.GetVarType(someIntegerNumbers)
	someIntegerNumbersForVariadicUnpacking := make([]interface{}, len(someIntegerNumbers))
	for i, v := range someIntegerNumbers {
		someIntegerNumbersForVariadicUnpacking[i] = v
	}
	utils.GetVarType(someIntegerNumbersForVariadicUnpacking...)

	utils.GetVarType(firstNumFloat64, secondNumFloat64, thirdNumFloat64)

	utils.GetVarType(someStringOne, someStringTwo)

	utils.GetVarType(someBooleanOne, someBooleanTwo)

	utils.GetVarType(someComplex64Numbers.First, someComplex64Numbers.Second)

	// convert to string and cat
	concatenatedString := utils.ConcatenateStrings(
		utils.GetTypeAsString(someIntegerNumbersForVariadicUnpacking[0]),
		utils.GetTypeAsString(someIntegerNumbersForVariadicUnpacking[1]),
		utils.GetTypeAsString(someIntegerNumbersForVariadicUnpacking[2]),
		utils.GetTypeAsString(firstNumFloat64),
		utils.GetTypeAsString(secondNumFloat64),
		utils.GetTypeAsString(thirdNumFloat64),
		utils.GetTypeAsString(someStringOne),
		utils.GetTypeAsString(someStringTwo),
		utils.GetTypeAsString(someBooleanOne),
		utils.GetTypeAsString(someBooleanTwo),
		utils.GetTypeAsString(someComplex64Numbers.First),
		utils.GetTypeAsString(someComplex64Numbers.Second),
	)
	fmt.Printf("Concatenated string is: %s\n", concatenatedString)
	fmt.Printf("#  Print variables' types  =========================================<\n\n")
	fmt.Printf("#  Print concatenated string as runes  =============================>\n")
	fmt.Printf("Here goes concatenated string as rune slice:\n  - %q\n", utils.StringToRunes(concatenatedString))
	fmt.Printf("#  Print concatenated string as runes  =============================<\n\n")
}
