package printrange

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestPrintRangeNumbers(t *testing.T) {
	var tests = []struct {
		start    int
		end      int
		sep      string
		expected string
	}{
		{1, 5, "\n", "i = 1\ni = 2\ni = 3\ni = 4\ni = 5"},
		{2, 6, ", ", "i = 2, i = 3, i = 4, i = 5, i = 6"},
		{3, 3, " | ", "i = 3"},
	}

	for _, tt := range tests {
		t.Run(tt.sep, func(t *testing.T) {
			output := captureOutput(func() {
				PrintRangeNumbers(tt.start, tt.end, tt.sep)
			})

			// Trim the output to remove trailing newline
			output = strings.TrimSpace(output)
			if output != tt.expected {
				t.Errorf("PrintRangeNumbers(%d, %d, %q) = %q; expected %q", tt.start, tt.end, tt.sep, output, tt.expected)
			}
		})
	}
}

// Helper function to capture the output of a function
func captureOutput(f func()) string {
	// Redirect standard output to a buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	f()

	// Close the writer and restore standard output
	w.Close()
	os.Stdout = old

	// Read the captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
