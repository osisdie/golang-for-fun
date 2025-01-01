package codechallenge

import (
	"fmt"
	"testing"
)

func TowerBreakers(n, m int) int {
	if m == 1 { // Player 1 cannot move
		return 2
	}

	if n%2 == 1 { // Player 1 can always make one tower height 1
		return 1
	} else { // Player 2 can always copy Player 1's move
		return 2
	}
}

func TestTowerBreakers(t *testing.T) {
	// Define test cases as a slice of structs
	testCases := []struct {
		n        int
		m        int
		expected int
	}{
		{2, 2, 2},
		{1, 4, 1},
		{1, 7, 1},
		{3, 7, 1},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%d, m=%d", tc.n, tc.m), func(t *testing.T) {
			result := TowerBreakers(tc.n, tc.m)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
