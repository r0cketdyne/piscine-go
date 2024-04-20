package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 9; j++ {
			for k := j + 1; k < 10; k++ {
				if k > j && j > i {
					z01.PrintRune(intToChar(i))
					z01.PrintRune(intToChar(j))
					z01.PrintRune(intToChar(k))
					if !(i == 7) {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					} else {
						z01.PrintRune('\n')
					}
				}
			}
		}
	}
}
