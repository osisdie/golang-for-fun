package gfg

import (
	"fmt"
	"testing"
)

func Palindrome(n int) bool {
	original_n := n
	revsered_n := 0

	for n > 0 {
		revsered_n = (revsered_n * 10) + n%10
		n /= 10
	}

	return revsered_n == original_n
}

func TestPalindrome(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{
			input:    78987,
			expected: true,
		},
		{
			input:    8668,
			expected: true,
		},
		{
			input:    8,
			expected: true,
		},
		{
			input:    21,
			expected: false,
		},
		{
			input:    367,
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := Palindrome(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
