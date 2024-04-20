package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 0; i < 99; i++ {
		for j := i + 1; j < 100; j++ {
			if j > i {
				z01.PrintRune(intToChar(i / 10))
				z01.PrintRune(intToChar(i % 10))
				z01.PrintRune(' ')
				z01.PrintRune(intToChar(j / 10))
				z01.PrintRune(intToChar(j % 10))
				if !(i == 98) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				} else {
					z01.PrintRune('\n')
				}
			}
		}
	}
}
