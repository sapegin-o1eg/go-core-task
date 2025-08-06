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
	fmt.Printf("SliceExample1 (memory efficient): evenNumSlice1: %v\n", evenNumSlice1)

	fmt.Printf("SliceExample2 (CPU efficient): originalSlice: %v\n", originalSlice)
	evenNumSlice2 := utils.SliceExample2(originalSlice, true)
	fmt.Printf("SliceExample2 (CPU efficient): evenNumSlice2: %v\n", evenNumSlice2)
}
