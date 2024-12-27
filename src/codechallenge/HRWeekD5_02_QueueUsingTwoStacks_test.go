package codechallenge

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func QueueUsingTwoStacks(ary []string) []int {
	results := []int{}
	queue := []int{}

	for _, s := range ary {
		parts := strings.Fields(s)
		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]
		switch cmd {
		case "1": // Enqueue
			if len(parts) < 2 {
				continue
			}
			value, _ := strconv.Atoi(parts[1])
			queue = append(queue, value)
		case "2": // Dequeue
			if len(queue) > 0 {
				queue = queue[1:]
			}
		case "3": // Print
			if len(queue) > 0 {
				results = append(results, queue[0])
			}
		}
	}

	return results
}

func TestQueueUsingTwoStacks(t *testing.T) {
	// Define test cases as a slice of structs
	testCases := []struct {
		input    []string
		expected []int
	}{
		{
			[]string{
				"1 42", // Enqueue 42
				"2",    // Dequeue
				"1 14", // Enqueue 14
				"3",    // Print front element
				"1 28", // Enqueue 28
				"3",    // Print front element
				"1 60", // Enqueue 60
				"1 78", // Enqueue 78
				"2",    // Dequeue
				"2",    // Dequeue
			},
			[]int{14, 14},
		},
		{
			[]string{
				"1 10", // Enqueue 10
				"1 20", // Enqueue 20
				"3",    // Print front element
				"2",    // Dequeue
				"3",    // Print front element
			},
			[]int{10, 20},
		},
		{
			[]string{
				"1 5", // Enqueue 5
				"2",   // Dequeue
				"3",   // Print front element (queue is empty)
			},
			[]int{},
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := QueueUsingTwoStacks(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
