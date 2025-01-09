package gfg

import (
	"fmt"
	"testing"
)

func IterativePower(x, n int) int {
	res := 1

	for i := 0; i < n; i++ {
		res *= x
	}

	return res
}

// O(logN)
func IterativePower_Optimized(x, n, m int) int {
	if n < 0 {
		return 0
	}

	res := 1

	for n > 0 {
		if n&1 == 1 {
			res = (res * x) % m
		}

		x = (x * x) % m
		n >>= 1
	}

	return res
}

func TestIterativePower(t *testing.T) {
	testCases := []struct {
		x, y     int
		expected int
	}{
		{
			x:        3,
			y:        5,
			expected: 243,
		},
		{
			x:        2,
			y:        5,
			expected: 32,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v, %v", tc.x, tc.y), func(t *testing.T) {
			result := IterativePower(tc.x, tc.y)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}

			result = IterativePower_Optimized(tc.x, tc.y, 1000000007)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

// go test codechallenge/gfg/GFG_DSA_IterativePower_test.go -count=1 -v

// Example benchmark usage:
//  go test codechallenge/gfg/GFG_DSA_IterativePower_test.go -count=1 -bench=.
// Example benchmark output:
// 	goos: linux
// 	goarch: amd64
// 	cpu: AMD Ryzen 9 PRO 6950HS with Radeon Graphics
// 	BenchmarkIterativePower-16                      348495673                3.562 ns/op
// 	BenchmarkIterativePower_Optimized-16            382027060                3.078 ns/op
// 	PASS
// 	ok      command-line-arguments  3.097s

func BenchmarkIterativePower(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IterativePower(10, 10)
	}
}

func BenchmarkIterativePower_Optimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IterativePower_Optimized(10, 10, 1000000007)
	}
}
