package gfg

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func AllDivisors(n int) []int {
	res := []int{}

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			res = append(res, i)
		}
	}

	return res
}

func AllDivisors_Optimized(n int) []int {
	res := []int{}

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			res = append(res, i)
			if n/i != i {
				res = append(res, n/i)
			}
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return res
}

func TestAllDivisors(t *testing.T) {
	testCases := []struct {
		input    int
		expected []int
	}{
		{
			input:    15,
			expected: []int{1, 3, 5, 15},
		},
		{
			input:    100,
			expected: []int{1, 2, 4, 5, 10, 20, 25, 50, 100},
		},
		{
			input:    7,
			expected: []int{1, 7},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := AllDivisors(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}

			result = AllDivisors_Optimized(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

// go test codechallenge/gfg/GFG_DSA_AllDivisors_test.go -count=1 -v
