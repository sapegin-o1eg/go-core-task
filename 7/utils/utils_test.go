package utils

import (
	"testing"
	"time"
)

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	merged := MergeChannels(ch1, ch2, ch3)

	go func() {
		ch1 <- 1
		ch2 <- 2
		ch3 <- 3
		close(ch1)
		close(ch2)
		close(ch3)
	}()

	var results []int
	for val := range merged {
		results = append(results, val)
	}

	expected := map[int]bool{1: true, 2: true, 3: true}
	if len(results) != 3 {
		t.Errorf("Expected 3 values, got %d", len(results))
	}

	for _, val := range results {
		if !expected[val] {
			t.Errorf("Unexpected value %d", val)
		}
	}
}

func TestMergeChannelsWithDifferentRates(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	merged := MergeChannels(ch1, ch2)

	go func() {
		for i := 0; i < 3; i++ {
			ch1 <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch1)
	}()

	go func() {
		for i := 10; i < 12; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	var results []int
	for val := range merged {
		results = append(results, val)
	}

	if len(results) != 5 {
		t.Errorf("Expected 5 values, got %d", len(results))
	}
}

func TestMergeEmptyChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	merged := MergeChannels(ch1, ch2)

	close(ch1)
	close(ch2)

	_, ok := <-merged
	if ok {
		t.Error("Merged channel should be closed")
	}
}
