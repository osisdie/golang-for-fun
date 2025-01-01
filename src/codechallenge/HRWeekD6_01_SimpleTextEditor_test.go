package codechallenge

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

// SimpleTextEditor processes a list of operations and returns the results
func SimpleTextEditor(operations []string) []string {
	undoStack := []string{}
	results := []string{}
	text := ""

	for _, ops := range operations {
		parts := strings.Fields(ops)
		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]
		switch cmd {
		case "1": // Append
			undoStack = append(undoStack, text) // Save previous text
			text = xAppend(text, parts[1])
		case "2": // Delete
			lastChars, _ := strconv.Atoi(parts[1])
			updated, before := xDelete(text, lastChars)
			if updated != before {
				undoStack = append(undoStack, before) // Save previous text
			}
			text = updated
		case "3": // Print
			printIdx, _ := strconv.Atoi(parts[1])
			if printIdx > 0 && printIdx <= len(text) {
				results = append(results, string(text[printIdx-1]))
			}
		case "4": // Undo
			if len(undoStack) > 0 {
				text, undoStack = undoStack[len(undoStack)-1], undoStack[:len(undoStack)-1]
			}
		}
	}

	return results
}

// xAppend appends a string to the existing text
func xAppend(existing, appending string) string {
	return existing + appending
}

// xDelete deletes the last `lastChars` characters from the existing text
func xDelete(existing string, lastChars int) (string, string) {
	if len(existing) >= lastChars {
		return existing[:len(existing)-lastChars], existing
	}
	return existing, existing
}

func TestSimpleTextEditor(t *testing.T) {
	// Define test cases as a slice of structs
	testCases := []struct {
		input    []string
		expected []string
	}{
		{
			[]string{
				"1 abcde", // abcde
				"1 fg",    // abcdefg
				"3 6",     // abcdefg, print f
				"2 5",     // ab
				"4",       // abcdefg
				"3 7",     // abcdefg, print g
				"4",       // ab -> abcdefg
				"3 4",     // abcdefg, print d
			},
			[]string{"f", "g", "d"},
		},
		{
			[]string{
				"1 abc", // abc
				"3 3",   // abc, print c
				"2 3",   // empty
				"1 xy",  // xy
				"3 2",   // xy, print y
				"4",     // xy -> empty
				"4",     // empty -> abc
				"3 1",   // abc, print a
			},
			[]string{"c", "y", "a"},
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := SimpleTextEditor(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
