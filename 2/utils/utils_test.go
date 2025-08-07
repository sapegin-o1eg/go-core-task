package utils

import (
	"reflect"
	"testing"
)

func TestSliceExample1(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		verbose  bool
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			verbose:  false,
			expected: []int{},
		},
		{
			name:     "no even numbers",
			input:    []int{1, 3, 5, 7, 9},
			verbose:  false,
			expected: []int{},
		},
		{
			name:     "all even numbers",
			input:    []int{2, 4, 6, 8, 10},
			verbose:  false,
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "mixed even and odd numbers",
			input:    []int{1, 2, 3, 4, 5, 6},
			verbose:  false,
			expected: []int{2, 4, 6},
		},
		{
			name:     "single even number",
			input:    []int{1, 2, 3},
			verbose:  false,
			expected: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SliceExample1(tt.input, tt.verbose)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SliceExample1() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSliceExample2(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		verbose  bool
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			verbose:  false,
			expected: []int{},
		},
		{
			name:     "no even numbers",
			input:    []int{1, 3, 5, 7, 9},
			verbose:  false,
			expected: []int{},
		},
		{
			name:     "all even numbers",
			input:    []int{2, 4, 6, 8, 10},
			verbose:  false,
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "mixed even and odd numbers",
			input:    []int{1, 2, 3, 4, 5, 6},
			verbose:  false,
			expected: []int{2, 4, 6},
		},
		{
			name:     "single even number",
			input:    []int{1, 2, 3},
			verbose:  false,
			expected: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SliceExample2(tt.input, tt.verbose)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SliceExample2() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		values   []int
		expected []int
	}{
		{
			name:     "empty slice add elements",
			input:    []int{},
			values:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "add to existing slice",
			input:    []int{1, 2, 3},
			values:   []int{4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "add no elements",
			input:    []int{1, 2, 3},
			values:   []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "add single element",
			input:    []int{1, 2},
			values:   []int{3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddElements(tt.input, tt.values...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("AddElements() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "multiple elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CopySlice(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CopySlice() = %v, want %v", result, tt.expected)
			}
		})
	}

	// Test that copy is independent
	t.Run("copy independence", func(t *testing.T) {
		original := []int{1, 2, 3, 4, 5}
		copied := CopySlice(original)

		// Modify original
		original[0] = 999

		if copied[0] == 999 {
			t.Error("CopySlice() should create independent copy, but original modification affected copied slice")
		}
		if copied[0] != 1 {
			t.Errorf("CopySlice() should preserve original values, got %v, want 1", copied[0])
		}
	})
}

func TestRemoveElements(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		indexes  []int
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			indexes:  []int{0, 1},
			expected: []int{},
		},
		{
			name:     "remove first element",
			input:    []int{1, 2, 3, 4, 5},
			indexes:  []int{0},
			expected: []int{2, 3, 4, 5},
		},
		{
			name:     "remove last element",
			input:    []int{1, 2, 3, 4, 5},
			indexes:  []int{4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "remove multiple elements",
			input:    []int{1, 2, 3, 4, 5},
			indexes:  []int{0, 2, 4},
			expected: []int{2, 4},
		},
		{
			name:     "remove no elements",
			input:    []int{1, 2, 3, 4, 5},
			indexes:  []int{},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "remove out of bounds indexes",
			input:    []int{1, 2, 3},
			indexes:  []int{5, 10},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveElements(tt.input, tt.indexes...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("RemoveElements() = %v, want %v", result, tt.expected)
			}
		})
	}

	// Test that original slice is not modified
	t.Run("original slice unchanged", func(t *testing.T) {
		original := []int{1, 2, 3, 4, 5}
		originalCopy := make([]int, len(original))
		copy(originalCopy, original)

		RemoveElements(original, 0, 1)

		if !reflect.DeepEqual(original, originalCopy) {
			t.Error("RemoveElements() should not modify original slice")
		}
	})
}

// Benchmark tests for performance comparison
func BenchmarkSliceExample1(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceExample1(input, false)
	}
}

func BenchmarkSliceExample2(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceExample2(input, false)
	}
}
