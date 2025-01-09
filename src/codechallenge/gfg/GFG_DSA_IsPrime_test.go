package gfg

import (
	"fmt"
	"math"
	"testing"
)

func IsPrime(n int) bool {
	if n == 1 {
		return false
	}

	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func IsPrime_Optimized(n int) bool {
	if n == 1 {
		return false
	}

	if n == 2 || n == 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	sqrtN := int(math.Sqrt(float64(n)))
	for i := 5; i <= sqrtN; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func TestIsPrime(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{
			input:    2,
			expected: true,
		},
		{
			input:    4,
			expected: false,
		},
		{
			input:    13,
			expected: true,
		},
		{
			input:    14,
			expected: false,
		},
		{
			input:    101,
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := IsPrime(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}

			result = IsPrime_Optimized(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

// go test -run ^Test.*IsPrime$ example.com/app/codechallenge/gfg -count=1 -v

// Example benchmark usage:
//  go test codechallenge/gfg/GFG_DSA_IsPrime_test.go -count=1 -bench=.
//
// Example benchmark output:
// 	goos: linux
// 	goarch: amd64
// 	cpu: AMD Ryzen 9 PRO 6950HS with Radeon Graphics
// 	BenchmarkIsPrime_Optimized-16           70162160                17.58 ns/op
// 	BenchmarkIsPrime-16                     22306508                57.64 ns/op
// 	PASS
// 	ok      command-line-arguments  2.602s

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime(1031)
	}
}

func BenchmarkIsPrime_Optimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPrime_Optimized(1031)
	}
}
