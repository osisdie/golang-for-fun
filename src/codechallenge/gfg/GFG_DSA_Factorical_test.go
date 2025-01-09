package gfg

import (
	"fmt"
	"testing"
)

func Factorical_Recursive(n int) int {
	if n == 0 {
		return 1
	}

	return n * Factorical_Recursive(n-1)
}

func Factorical(n int) int {
	fact := 1
	for i := 2; i <= n; i++ {
		fact *= i
	}

	return fact
}

func TestFactorical(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{
			input:    4,
			expected: 24,
		},
		{
			input:    6,
			expected: 720,
		},
		{
			input:    0,
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := Factorical(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

// Example benchmark usage:
//  go test codechallenge/gfg/GFG_DSA_Factorical_test.go -count=1 -bench=.

// Example benchmark output:
// 	goos: linux
// 	goarch: amd64
// 	cpu: AMD Ryzen 9 PRO 6950HS with Radeon Graphics
// 	BenchmarkFactorical_Recursive-16        37806259                29.74 ns/op
// 	BenchmarkFactorical-16                  128587597                9.678 ns/op
// 	PASS
// 	ok      command-line-arguments  3.356s

func BenchmarkFactorical(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorical(20)
	}
}

func BenchmarkFactorical_Recursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorical_Recursive(20)
	}
}
