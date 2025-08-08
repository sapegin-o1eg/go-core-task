package main

import "fmt"

func main() {
	fmt.Printf("Hi!\n")

	slice1 := []string{"apple", "banana", "banana", "cherry", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	fmt.Println(Difference(slice1, slice2))
}

func Difference(source, exclude []string) []string {
	excludeSet := make(map[string]bool)
	resultSlice := make([]string, 0)

	for _, item := range exclude {
		excludeSet[item] = true
	}

	for _, item := range source {
		if _, ok := excludeSet[item]; !ok {
			resultSlice = append(resultSlice, item)
		}
	}

	return resultSlice
}
