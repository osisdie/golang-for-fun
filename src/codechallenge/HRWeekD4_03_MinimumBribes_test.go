package codechallenge

import (
	"fmt"
	"strconv"
	"testing"
)

func MinimumBribes(q []int) string {
	ary := make([]int, len(q))
	copy(ary, q)

	moves := 0
	person := len(ary)
	for person > 1 {
		initPos := person - 1 // Initial position
		currPos := findIndex(ary, person)
		move := initPos - currPos
		moves += move
		if move > 2 {
			return "Too chaotic"
		}

		if move > 0 {
			// Swap current position to the right position
			for j := currPos; j < initPos; j++ {
				ary[j], ary[j+1] = ary[j+1], ary[j]
			}
		}

		person--
	}

	return strconv.Itoa(moves)
}

// Helper function to find the index of a value in a slice
func findIndex(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func TestMinimumBribes(t *testing.T) {
	// Define test cases as a slice of structs
	testCases := []struct {
		input    []int
		expected string
	}{
		{[]int{2, 1, 5, 3, 4}, "3"},
		{[]int{2, 5, 1, 3, 4}, "Too chaotic"},
		{[]int{1, 2, 5, 3, 4, 7, 8, 6}, "4"},
		{[]int{5, 1, 2, 3, 7, 8, 6, 4}, "Too chaotic"},
		{[]int{1, 2, 5, 3, 7, 8, 6, 4}, "7"},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := MinimumBribes(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
