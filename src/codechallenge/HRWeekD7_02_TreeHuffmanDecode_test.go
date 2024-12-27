package codechallenge

import (
	"sort"
	"strings"
	"testing"
)

// Huffman TreeNode
type HF_Node struct {
	Character string
	Frequency int
	Left      *HF_Node
	Right     *HF_Node
}

// Decoder
func TreeHuffmanDecode(root *HF_Node, encoded string) string {
	var sb strings.Builder
	current := root

	for _, ch := range encoded {
		switch ch {
		case '0':
			current = current.Left
		case '1':
			current = current.Right
		}

		if current.Left == nil && current.Right == nil {
			sb.WriteString(current.Character)
			current = root
		}
	}

	return sb.String()
}

// Encoder
func HF_Encoding(root *HF_Node, ch string, encoded string) (bool, string) {
	if root == nil {
		return false, encoded
	}

	if root.Character != "" && root.Character == ch {
		return true, encoded
	}

	if root.Left != nil {
		found, updated := HF_Encoding(root.Left, ch, encoded+"0")
		if found {
			return true, updated
		}
	}

	if root.Right != nil {
		found, updated := HF_Encoding(root.Right, ch, encoded+"1")
		if found {
			return true, updated
		}
	}

	return false, encoded
}

// Construct a Huffman Tree
func BuildHF_Tree(raw string) (*HF_Node, string) {
	// Computing each char's frequency
	freqMap := make(map[string]int)
	for _, ch := range raw {
		freqMap[string(ch)]++
	}

	// Sort by char's frequency
	type charFreq struct {
		Character string
		Frequency int
	}
	var items []charFreq
	for ch, freq := range freqMap {
		items = append(items, charFreq{Character: ch, Frequency: freq})
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Frequency < items[j].Frequency
	})

	// Start constructing Huffman Tree and TreeNodes
	leafTwo := items[:2]
	left := leafTwo[0]
	right := leafTwo[1]

	leaf := &HF_Node{
		Frequency: left.Frequency + right.Frequency,
		Left: &HF_Node{
			Frequency: left.Frequency,
			Character: left.Character,
		},
		Right: &HF_Node{
			Frequency: right.Frequency,
			Character: right.Character,
		},
	}

	root := leaf
	for _, item := range items[2:] {
		current := &HF_Node{
			Frequency: item.Frequency + root.Frequency,
			Left:      root,
			Right: &HF_Node{
				Frequency: item.Frequency,
				Character: item.Character,
			},
		}
		root = current
	}

	// Encoding raw string to encoded string
	var sb strings.Builder
	encodingCache := make(map[string]string)
	for _, ch := range raw {
		charStr := string(ch)
		if _, exists := encodingCache[charStr]; !exists {
			_, encoded := HF_Encoding(root, charStr, "")
			encodingCache[charStr] = encoded
		}
		sb.WriteString(encodingCache[charStr])
	}

	return root, sb.String()
}

func TestHuffman(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test case 1",
			input:    "ABRACADABRA",
			expected: "ABRACADABRA",
		},
		{
			name:     "Test case 2",
			input:    "HELLO",
			expected: "HELLO",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root, encoded := BuildHF_Tree(tt.input)
			result := TreeHuffmanDecode(root, encoded)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
