package codechallenge

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func FindZigZagSequence(arr []int) []int {
	// Sort the input array
	sort.Ints(arr)

	len := len(arr)
	middleIdx := len / 2
	leftIdx := 0
	rightIdx := len - 1

	// Create a result array
	result := make([]int, len)
	result[middleIdx] = arr[len-1] // Place the maximum value in the middle

	// Fill the result array
	for i := 0; i < len-1; i++ {
		if i < middleIdx {
			result[leftIdx] = arr[i]
			leftIdx++
		} else {
			result[rightIdx] = arr[i]
			rightIdx--
		}
	}

	return result
}

func TestFindZigZagSequence(t *testing.T) {
	// Define test cases as a slice of structs
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{2, 3, 5, 1, 4},
			expected: []int{1, 2, 5, 4, 3},
		},
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			expected: []int{1, 2, 3, 7, 6, 5, 4},
		},
	}

	// Add test cases from files
	testCases = append(testCases, readTestCasesFromFile(t)...)

	// Run tests
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %v", tc.input), func(t *testing.T) {
			result := FindZigZagSequence(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

// Helper function to read test cases from files
func readTestCasesFromFile(t *testing.T) []struct {
	input    []int
	expected []int
} {
	var testCases []struct {
		input    []int
		expected []int
	}

	// Get the pwd directory
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}

	// Read input and output files
	inputFile := filepath.Join(pwd,
		"test_data/HRWeekD3_01_FindZigZagSequence_test/1/input.txt")
	outputFile := filepath.Join(pwd,
		"test_data/HRWeekD3_01_FindZigZagSequence_test/1/output.txt")

	inputLines, err := readLines(inputFile)
	if err != nil {
		t.Fatalf("Failed to read input file: %v", err)
	}

	outputLines, err := readLines(outputFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	// Parse the number of test cases
	inputN, err := strconv.Atoi(inputLines[0])
	if err != nil {
		t.Fatalf("Failed to parse number of test cases: %v", err)
	}

	// Parse input and output arrays
	for i := 0; i < inputN; i++ {
		inputAry := parseArray(inputLines[(i+1)*2])
		outputAry := parseArray(outputLines[i])
		testCases = append(testCases, struct {
			input    []int
			expected []int
		}{inputAry, outputAry})
	}

	return testCases
}

// Helper function to read lines from a file
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// Helper function to parse a space-separated string into an array of integers
func parseArray(line string) []int {
	var result []int
	for _, s := range strings.Split(line, " ") {
		num, _ := strconv.Atoi(s)
		result = append(result, num)
	}
	return result
}
