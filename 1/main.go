package main

import (
	"fmt"
	"task1/utils"
)

func main() {
	someIntegerNumbers := utils.PrintSomeIntegerNumbers()
	fmt.Println(someIntegerNumbers)
	utils.PrintSomeFloat64Numbers()
	utils.PrintSomeStrings()
	utils.PrintSomeBooleans()
	utils.PrintSomeComplex64Numbers()
}
