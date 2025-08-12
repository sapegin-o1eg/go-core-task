package utils

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		expected bool
		result   []int
	}{
		{
			name:     "Normal intersection",
			a:        []int{65, 3, 58, 678, 64},
			b:        []int{64, 2, 3, 43},
			expected: true,
			result:   []int{3, 64},
		},
		{
			name:     "Empty first array",
			a:        []int{},
			b:        []int{1, 2, 3},
			expected: false,
			result:   []int{},
		},
		{
			name:     "Empty second array",
			a:        []int{1, 2, 3},
			b:        []int{},
			expected: false,
			result:   []int{},
		},
		{
			name:     "Both arrays empty",
			a:        []int{},
			b:        []int{},
			expected: false,
			result:   []int{},
		},
		{
			name:     "No intersection",
			a:        []int{1, 2, 3},
			b:        []int{4, 5, 6},
			expected: false,
			result:   []int{},
		},
		{
			name:     "All elements intersect",
			a:        []int{1, 2, 3},
			b:        []int{1, 2, 3},
			expected: true,
			result:   []int{1, 2, 3},
		},
		{
			name:     "Duplicate elements in first array",
			a:        []int{1, 1, 2, 3},
			b:        []int{1, 2, 4},
			expected: true,
			result:   []int{1, 2},
		},
		{
			name:     "Duplicate elements in second array",
			a:        []int{1, 2, 3},
			b:        []int{1, 1, 2, 4},
			expected: true,
			result:   []int{1, 2},
		},
		{
			name:     "First array larger than second",
			a:        []int{1, 2, 3, 4, 5},
			b:        []int{2, 4},
			expected: true,
			result:   []int{2, 4},
		},
		{
			name:     "Second array larger than first",
			a:        []int{2, 4},
			b:        []int{1, 2, 3, 4, 5},
			expected: true,
			result:   []int{2, 4},
		},
		{
			name:     "Single element arrays with intersection",
			a:        []int{42},
			b:        []int{42},
			expected: true,
			result:   []int{42},
		},
		{
			name:     "Single element arrays without intersection",
			a:        []int{42},
			b:        []int{24},
			expected: false,
			result:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, result := Intersection(tt.a, tt.b)

			if ok != tt.expected {
				t.Errorf("Intersection() ok = %v, expected %v", ok, tt.expected)
			}

			if !reflect.DeepEqual(result, tt.result) {
				t.Errorf("Intersection() result = %v, expected %v", result, tt.result)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "First number smaller",
			a:        5,
			b:        10,
			expected: 5,
		},
		{
			name:     "Second number smaller",
			a:        15,
			b:        8,
			expected: 8,
		},
		{
			name:     "Equal numbers",
			a:        7,
			b:        7,
			expected: 7,
		},
		{
			name:     "Negative numbers",
			a:        -5,
			b:        -3,
			expected: -5,
		},
		{
			name:     "Zero and positive",
			a:        0,
			b:        100,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := min(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("min(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
