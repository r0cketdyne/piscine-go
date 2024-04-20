package piscine

import "github.com/01-edu/z01"

func PrintNbr(num int) {
	n1 := num / 10
	n2 := num % 10
	n3 := n1 % 10
	n1 = n1 / 10
	n4 := n1 % 10
	n1 = n1 / 10
	n5 := n1 % 10
	n1 = n1 / 10
	n6 := n1 % 10
	n1 = n1 / 10
	if num < 0 {
		z01.PrintRune('-')
		n1 = -n1
		n2 = -n2
		n3 = -n3
		n4 = -n4
		n5 = -n5
		n6 = -n6
	}

	// Handle the special case of zero
	if num == 0 {
		z01.PrintRune('0')
		return
	}

	// Create a buffer to hold the digits of the number
	digits := make([]int, 0, 20) // Assuming max number of digits is 20
	digits = append(digits, n2)
	digits = append(digits, n3)
	digits = append(digits, n4)
	digits = append(digits, n5)
	digits = append(digits, n6)
	// Extract each digit of the number
	for n1 > 0 {
		digits = append(digits, n1%10)
		n1 /= 10
	}

	// Print the digits in reverse order
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(intToChar(digits[i]))
	}
}
