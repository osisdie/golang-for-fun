package codechallenge

import (
	"fmt"
	"reflect"
	"testing"
)

func PlusMinus(arr []int) []string {
	len := len(arr)
	pCount, nCount, zCount := 0, 0, 0

	for _, num := range arr {
		if num > 0 {
			pCount++
		} else if num < 0 {
			nCount++
		} else {
			zCount++
		}
	}

	results := []string{
		fmt.Sprintf("%.6f", float64(pCount)/float64(len)),
		fmt.Sprintf("%.6f", float64(nCount)/float64(len)),
		fmt.Sprintf("%.6f", float64(zCount)/float64(len)),
	}

	return results
}

func TestPlusMinus(t *testing.T) {
	// Define test cases as a slice of structs
	testCases := []struct {
		input    []int
		expected []string
	}{
		{
			input:    []int{-4, 3, -9, 0, 4, 1},
			expected: []string{"0.500000", "0.333333", "0.166667"},
		},
		{
			input:    []int{1, 2, 3, -1, -2, -3, 0, 0},
			expected: []string{"0.375000", "0.375000", "0.250000"},
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := PlusMinus(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
