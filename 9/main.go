package main

import (
	"fmt"
	"sync"
	"task9/utils"
)

func main() {
	testNumbers := []uint8{1, 2, 3, 4, 5}

	input := make(chan uint8, len(testNumbers))
	output := make(chan float64, len(testNumbers))

	var wg sync.WaitGroup

	wg.Add(1)
	go utils.Pipeline(input, output, &wg)

	wg.Add(1)
	go utils.ReadResults(output, &wg)

	wg.Add(1)
	go utils.GenerateNumbers(input, testNumbers, &wg)

	wg.Wait()

	fmt.Println("Pipeline processing completed!")
}
