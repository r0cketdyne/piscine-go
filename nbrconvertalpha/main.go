package main

import (
	"os"

	"github.com/01-edu/z01"
)

func atoi(s string) int {
	var result int
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		result = result*10 + int(s[i]-'0')
	}
	return result
}

func getLetter(n int, upper bool) string {
	// Offset to get the corresponding ASCII value for 'a'
	offset := 96

	// If upper case flag is provided, adjust offset for upper case letters
	if upper {
		offset = 64
	}

	// Calculate ASCII value for the corresponding letter
	asciiValue := offset + n

	// Convert ASCII value to string
	return string(asciiValue)
}

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]

		// Check if the --upper flag is provided
		isUpper := false
		if len(args) > 0 && args[0] == "--upper" {
			isUpper = true
			args = args[1:]
		}
		nums := make([]int, 0)
		for _, arg := range args {
			nums = append(nums, atoi(arg))
		}
		for _, n := range nums {
			if n < 1 || n > 26 {
				// If not, print a space
				z01.PrintRune(' ')
				continue
			}
			z01.PrintRune(rune(getLetter(n, isUpper)[0]))
		}
		z01.PrintRune('\n')
	}
}
