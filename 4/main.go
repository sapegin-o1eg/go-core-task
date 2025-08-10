package main

import (
	"fmt"
	"task4/utils"
)

func main() {
	slice1 := []string{"apple", "banana", "banana", "cherry", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "banana", "date", "fig"}

	fmt.Printf("Slice1: %+v\n", slice1)
	fmt.Printf("Slice2: %+v\n\n\n", slice2)

	fmt.Printf("utils.Difference: %+v\n", utils.Difference(slice1, slice2))
}
