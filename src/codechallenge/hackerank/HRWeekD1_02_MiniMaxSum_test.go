package hackerank

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func MiniMaxSum(arr []int) []string {
	sum := int64(0)
	min := math.MaxInt
	max := math.MinInt

	for _, num := range arr {
		sum += int64(num)
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	results := []string{
		fmt.Sprintf("%d", sum-int64(max)),
		fmt.Sprintf("%d", sum-int64(min)),
	}

	return results
}

func TestMiniMaxSum(t *testing.T) {
	// Define test cases as a slice of structs
	testCases := []struct {
		input    []int
		expected []string
	}{
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []string{"10", "14"},
		},
		{
			input:    []int{256741038, 623958417, 467905213, 714532089, 938071625},
			expected: []string{"2063136757", "2744467344"},
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := MiniMaxSum(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
