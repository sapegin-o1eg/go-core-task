package utils

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	testsWithoutPassingRemoveDuplicates := []struct {
		name    string
		source  []string
		exclude []string
		want    []string
	}{
		{
			name:    "Test unique items are not removed",
			source:  []string{"apple", "banana", "banana", "cherry", "cherry", "date", "43", "lead", "gno1"},
			exclude: []string{"banana", "banana", "date", "fig"},
			want:    []string{"apple", "cherry", "cherry", "43", "lead", "gno1"},
		},
		{
			name:    "Test empty source",
			source:  []string{},
			exclude: []string{"banana", "banana", "date", "fig"},
			want:    []string{},
		},
		{
			name:    "Test empty exclude",
			source:  []string{"apple", "banana", "banana", "cherry", "cherry", "date", "43", "lead", "gno1"},
			exclude: []string{},
			want:    []string{"apple", "banana", "banana", "cherry", "cherry", "date", "43", "lead", "gno1"},
		},
		{
			name:    "Test empty source and exclude",
			source:  []string{},
			exclude: []string{},
			want:    []string{},
		},
	}

	for _, tt := range testsWithoutPassingRemoveDuplicates {
		t.Run(tt.name, func(t *testing.T) {
			got := Difference(tt.source, tt.exclude)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference(%v, %v) = %v, want %v", tt.source, tt.exclude, got, tt.want)
			}
		})
	}
}
