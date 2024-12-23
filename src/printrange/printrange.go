package printrange

import (
    "fmt"
    "strings"
)

// PrintRangeNumbers prints numbers from start to end, separated by the given separator.
func PrintRangeNumbers(start int, end int, sep string) {
    var numbers []string
    for i := start; i <= end; i++ {
        numbers = append(numbers, fmt.Sprintf("i = %d", i))
    }
    fmt.Println(strings.Join(numbers, sep))
}