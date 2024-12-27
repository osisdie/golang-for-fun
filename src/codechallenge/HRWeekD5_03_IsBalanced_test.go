package codechallenge

import (
	"fmt"
	"testing"
)

func IsBalanced(s string) string {
	if len(s) == 0 {
		return "YES"
	}

	stack := []rune{}
	brackets := map[rune]bool{
		'{': true,
		'[': true,
		'(': true,
	}

	for _, ch := range s {
		if brackets[ch] {
			// Push opening brackets onto the stack
			stack = append(stack, ch)
		} else {
			// If the stack is empty, it's unbalanced
			if len(stack) == 0 {
				return "NO"
			}

			// Pop the top element from the stack
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Check if the popped bracket matches the current closing bracket
			if (ch == ')' && top != '(') ||
				(ch == '}' && top != '{') ||
				(ch == ']' && top != '[') {
				return "NO"
			}
		}
	}

	// If the stack is empty, all brackets are balanced
	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func TestIsBalanced(t *testing.T) {
	// Define test cases as a slice of structs
	testCases := []struct {
		input    string
		expected string
	}{
		{"{[()]}", "YES"},
		{"{[(])}", "NO"},
		{"{{[[(())]]}}", "YES"},
		{"{{}(", "NO"},
		{"{{([])}}", "YES"},
		{"{{)[](}}", "NO"},
		{"{(([])[])[]}", "YES"},
		{"{(([])[])[]]}", "NO"},
		{"{(([])[])[]}[]", "YES"},
		{"}][}}(}][))]", "NO"},
		{"[](){()}", "YES"},
		{"()", "YES"},
		{"({}([][]))[]()", "YES"},
		{"{)[](}]}]}))}(())(", "NO"},
		{"([[)", "NO"}, // Corner case
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := IsBalanced(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
