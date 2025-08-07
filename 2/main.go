package main

import (
	"fmt"
	"math/rand"
	"task2/utils"
)

func main() {
	originalSlice := make([]int, 10)
	for index := range originalSlice {
		originalSlice[index] = rand.Int()
	}
	fmt.Printf("SliceExample1 (memory efficient): originalSlice: %v\n", originalSlice)
	evenNumSlice1 := utils.SliceExample1(originalSlice, true)
	fmt.Printf("SliceExample1 (memory efficient): evenNumSlice1: %v\n\n\n", evenNumSlice1)

	fmt.Printf("SliceExample2 (CPU efficient): originalSlice: %v\n", originalSlice)
	evenNumSlice2 := utils.SliceExample2(originalSlice, true)
	fmt.Printf("SliceExample2 (CPU efficient): evenNumSlice2: %v\n\n\n", evenNumSlice2)

	fmt.Printf("AddElements: originalSlice: %v\n", originalSlice)
	newSlice := utils.AddElements(originalSlice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("AddElements: newSlice: %v\n\n\n", newSlice)

	fmt.Printf("CopySlice: originalSlice: %v\n", originalSlice)
	copiedSlice := utils.CopySlice(originalSlice)
	fmt.Printf("CopySlice: copiedSlice: %v\n\n\n", copiedSlice)

	originalSlice[0] = 1000000000
	originalSlice[1] = 1000000000
	originalSlice[2] = 1000000000
	fmt.Printf("CopySlice: originalSlice: %v\n", originalSlice)
	fmt.Printf("CopySlice: copiedSlice: %v\n\n\n", copiedSlice)
}
