package gfg

import (
	"fmt"
	"testing"
)

func GCD(a, b int) int {
	if a < 0 {
		a = -a
	}

	if b < 0 {
		b = -b
	}

	if b == 0 {
		return a
	}

	return GCD(b, a%b)
}

func TestGCD(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected int
	}{
		{
			a:        4,
			b:        6,
			expected: 2,
		},
		{
			a:        100,
			b:        200,
			expected: 100,
		},
		{
			a:        7,
			b:        13,
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v, %v", tc.a, tc.b), func(t *testing.T) {
			result := GCD(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

// go test -run ^Test.*GCD$ example.com/app/codechallenge/gfg -count=1 -v
