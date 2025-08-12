package utils

import (
	"fmt"
	"sync"
)

func Pipeline(input <-chan uint8, output chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(output)

	for num := range input {
		result := float64(num) * float64(num) * float64(num)
		output <- result
	}
}

func GenerateNumbers(input chan<- uint8, numbers []uint8, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(input)

	for _, num := range numbers {
		input <- num
	}
}

func ReadResults(output <-chan float64, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Pipeline results:")
	for result := range output {
		fmt.Printf("Result: %.2f\n", result)
	}
}
