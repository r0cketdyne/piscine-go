package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	prName := make([]rune, 0)
	// Check if any command line arguments were provided
	if len(os.Args) > 1 {
		// Print the program name
		prName = append(prName, []rune(os.Args[0])...)
		for rn := range prName {
			if rn >= 2 {
				z01.PrintRune(prName[rn])
			}
		}
		z01.PrintRune('\n')
		return
	}
	// Print the program name if no command line arguments were provided
	prName = append(prName, []rune(os.Args[0])...)
	for rn := range prName {
		if rn >= 2 {
			z01.PrintRune(prName[rn])
		}
	}
	z01.PrintRune('\n')
}
