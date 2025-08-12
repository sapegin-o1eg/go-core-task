package utils

import (
	"sync"
	"testing"
	"time"
)

// TestNewCustomWaitGroup tests the constructor
func TestNewCustomWaitGroup(t *testing.T) {
	wg := NewCustomWaitGroup()
	if wg == nil {
		t.Fatal("NewCustomWaitGroup returned nil")
	}
	if wg.counter != 0 {
		t.Errorf("Expected counter to be 0, got %d", wg.counter)
	}
	if wg.sem == nil {
		t.Fatal("Semaphore channel is nil")
	}
}

// TestCustomWaitGroup_Add tests the Add method
func TestCustomWaitGroup_Add(t *testing.T) {
	wg := NewCustomWaitGroup()

	// Test positive delta
	wg.Add(5)
	if wg.counter != 5 {
		t.Errorf("Expected counter to be 5, got %d", wg.counter)
	}

	// Test negative delta
	wg.Add(-2)
	if wg.counter != 3 {
		t.Errorf("Expected counter to be 3, got %d", wg.counter)
	}

	// Test zero delta
	wg.Add(0)
	if wg.counter != 3 {
		t.Errorf("Expected counter to be 3, got %d", wg.counter)
	}
}

// TestCustomWaitGroup_AddNegative tests that Add panics with negative counter
func TestCustomWaitGroup_AddNegative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when counter becomes negative")
		}
	}()

	wg := NewCustomWaitGroup()
	wg.Add(1)
	wg.Add(-2) // This should panic
}

// TestCustomWaitGroup_Done tests the Done method
func TestCustomWaitGroup_Done(t *testing.T) {
	wg := NewCustomWaitGroup()
	wg.Add(3)

	wg.Done()
	if wg.counter != 2 {
		t.Errorf("Expected counter to be 2, got %d", wg.counter)
	}

	wg.Done()
	if wg.counter != 1 {
		t.Errorf("Expected counter to be 1, got %d", wg.counter)
	}
}

// TestCustomWaitGroup_Wait tests the Wait method
func TestCustomWaitGroup_Wait(t *testing.T) {
	wg := NewCustomWaitGroup()

	// Test immediate return when counter is 0
	start := time.Now()
	wg.Wait()
	elapsed := time.Since(start)

	if elapsed > 10*time.Millisecond {
		t.Error("Wait should return immediately when counter is 0")
	}
}

// TestCustomWaitGroup_WaitWithGoroutines tests Wait with multiple goroutines
func TestCustomWaitGroup_WaitWithGoroutines(t *testing.T) {
	wg := NewCustomWaitGroup()
	const numGoroutines = 5

	// Start multiple goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
		}()
	}

	// Wait for all goroutines to complete
	start := time.Now()
	wg.Wait()
	elapsed := time.Since(start)

	if elapsed < 10*time.Millisecond {
		t.Error("Wait returned too quickly, goroutines may not have completed")
	}
}

// TestCustomWaitGroup_ConcurrentAccess tests concurrent access to WaitGroup
func TestCustomWaitGroup_ConcurrentAccess(t *testing.T) {
	wg := NewCustomWaitGroup()
	const numGoroutines = 100

	var wgSync sync.WaitGroup
	wgSync.Add(numGoroutines)

	// Start goroutines that add and then done
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wgSync.Done()
			wg.Add(1)
			time.Sleep(1 * time.Millisecond)
			wg.Done()
		}()
	}

	// Wait for all test goroutines to complete
	wgSync.Wait()

	// Final counter should be 0
	if wg.counter != 0 {
		t.Errorf("Expected final counter to be 0, got %d", wg.counter)
	}
}

// TestCustomWaitGroup_Reuse tests reusing the same WaitGroup
func TestCustomWaitGroup_Reuse(t *testing.T) {
	wg := NewCustomWaitGroup()

	// First use
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Millisecond)
	}()
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Millisecond)
	}()
	wg.Wait()

	wg.Reset()

	// Second use
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Millisecond)
	}()
	wg.Wait()

	if wg.counter != 0 {
		t.Errorf("Expected final counter to be 0, got %d", wg.counter)
	}
}

// TestCustomWaitGroup_Reset tests the Reset method
func TestCustomWaitGroup_Reset(t *testing.T) {
	wg := NewCustomWaitGroup()

	// Add some work
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Millisecond)
	}()

	// Wait and reset
	wg.Wait()
	wg.Reset()

	// Verify reset state
	if wg.counter != 0 {
		t.Errorf("Expected counter to be 0 after reset, got %d", wg.counter)
	}
	if wg.closed {
		t.Error("Expected closed to be false after reset")
	}

	// Test reuse after reset
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Millisecond)
	}()
	wg.Wait()

	if wg.counter != 0 {
		t.Errorf("Expected final counter to be 0, got %d", wg.counter)
	}
}

// BenchmarkCustomWaitGroup benchmarks the performance
func BenchmarkCustomWaitGroup(b *testing.B) {
	b.Run("Add", func(b *testing.B) {
		wg := NewCustomWaitGroup()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			wg.Add(1)
		}
	})

	b.Run("Done", func(b *testing.B) {
		wg := NewCustomWaitGroup()
		wg.Add(int64(b.N))
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			wg.Done()
		}
	})

	b.Run("Wait", func(b *testing.B) {
		wg := NewCustomWaitGroup()
		wg.Add(1)
		go func() {
			defer wg.Done()
		}()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			wg.Wait()
		}
	})
}
