package codechallenge

import (
	"container/heap"
	"fmt"
	"testing"
)

// PriorityQueue implements heap.Interface and holds items.
type PriorityQueue []int

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return pq[i] < pq[j] }

func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// JesseAndCookies processes a list of operations and returns the results
func JesseAndCookies(k int, A []int) int {
	steps := 0

	// Initialize a min-heap (priority queue)
	pq := &PriorityQueue{}
	heap.Init(pq)
	for _, a := range A {
		heap.Push(pq, a)
	}

	// Process the heap until all cookies have sweetness >= k or no more steps can be taken
	for pq.Len() >= 2 && (*pq)[0] < k {
		steps++
		first := heap.Pop(pq).(int)
		second := heap.Pop(pq).(int)
		newVal := first + 2*second
		heap.Push(pq, newVal)
	}

	// Check if the smallest cookie meets the requirement
	if pq.Len() > 0 && (*pq)[0] >= k {
		return steps
	}
	return -1
}

func TestJesseAndCookies(t *testing.T) {
	// Define test cases as a slice of structs
	testCases := []struct {
		k        int
		A        []int
		expected int
	}{
		{7, []int{1, 2, 3, 9, 10, 12}, 2},
		{90, []int{13, 47, 74, 12, 89, 74, 18, 38}, 5},
		{10, []int{1, 1, 1}, -1}, // Not possible to achieve sweetness >= 10
		{5, []int{5, 6, 7}, 0},   // Already meets the requirement
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("k=%d, A=%v", tc.k, tc.A), func(t *testing.T) {
			result := JesseAndCookies(tc.k, tc.A)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
