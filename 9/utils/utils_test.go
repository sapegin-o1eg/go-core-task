package utils

import (
	"sync"
	"testing"
)

func TestPipeline(t *testing.T) {
	// Test data
	testNumbers := []uint8{1, 2, 3}
	expectedResults := []float64{1.0, 8.0, 27.0}

	// Create channels
	input := make(chan uint8, len(testNumbers))
	output := make(chan float64, len(testNumbers))

	// Create wait group
	var wg sync.WaitGroup

	// Start pipeline in goroutine
	wg.Add(1)
	go Pipeline(input, output, &wg)

	// Send test numbers in main thread (sequential)
	for _, num := range testNumbers {
		input <- num
	}
	close(input) // Close input after sending all data

	// Wait for pipeline to complete
	wg.Wait()

	// Collect results from output channel
	var results []float64
	for result := range output {
		results = append(results, result)
	}

	// Verify results
	if len(results) != len(expectedResults) {
		t.Errorf("Expected %d results, got %d", len(expectedResults), len(results))
	}

	for i, result := range results {
		if result != expectedResults[i] {
			t.Errorf("Expected result[%d] = %.1f, got %.1f", i, expectedResults[i], result)
		}
	}
}

func TestGenerateNumbers(t *testing.T) {
	testNumbers := []uint8{5, 10, 15}
	input := make(chan uint8, len(testNumbers))

	var wg sync.WaitGroup
	wg.Add(1)

	// Start generator in goroutine
	go GenerateNumbers(input, testNumbers, &wg)

	// Collect generated numbers in main thread (sequential)
	var received []uint8
	for num := range input {
		received = append(received, num)
	}

	wg.Wait()

	// Verify all numbers were generated
	if len(received) != len(testNumbers) {
		t.Errorf("Expected %d numbers, got %d", len(testNumbers), len(received))
	}

	for i, num := range received {
		if num != testNumbers[i] {
			t.Errorf("Expected number[%d] = %d, got %d", i, testNumbers[i], num)
		}
	}
}

func TestReadResults(t *testing.T) {
	// This test verifies that ReadResults can process results from a channel
	// Since it's mainly for output, we'll test the channel handling

	output := make(chan float64, 3)
	testResults := []float64{1.0, 8.0, 27.0}

	var wg sync.WaitGroup
	wg.Add(1)

	// Start reader in goroutine
	go ReadResults(output, &wg)

	// Send test results in main thread (sequential)
	for _, result := range testResults {
		output <- result
	}
	close(output) // Close output after sending all data

	wg.Wait()
	// Test passes if no panic occurs during processing
}

func TestPipelineWithZero(t *testing.T) {
	// Test edge case with zero
	expectedResult := 0.0

	input := make(chan uint8, 1)
	output := make(chan float64, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	go Pipeline(input, output, &wg)

	// Send zero
	go func() {
		input <- 0
		close(input) // Close input AFTER sending data
	}()

	// Get result
	result := <-output
	wg.Wait()

	if result != expectedResult {
		t.Errorf("Expected result for 0 = %.1f, got %.1f", expectedResult, result)
	}
}

func TestPipelineWithMaxUint8(t *testing.T) {
	// Test edge case with maximum uint8 value
	maxUint8 := uint8(255)
	expectedResult := float64(maxUint8) * float64(maxUint8) * float64(maxUint8)

	input := make(chan uint8, 1)
	output := make(chan float64, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	go Pipeline(input, output, &wg)

	// Send max uint8
	go func() {
		input <- maxUint8
		close(input) // Close input AFTER sending data
	}()

	// Get result
	result := <-output
	wg.Wait()

	if result != expectedResult {
		t.Errorf("Expected result for 255 = %.1f, got %.1f", expectedResult, result)
	}
}
