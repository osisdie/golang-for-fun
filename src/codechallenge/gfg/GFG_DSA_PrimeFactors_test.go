package gfg

import (
	"fmt"
	"reflect"
	"testing"
)

func PrimeFactors(n int) []int {
	res := []int{}

	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			res = append(res, i)
			n /= i
		}
	}

	if n > 1 {
		res = append(res, n)
	}

	return res
}

func PrimeFactors_Optimized(n int) []int {
	res := []int{}

	// 2
	for n%2 == 0 {
		res = append(res, 2)
		n /= 2
	}
	// 3
	for n%3 == 0 {
		res = append(res, 3)
		n /= 3
	}

	for i := 5; i*i <= n; i += 6 {
		for n%i == 0 {
			res = append(res, i)
			n /= i
		}

		for n%(i+2) == 0 {
			res = append(res, (i + 2))
			n /= (i + 2)
		}
	}

	if n > 3 {
		res = append(res, n)
	}

	return res
}

func TestPrimeFactors(t *testing.T) {
	testCases := []struct {
		input    int
		expected []int
	}{
		{
			input:    12,
			expected: []int{2, 2, 3},
		},
		{
			input:    13,
			expected: []int{13},
		},
		{
			input:    315,
			expected: []int{3, 3, 5, 7},
		},
		{
			input:    84,
			expected: []int{2, 2, 3, 7},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := PrimeFactors(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}

			result = PrimeFactors_Optimized(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

// go test codechallenge/gfg/GFG_DSA_PrimeFactors_test.go -count=1 -v
