package utils

import (
	"context"
	"testing"
	"time"
)

func TestRandomGenerator_WithContext(t *testing.T) {
	ctx := t.Context()

	const (
		count = 1000
		min   = 1
		max   = 10
	)

	randCh := RandomGenerator(ctx, count, min, max)

	for i := 0; i < 10; i++ {
		num, ok := <-randCh
		if !ok {
			t.Error("Channel closed prematurely")
			return
		}
		if num < min || num > max {
			t.Errorf("Number %d out of range", num)
		}
	}

	select {
	case _, ok := <-randCh:
		if !ok {
			t.Error("Channel should not be closed yet")
		}
	default:
	}
}

func TestRandomGenerator_Cancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(t.Context())

	const count = 100000
	randCh := RandomGenerator(ctx, count, 1, 10)

	cancel()

	time.Sleep(50 * time.Millisecond)

	if num, ok := <-randCh; ok {
		t.Errorf("Expected closed channel, but got value: %d", num)
	}

	select {
	case _, ok := <-randCh:
		if ok {
			t.Error("Channel should be closed")
		}
	default:
		t.Error("Channel should be closed and not blocking")
	}
}