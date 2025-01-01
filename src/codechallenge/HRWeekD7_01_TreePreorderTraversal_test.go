package codechallenge

import (
	"strconv"
	"strings"
	"testing"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(data int) *Node {
	return &Node{data: data, left: nil, right: nil}
}

func Insert(node *Node, data int) *Node {
	if node == nil {
		return NewNode(data)
	}

	if data <= node.data {
		node.left = Insert(node.left, data)
	} else {
		node.right = Insert(node.right, data)
	}

	return node
}

func PreOrderUtil(node *Node) string {
	results := []int{}
	TraverseDFSToList(node, &results)
	strResults := make([]string, len(results))
	for i, v := range results {
		strResults[i] = strconv.Itoa(v)
	}
	return strings.Join(strResults, " ")
}

func TraverseDFSToList(node *Node, outList *[]int) {
	if node == nil {
		return
	}

	*outList = append(*outList, node.data)

	TraverseDFSToList(node.left, outList)
	TraverseDFSToList(node.right, outList)
}

func TestPreOrderUtil(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test case 1",
			input:    "1 2 5 3 6 4",
			expected: "1 2 5 3 4 6",
		},
		{
			name:     "Test case 2",
			input:    "1 14 3 7 4 5 15 6 13 10 11 2 12 8 9",
			expected: "1 14 3 2 7 4 5 6 13 10 8 9 11 12 15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var root *Node
			inputAry := strings.Split(tt.input, " ")
			for _, str := range inputAry {
				num, _ := strconv.Atoi(str)
				root = Insert(root, num)
			}

			result := PreOrderUtil(root)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
