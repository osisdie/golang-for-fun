package gfg

import (
	"fmt"
	"testing"
)

func ComputingPow(x, n int) int {
	if n == 0 {
		return 1
	}

	res := 1
	for i := 0; i < n; i++ {
		res *= x
	}

	return res
}

func ComputingPow_Recursive(x, n int) int {
	if n == 0 {
		return 1
	}

	if n%2 == 0 {
		return ComputingPow_Recursive(x, n/2) * ComputingPow_Recursive(x, n/2)
	} else {
		return ComputingPow_Recursive(x, n-1) * x
	}
}

func ComputingPow_Optimized(x, n int) int {
	if n == 0 {
		return 1
	}

	temp := ComputingPow_Optimized(x, n/2)
	temp = temp * temp

	if n%2 == 0 {
		return temp
	} else {
		return temp * x
	}
}

func TestComputingPow(t *testing.T) {
	testCases := []struct {
		x, n     int
		expected int
	}{
		{
			x:        2,
			n:        3,
			expected: 8,
		},
		{
			x:        3,
			n:        4,
			expected: 81,
		},
		{
			x:        5,
			n:        0,
			expected: 1,
		},
		{
			x:        5,
			n:        1,
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v, %v", tc.x, tc.n), func(t *testing.T) {
			result := ComputingPow(tc.x, tc.n)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}

			result = ComputingPow_Recursive(tc.x, tc.n)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}

			result = ComputingPow_Optimized(tc.x, tc.n)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

// go test -run ^Test.*ComputingPow$ example.com/app/codechallenge/gfg -count=1 -v

// Example benchmark usage:
//  go test codechallenge/gfg/GFG_DSA_ComputingPow_test.go -count=1 -bench=.
// Example benchmark output:
// 	goos: linux
// 	goarch: amd64
// 	cpu: AMD Ryzen 9 PRO 6950HS with Radeon Graphics
// 	BenchmarkComputingPow-16                346950804                3.584 ns/op
// 	BenchmarkComputingPow_Recursive-16      31205078                37.49 ns/op
// 	PASS
// 	ok      command-line-arguments  2.814s

func BenchmarkComputingPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ComputingPow(10, 10)
	}
}

func BenchmarkComputingPow_Recursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ComputingPow_Recursive(10, 10)
	}
}

func BenchmarkComputingPow_Optimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ComputingPow_Optimized(10, 10)
	}
}
