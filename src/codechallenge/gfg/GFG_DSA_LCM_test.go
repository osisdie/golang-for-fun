package gfg

import (
	"fmt"
	"testing"
)

// LCM(a, b) = (a x b) / GCD(a, b)
func LCM(a, b int) int {
	return (a / GCD(a, b)) * b
}

func TestLCM(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected int
	}{
		{
			a:        4,
			b:        6,
			expected: 12,
		},
		{
			a:        12,
			b:        15,
			expected: 60,
		},
		{
			a:        2,
			b:        8,
			expected: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v, %v", tc.a, tc.b), func(t *testing.T) {
			result := LCM(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

// go test -run ^Test.*LCM$ example.com/app/codechallenge/gfg -count=1 -v
