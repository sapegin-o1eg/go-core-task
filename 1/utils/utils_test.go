package utils

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetTypeAsString(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     "int type",
			input:    42,
			expected: "int",
		},
		{
			name:     "string type",
			input:    "hello",
			expected: "string",
		},
		{
			name:     "float64 type",
			input:    3.14,
			expected: "float64",
		},
		{
			name:     "bool type",
			input:    true,
			expected: "bool",
		},
		{
			name:     "slice type",
			input:    []int{1, 2, 3},
			expected: "[]int",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetTypeAsString(tt.input)
			if result != tt.expected {
				t.Errorf("GetTypeAsString(%v) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConcatenateStrings(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "empty args",
			args:     []string{},
			expected: "",
		},
		{
			name:     "single string",
			args:     []string{"hello"},
			expected: "hello",
		},
		{
			name:     "multiple strings",
			args:     []string{"hello", "world", "test"},
			expected: "helloworldtest",
		},
		{
			name:     "strings with spaces",
			args:     []string{"hello ", "world ", "test"},
			expected: "hello world test",
		},
		{
			name:     "empty strings",
			args:     []string{"", "", ""},
			expected: "",
		},
		{
			name:     "mixed content",
			args:     []string{"int", "string", "bool"},
			expected: "intstringbool",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConcatenateStrings(tt.args...)
			if result != tt.expected {
				t.Errorf("ConcatenateStrings(%v) = %s, want %s", tt.args, result, tt.expected)
			}
		})
	}
}

func TestStringToRunes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []rune
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []rune{},
		},
		{
			name:     "simple string",
			input:    "hello",
			expected: []rune{'h', 'e', 'l', 'l', 'o'},
		},
		{
			name:     "unicode string",
			input:    "привет",
			expected: []rune{'п', 'р', 'и', 'в', 'е', 'т'},
		},
		{
			name:     "mixed unicode",
			input:    "hello世界",
			expected: []rune{'h', 'e', 'l', 'l', 'o', '世', '界'},
		},
		{
			name:     "special characters",
			input:    "!@#$%",
			expected: []rune{'!', '@', '#', '$', '%'},
		},
		{
			name:     "numbers as string",
			input:    "12345",
			expected: []rune{'1', '2', '3', '4', '5'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StringToRunes(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("StringToRunes(%s) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetVarType(t *testing.T) {
	// This function only prints to stdout, so we can only test that it doesn't panic
	// and that it handles various input types correctly

	testCases := []struct {
		name string
		args []interface{}
	}{
		{
			name: "no args",
			args: []interface{}{},
		},
		{
			name: "single arg",
			args: []interface{}{42},
		},
		{
			name: "multiple args",
			args: []interface{}{42, "hello", 3.14, true},
		},
		{
			name: "nil args",
			args: []interface{}{nil, nil},
		},
		{
			name: "complex types",
			args: []interface{}{[]int{1, 2, 3}, map[string]int{"a": 1}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test that function doesn't panic
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("GetVarType panicked: %v", r)
				}
			}()

			GetVarType(tc.args...)
		})
	}
}

func TestHashRunesWithSalt(t *testing.T) {
	tests := []struct {
		name  string
		runes []rune
		salt  []string
	}{
		{
			name:  "empty runes with default salt",
			runes: []rune{},
			salt:  []string{},
		},
		{
			name:  "simple string with default salt",
			runes: []rune{'h', 'e', 'l', 'l', 'o'},
			salt:  []string{},
		},
		{
			name:  "simple string with custom salt",
			runes: []rune{'h', 'e', 'l', 'l', 'o'},
			salt:  []string{"custom"},
		},
		{
			name:  "unicode string with default salt",
			runes: []rune{'п', 'р', 'и', 'в', 'е', 'т'},
			salt:  []string{},
		},
		{
			name:  "single rune with default salt",
			runes: []rune{'a'},
			salt:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HashRunesWithSalt(tt.runes, tt.salt...)

			// Check that result is a valid hex string
			if len(result) != 64 {
				t.Errorf("HashRunesWithSalt() returned hash of length %d, expected 64", len(result))
			}

			// Check that result contains only hex characters
			for _, char := range result {
				if !strings.Contains("0123456789abcdef", string(char)) {
					t.Errorf("HashRunesWithSalt() returned non-hex character: %c", char)
				}
			}
		})
	}
}

func TestHashRunesWithSalt_Consistency(t *testing.T) {
	// Test that the same input always produces the same output
	runes := []rune{'t', 'e', 's', 't'}
	salt := "test-salt"

	hash1 := HashRunesWithSalt(runes, salt)
	hash2 := HashRunesWithSalt(runes, salt)

	if hash1 != hash2 {
		t.Errorf("HashRunesWithSalt() is not consistent: %s != %s", hash1, hash2)
	}
}

func TestHashRunesWithSalt_DifferentSalts(t *testing.T) {
	// Test that different salts produce different hashes
	runes := []rune{'t', 'e', 's', 't'}

	hash1 := HashRunesWithSalt(runes, "salt1")
	hash2 := HashRunesWithSalt(runes, "salt2")

	if hash1 == hash2 {
		t.Errorf("HashRunesWithSalt() with different salts produced same hash: %s", hash1)
	}
}
