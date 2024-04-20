package piscine

import "github.com/01-edu/z01"

func printDigits(digits []int) {
	for digit := range digits {
		z01.PrintRune(intToChar(digits[digit]))
	}
	if !(digits[0] == (10 - (len(digits)))) {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}

func generateCombinations(digits []int, start, index int) {
	if index == len(digits) {
		printDigits(digits)
		return
	}

	for i := start; i <= 9; i++ {
		digits[index] = i
		generateCombinations(digits, i+1, index+1)
	}
}

func PrintCombN(n int) {
	digits := make([]int, n)
	generateCombinations(digits, 0, 0)
	z01.PrintRune('\n')
}
