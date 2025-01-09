package gfg

import (
	"fmt"
	"testing"
)

func CountDigits(n int) int {
	if n < 0 {
		n = -n
	}

	if n == 0 {
		return 1
	}

	count := 0
	for n > 0 {
		count += 1
		n /= 10
	}

	return count
}

func TestCountDigits(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{
			input:    9235,
			expected: 4,
		},
		{
			input:    38,
			expected: 2,
		},
		{
			input:    7,
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := CountDigits(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
