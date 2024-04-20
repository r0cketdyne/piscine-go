package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i, s := range a {
		for _, char := range s {
			z01.PrintRune(char)
		}
		if i != len(a)-1 {
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')
}
