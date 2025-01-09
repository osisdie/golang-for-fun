package gfg

import (
	"fmt"
	"reflect"
	"testing"
)

func _isPrime(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i < n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func AllPrimes(n int) []int {
	res := []int{}

	for i := 2; i <= n; i++ {
		if _isPrime(i) {
			res = append(res, i)
		}
	}

	return res
}

// Sieve of Eratosthenes
func AllPrimes_Optimized(n int) []int {
	var prime_list = make([]bool, n+1)
	res := []int{}

	for i := range prime_list {
		prime_list[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if prime_list[i] {
			for j := 2 * i; j <= n; j += i {
				prime_list[j] = false
			}
		}
	}

	for i := 2; i <= n; i++ {
		if prime_list[i] {
			res = append(res, i)
		}
	}

	return res
}

// O(nloglogN)
func AllPrimes_EvenOptimized(n int) []int {
	var composite_list = make([]bool, n+1)
	res := []int{}

	for i := 2; i <= n; i++ {
		if !composite_list[i] {
			res = append(res, i)
			for j := i * i; j <= n; j += i {
				composite_list[j] = true
			}
		}
	}

	return res
}

func TestAllPrimes(t *testing.T) {
	testCases := []struct {
		input    int
		expected []int
	}{
		{
			input:    10,
			expected: []int{2, 3, 5, 7},
		},
		{
			input:    23,
			expected: []int{2, 3, 5, 7, 11, 13, 17, 19, 23},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := AllPrimes(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}

			result = AllPrimes_Optimized(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}

			result = AllPrimes_EvenOptimized(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

// go test codechallenge/gfg/GFG_DSA_AllPrimes_test.go -count=1 -v

// Example benchmark usage:
//  go test codechallenge/gfg/GFG_DSA_AllPrimes_test.go -count=1 -bench=.
// Example benchmark output:
// 	goos: linux
// 	goarch: amd64
// 	cpu: AMD Ryzen 9 PRO 6950HS with Radeon Graphics
// 	BenchmarkAllPrimes-16                    2484266               483.0 ns/op
// 	BenchmarkAllPrimes_Optimized-16          2692304               440.5 ns/op
// 	BenchmarkAllPrimes_EvenOptimized-16      3014186               386.0 ns/op
// 	PASS
// 	ok      command-line-arguments  4.909s

func BenchmarkAllPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllPrimes(100)
	}
}

func BenchmarkAllPrimes_Optimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllPrimes_Optimized(100)
	}
}

func BenchmarkAllPrimes_EvenOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AllPrimes_EvenOptimized(100)
	}
}
