package codechallenge

import (
	"fmt"
	"sort"
	"testing"
)

func GridChallenge(grid []string) string {
	rows := len(grid)
	if rows == 0 {
		return "YES" // Empty grid is considered sorted
	}

	cols := len(grid[0])

	// Convert each row to a sorted slice of runes
	matrix := make([][]rune, rows)
	for i, row := range grid {
		runes := []rune(row)
		sort.Slice(runes, func(j, k int) bool {
			return runes[j] < runes[k]
		})
		matrix[i] = runes
	}

	// Check if each column is sorted
	for i := 0; i < cols; i++ {
		for j := 0; j < rows-1; j++ {
			if matrix[j][i] > matrix[j+1][i] {
				return "NO"
			}
		}
	}

	return "YES"
}

func TestGridChallenge(t *testing.T) {
	// Define test cases as a slice of structs
	testCases := []struct {
		input    []string
		expected string
	}{
		{
			[]string{"ebacd", "fghij", "olmkn", "trpqs", "xywuv"},
			"YES",
		},
		{
			[]string{"abc", "lmp", "qrt"},
			"YES",
		},
		{
			[]string{"mpxz", "abcd", "wlmf"},
			"NO",
		},
		{
			[]string{"abc", "hjk", "mpq", "rtv"},
			"YES",
		},
		{
			[]string{"zx", "fw"},
			"NO",
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := GridChallenge(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
