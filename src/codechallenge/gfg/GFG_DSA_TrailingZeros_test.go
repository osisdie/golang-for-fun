package gfg

import (
	"fmt"
	"testing"
)

func TrailingZeros(n int) int {
	count := 0

	for i := 5; i <= n; i *= 5 {
		count += n / i
	}

	return count
}

func TestTrailingZeros(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{
			input:    5,
			expected: 1,
		},
		{
			input:    10,
			expected: 2,
		},
		{
			input:    100,
			expected: 24,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := TrailingZeros(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
