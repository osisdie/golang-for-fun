package main

import (
	"example.com/app/printrange"
    "fmt"
    "runtime"
)

func main() {
	// Hello, Go go1.23.4
	fmt.Println("Hello, Go", runtime.Version())

	fmt.Println("Printing numbers with default separator:")
    printrange.PrintRangeNumbers(1, 5, "\n")

    fmt.Println("\nPrinting numbers with custom separator:")
    printrange.PrintRangeNumbers(2, 6, ", ")
}