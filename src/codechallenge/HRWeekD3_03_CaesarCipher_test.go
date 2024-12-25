package codechallenge

import (
	"fmt"
	"testing"
	"unicode"
)

func CaesarCipher(s string, k int) string {
	k = k % 26
	runes := []rune(s)

	for i, origin := range runes {
		shifted := origin + rune(k)

		if unicode.IsLower(origin) {
			if shifted < 'a' {
				shifted += 26
			} else if shifted > 'z' {
				shifted -= 26
			}
		} else if unicode.IsUpper(origin) {
			if shifted < 'A' {
				shifted += 26
			} else if shifted > 'Z' {
				shifted -= 26
			}
		} else {
			shifted = origin
		}

		runes[i] = shifted
	}

	return string(runes)
}

func TestCaesarCipher(t *testing.T) {
	// Define test cases as a slice of structs
	testCases := []struct {
		input    string
		shift    int
		expected string
	}{
		{"middle-Outz", 2, "okffng-Qwvb"},
		{"Always-Look-on-the-Bright-Side-of-Life", 5, "Fqbfdx-Qttp-ts-ymj-Gwnlmy-Xnij-tk-Qnkj"},
		{"www.abc.xy", 87, "fff.jkl.gh"},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s, Shift: %d", tc.input, tc.shift), func(t *testing.T) {
			result := CaesarCipher(tc.input, tc.shift)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}
