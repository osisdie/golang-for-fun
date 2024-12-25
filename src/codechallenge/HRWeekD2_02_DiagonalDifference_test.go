package codechallenge

import (
	"fmt"
	"math"
	"testing"
)

func DiagonalDifference(arr [][]int) int {
	len := len(arr)
	sumLR := 0
	sumRL := 0

	for i := 0; i < len; i++ {
		sumLR += arr[i][i]
		sumRL += arr[i][len-i-1]
	}

	return int(math.Abs(float64(sumLR - sumRL)))
}

func TestDiagonalDifference(t *testing.T) {
	// Define test cases as a slice of structs
	testCases := []struct {
		input    [][]int
		expected int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{9, 8, 9},
			},
			expected: 2,
		},
		{
			input: [][]int{
				{11, 2, 4},
				{4, 5, 6},
				{10, 8, -12},
			},
			expected: 15,
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := DiagonalDifference(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
