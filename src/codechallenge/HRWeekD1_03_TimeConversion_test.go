package codechallenge

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TimeConversion(s string) string {
	// Split the input string into components
	parts := strings.Split(s, ":")
	hh, _ := strconv.Atoi(parts[0])
	group := s[len(s)-2:] // Extract "AM" or "PM"

	// Handle AM and PM cases
	switch group {
	case "AM":
		if hh == 12 {
			hh = 0
		}
	case "PM":
		if hh != 12 {
			hh += 12
		}
	default:
		panic("Invalid time format")
	}

	// Format the hour and combine with the rest of the time
	hhStr := fmt.Sprintf("%02d", hh)
	result := hhStr + s[2:len(s)-2]

	return result
}

func TestTimeConversion(t *testing.T) {
	// Define test cases as a slice of structs
	testCases := []struct {
		input    string
		expected string
	}{
		{"07:05:45PM", "19:05:45"},
		{"12:40:22AM", "00:40:22"},
		{"12:45:54PM", "12:45:54"},
		{"12:00:00AM", "00:00:00"},
		{"12:00:00PM", "12:00:00"},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := TimeConversion(tc.input)
			if result != tc.expected {
				t.Errorf("Input: %s, Expected: %s, Got: %s", tc.input, tc.expected, result)
			}
		})
	}
}
