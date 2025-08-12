package utils

import (
	"testing"
)

func TestNewStringIntMap(t *testing.T) {
	// Test creating new StringIntMap
	sm := NewStringIntMap()

	if sm == nil {
		t.Fatal("NewStringIntMap returned nil")
	}

	if sm.M == nil {
		t.Fatal("Map field is nil")
	}

	if len(sm.M) != 0 {
		t.Errorf("Expected empty map, got map with %d elements", len(sm.M))
	}
}

func TestAdd(t *testing.T) {
	sm := NewStringIntMap()

	// Test adding single element
	sm.Add("key1", 100)

	if len(sm.M) != 1 {
		t.Errorf("Expected 1 element, got %d", len(sm.M))
	}

	if sm.M["key1"] != 100 {
		t.Errorf("Expected value 100, got %d", sm.M["key1"])
	}

	// Test adding multiple elements
	sm.Add("key2", 200)
	sm.Add("key3", 300)

	if len(sm.M) != 3 {
		t.Errorf("Expected 3 elements, got %d", len(sm.M))
	}

	// Test overwriting existing key
	sm.Add("key1", 150)

	if sm.M["key1"] != 150 {
		t.Errorf("Expected overwritten value 150, got %d", sm.M["key1"])
	}

	if len(sm.M) != 3 {
		t.Errorf("Expected still 3 elements after overwrite, got %d", len(sm.M))
	}
}

func TestRemove(t *testing.T) {
	sm := NewStringIntMap()

	// Add some elements
	sm.Add("key1", 100)
	sm.Add("key2", 200)
	sm.Add("key3", 300)

	// Test removing existing element
	sm.Remove("key2")

	if len(sm.M) != 2 {
		t.Errorf("Expected 2 elements after removal, got %d", len(sm.M))
	}

	if _, exists := sm.M["key2"]; exists {
		t.Error("Key 'key2' should not exist after removal")
	}

	// Test removing non-existent element
	sm.Remove("nonexistent")

	if len(sm.M) != 2 {
		t.Errorf("Expected still 2 elements after removing non-existent key, got %d", len(sm.M))
	}

	// Test removing all elements
	sm.Remove("key1")
	sm.Remove("key3")

	if len(sm.M) != 0 {
		t.Errorf("Expected empty map after removing all elements, got %d elements", len(sm.M))
	}
}

func TestCopy(t *testing.T) {
	sm := NewStringIntMap()

	// Test copying empty map
	copy1 := sm.Copy()

	if len(copy1) != 0 {
		t.Errorf("Expected empty copy, got map with %d elements", len(copy1))
	}

	// Add elements and test copying
	sm.Add("key1", 100)
	sm.Add("key2", 200)
	sm.Add("key3", 300)

	copy2 := sm.Copy()

	// Check that copy has same elements
	if len(copy2) != 3 {
		t.Errorf("Expected copy with 3 elements, got %d", len(copy2))
	}

	if copy2["key1"] != 100 || copy2["key2"] != 200 || copy2["key3"] != 300 {
		t.Error("Copy does not contain expected values")
	}

	// Test that copy is independent
	copy2["key1"] = 999
	copy2["newkey"] = 500

	if sm.M["key1"] != 100 {
		t.Error("Original map should not be affected by changes to copy")
	}

	if _, exists := sm.M["newkey"]; exists {
		t.Error("Original map should not contain new key added to copy")
	}

	if len(sm.M) != 3 {
		t.Errorf("Original map should still have 3 elements, got %d", len(sm.M))
	}
}

func TestGet(t *testing.T) {
	sm := NewStringIntMap()

	// Test getting from empty map
	value, ok := sm.Get("nonexistent")

	if ok {
		t.Error("Expected false for non-existent key")
	}

	if value != 0 {
		t.Errorf("Expected zero value for non-existent key, got %d", value)
	}

	// Add elements and test getting
	sm.Add("key1", 100)
	sm.Add("key2", 200)

	// Test getting existing key
	value, ok = sm.Get("key1")

	if !ok {
		t.Error("Expected true for existing key")
	}

	if value != 100 {
		t.Errorf("Expected value 100, got %d", value)
	}

	// Test getting another existing key
	value, ok = sm.Get("key2")

	if !ok {
		t.Error("Expected true for existing key")
	}

	if value != 200 {
		t.Errorf("Expected value 200, got %d", value)
	}

	// Test getting non-existent key
	value, ok = sm.Get("key3")

	if ok {
		t.Error("Expected false for non-existent key")
	}

	if value != 0 {
		t.Errorf("Expected zero value for non-existent key, got %d", value)
	}
}

func TestExists(t *testing.T) {
	sm := NewStringIntMap()

	// Test checking non-existent key in empty map
	if sm.Exists("nonexistent") {
		t.Error("Expected false for non-existent key in empty map")
	}

	// Add elements and test
	sm.Add("key1", 100)
	sm.Add("key2", 200)

	// Test checking existing keys
	if !sm.Exists("key1") {
		t.Error("Expected true for existing key 'key1'")
	}

	if !sm.Exists("key2") {
		t.Error("Expected true for existing key 'key2'")
	}

	// Test checking non-existent key
	if sm.Exists("key3") {
		t.Error("Expected false for non-existent key 'key3'")
	}

	// Test after removal
	sm.Remove("key1")

	if sm.Exists("key1") {
		t.Error("Expected false for removed key 'key1'")
	}

	if !sm.Exists("key2") {
		t.Error("Expected true for remaining key 'key2'")
	}
}

func TestStringIntMap_Integration(t *testing.T) {
	// Test integration of all methods together
	sm := NewStringIntMap()

	// Add multiple elements
	elements := map[string]int{
		"apple":  1,
		"banana": 2,
		"cherry": 3,
		"date":   4,
		"elder":  5,
	}

	for key, value := range elements {
		sm.Add(key, value)
	}

	// Verify all elements exist
	for key, expectedValue := range elements {
		if !sm.Exists(key) {
			t.Errorf("Key '%s' should exist", key)
		}

		value, ok := sm.Get(key)
		if !ok {
			t.Errorf("Get should return true for key '%s'", key)
		}

		if value != expectedValue {
			t.Errorf("Expected value %d for key '%s', got %d", expectedValue, key, value)
		}
	}

	// Test copy
	copy := sm.Copy()
	if len(copy) != len(elements) {
		t.Errorf("Copy should have %d elements, got %d", len(elements), len(copy))
	}

	// Remove some elements
	sm.Remove("banana")
	sm.Remove("date")

	// Verify removal
	if sm.Exists("banana") {
		t.Error("Removed key 'banana' should not exist")
	}

	if sm.Exists("date") {
		t.Error("Removed key 'date' should not exist")
	}

	if !sm.Exists("apple") {
		t.Error("Non-removed key 'apple' should still exist")
	}

	// Verify copy is unchanged
	if copy["banana"] != 2 {
		t.Error("Copy should not be affected by removal from original")
	}

	if copy["date"] != 4 {
		t.Error("Copy should not be affected by removal from original")
	}
}
