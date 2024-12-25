package codechallenge

import (
	"fmt"
	"testing"
)

func Lonelyinteger(a []int) int {
	numMap := make(map[int]int)

	// Populate the map with counts of each number
	for _, num := range a {
		numMap[num]++
	}

	// Find the first number with a count of 1
	for num, count := range numMap {
		if count == 1 {
			return num
		}
	}

	// If no unique number is found, return 0 (or handle as needed)
	return 0
}

func TestLonelyinteger(t *testing.T) {
	// Define test cases as a slice of structs
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{1, 1, 2}, 2},
		{[]int{1, 2, 3, 4, 3, 2, 1}, 4},
		{[]int{4, 9, 95, 93, 57, 4, 57, 93, 9}, 95},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := Lonelyinteger(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
