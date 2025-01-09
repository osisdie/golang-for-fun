package hackerank

import (
	"fmt"
	"strings"
	"testing"
)

// SinglyLinkedListNode represents a node in the linked list
type SinglyLinkedListNode struct {
	data int
	next *SinglyLinkedListNode
}

// SinglyLinkedList represents a linked list
type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

// InsertNode inserts a new node at the end of the linked list
func (llist *SinglyLinkedList) InsertNode(nodeData int) {
	node := &SinglyLinkedListNode{data: nodeData}

	if llist.head == nil {
		llist.head = node
	} else {
		llist.tail.next = node
	}

	llist.tail = node
}

// MergeLists merges two sorted linked lists into one sorted linked list
func MergeLists(head1, head2 *SinglyLinkedListNode) *SinglyLinkedListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	// Dummy head to simplify the merging process
	dummy := &SinglyLinkedListNode{data: -1}
	currentMerged := dummy
	current1 := head1
	current2 := head2

	// Merge the two lists
	for current1 != nil && current2 != nil {
		if current1.data <= current2.data {
			currentMerged.next = current1
			current1 = current1.next
		} else {
			currentMerged.next = current2
			current2 = current2.next
		}

		currentMerged = currentMerged.next
	}

	if current1 != nil {
		currentMerged.next = current1
	} else if current2 != nil {
		currentMerged.next = current2
	}

	return dummy.next
}

// PrintSinglyLinkedList prints the linked list to a string
func PrintSinglyLinkedList(node *SinglyLinkedListNode, sep string) string {
	var result []string
	for node != nil {
		result = append(result, fmt.Sprintf("%d", node.data))
		node = node.next
	}

	return strings.Join(result, sep)
}

func TestMergeLists(t *testing.T) {
	// Define test cases as a slice of structs
	testCases := []struct {
		list1    []int
		list2    []int
		expected string
	}{
		{[]int{1, 3, 7}, []int{1, 2}, "1 1 2 3 7"},
		{[]int{1, 2, 3}, []int{3, 4}, "1 2 3 3 4"},
		{[]int{1, 5, 9}, []int{2, 4, 6}, "1 2 4 5 6 9"},
		{[]int{}, []int{1, 2, 3}, "1 2 3"},
		{[]int{1, 2, 3}, []int{}, "1 2 3"},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("List1: %v, List2: %v", tc.list1, tc.list2), func(t *testing.T) {
			// Create the first linked list
			llist1 := &SinglyLinkedList{}
			for _, data := range tc.list1 {
				llist1.InsertNode(data)
			}

			// Create the second linked list
			llist2 := &SinglyLinkedList{}
			for _, data := range tc.list2 {
				llist2.InsertNode(data)
			}

			// Merge the two lists
			mergedList := MergeLists(llist1.head, llist2.head)

			// Print the merged list to a string
			result := PrintSinglyLinkedList(mergedList, " ")

			// Compare the result with the expected output
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
