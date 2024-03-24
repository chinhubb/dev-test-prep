package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		name     string
		coll1    []int
		coll2    []int
		coll3    []int
		expected []int
	}{
		{
			name:     "Typical Case",
			coll1:    []int{2, 5, 9},
			coll2:    []int{20, 13, 7},
			coll3:    []int{3, 6, 15},
			expected: []int{2, 3, 5, 6, 9, 20, 13, 7, 15},
		},
		{
			name:     "Empty Collection",
			coll1:    []int{1, 3},
			coll2:    []int{},
			coll3:    []int{4, 6},
			expected: []int{1, 3, 4, 6},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := merge(tc.coll1, tc.coll2, tc.coll3)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected: %v, got: %v", tc.expected, result)
			}
		})
	}
}
