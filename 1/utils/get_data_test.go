package utils

import (
	"reflect"
	"testing"
	"task1/models"
)

func TestGetSomeIntegerNumbers(t *testing.T) {
	result := GetSomeIntegerNumbers()
	
	// Check if result is not nil
	if result == nil {
		t.Error("Expected non-nil slice, got nil")
	}
	
	// Check expected length
	expectedLength := 9
	if len(result) != expectedLength {
		t.Errorf("Expected length %d, got %d", expectedLength, len(result))
	}
	
	// Check expected values
	expectedValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(result, expectedValues) {
		t.Errorf("Expected %v, got %v", expectedValues, result)
	}
	
	// Check that all values are integers
	for i, val := range result {
		if reflect.TypeOf(val).Kind() != reflect.Int {
			t.Errorf("Expected int type at index %d, got %T", i, val)
		}
	}
}

func TestGetSomeFloat64Numbers(t *testing.T) {
	first, second, third := GetSomeFloat64Numbers()
	
	// Check expected values
	expectedFirst := 1.01
	expectedSecond := 2.02
	expectedThird := 3.03
	
	if first != expectedFirst {
		t.Errorf("Expected first value %f, got %f", expectedFirst, first)
	}
	
	if second != expectedSecond {
		t.Errorf("Expected second value %f, got %f", expectedSecond, second)
	}
	
	if third != expectedThird {
		t.Errorf("Expected third value %f, got %f", expectedThird, third)
	}
	
	// Check types
	if reflect.TypeOf(first).Kind() != reflect.Float64 {
		t.Errorf("Expected float64 type for first value, got %T", first)
	}
	
	if reflect.TypeOf(second).Kind() != reflect.Float64 {
		t.Errorf("Expected float64 type for second value, got %T", second)
	}
	
	if reflect.TypeOf(third).Kind() != reflect.Float64 {
		t.Errorf("Expected float64 type for third value, got %T", third)
	}
}

func TestGetSomeStrings(t *testing.T) {
	first, second := GetSomeStrings()
	
	// Check expected values
	expectedFirst := "It's the first string"
	expectedSecond := "It's the second string"
	
	if first != expectedFirst {
		t.Errorf("Expected first string '%s', got '%s'", expectedFirst, first)
	}
	
	if second != expectedSecond {
		t.Errorf("Expected second string '%s', got '%s'", expectedSecond, second)
	}
	
	// Check types
	if reflect.TypeOf(first).Kind() != reflect.String {
		t.Errorf("Expected string type for first value, got %T", first)
	}
	
	if reflect.TypeOf(second).Kind() != reflect.String {
		t.Errorf("Expected string type for second value, got %T", second)
	}
}

func TestGetSomeBooleans(t *testing.T) {
	first, second := GetSomeBooleans()
	
	// Check expected values
	expectedFirst := true
	expectedSecond := false
	
	if first != expectedFirst {
		t.Errorf("Expected first boolean %t, got %t", expectedFirst, first)
	}
	
	if second != expectedSecond {
		t.Errorf("Expected second boolean %t, got %t", expectedSecond, second)
	}
	
	// Check types
	if reflect.TypeOf(first).Kind() != reflect.Bool {
		t.Errorf("Expected bool type for first value, got %T", first)
	}
	
	if reflect.TypeOf(second).Kind() != reflect.Bool {
		t.Errorf("Expected bool type for second value, got %T", second)
	}
}

func TestGetSomeComplex64Numbers(t *testing.T) {
	result := GetSomeComplex64Numbers()
	
	// Check if result is not nil
	if reflect.ValueOf(result).IsZero() {
		t.Error("Expected non-zero struct, got zero value")
	}
	
	// Check expected values
	expectedFirst := complex64(1 + 1i)
	expectedSecond := complex64(2 + 2i)
	
	if result.First != expectedFirst {
		t.Errorf("Expected first complex64 %v, got %v", expectedFirst, result.First)
	}
	
	if result.Second != expectedSecond {
		t.Errorf("Expected second complex64 %v, got %v", expectedSecond, result.Second)
	}
	
	// Check types
	if reflect.TypeOf(result.First).Kind() != reflect.Complex64 {
		t.Errorf("Expected complex64 type for first value, got %T", result.First)
	}
	
	if reflect.TypeOf(result.Second).Kind() != reflect.Complex64 {
		t.Errorf("Expected complex64 type for second value, got %T", result.Second)
	}
	
	// Check that result is of correct struct type
	expectedType := reflect.TypeOf(models.Complex64Numbers{})
	if reflect.TypeOf(result) != expectedType {
		t.Errorf("Expected type %v, got %T", expectedType, result)
	}
} 