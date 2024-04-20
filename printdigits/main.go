package main

import (
	"github.com/01-edu/z01"
)

func intToChar(num int) rune {
	return rune('0' + num)
}

func main() {
	for i := 0; i < 10; i++ {
		char := intToChar(i)
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
