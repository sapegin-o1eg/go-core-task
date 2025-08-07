package utils

import (
	"fmt"
	"slices"
)

func SliceExample1(s []int, verbose ...bool) []int {
	// lets make this func memory efficient
	// but drawback of this is that we need to iterate over the slice twice
	isVerbose := false
	if len(verbose) > 0 {
		isVerbose = verbose[0]
	}
	evenNumCount := 0
	for _, value := range s {
		if value%2 == 0 {
			evenNumCount++
		}
	}
	evenNumSlice := make([]int, evenNumCount)

	evenNumSliceIndex := 0
	for index, value := range s {
		if value%2 == 0 {
			evenNumSlice[evenNumSliceIndex] = value
			evenNumSliceIndex++
			if isVerbose {
				fmt.Printf("Found even number at index %v : %v\n", index, value)
			}
		}
	}
	return evenNumSlice
}

func SliceExample2(s []int, verbose ...bool) []int {
	// lets make this func CPU efficient
	// but drawback of this is that new slice capacity will be equal to the length of original slice
	// and RAM will be preallocated for the length of original slice
	isVerbose := false
	evenNumSliceIndex := 0

	if len(verbose) > 0 {
		isVerbose = verbose[0]
	}

	evenNumSlice := make([]int, 0, len(s))

	for index, value := range s {
		if value%2 == 0 {
			evenNumSlice = append(evenNumSlice, value)
			evenNumSliceIndex++
			if isVerbose {
				fmt.Printf("Found even number at index %v : %v\n", index, value)
			}
		}
	}
	return evenNumSlice
}

func AddElements(s []int, values ...int) []int {
	newSlice := make([]int, len(s), len(s)+len(values))
	copy(newSlice, s)
	for _, value := range values {
		newSlice = append(newSlice, value)
	}
	return newSlice
}

func CopySlice(s []int) []int {
	newSlice := make([]int, len(s))
	copy(newSlice, s)
	return newSlice
}

func RemoveElements(s []int, indexesForRemoval ...int) []int {
	capacity := len(s) - len(indexesForRemoval)
	if capacity < 0 {
		capacity = 0
	}
	newSlice := make([]int, 0, capacity)
	for index, value := range s {
		if !slices.Contains(indexesForRemoval, index) {
			newSlice = append(newSlice, value)
		}
	}
	return newSlice
}
