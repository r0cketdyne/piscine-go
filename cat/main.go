package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]
		for _, filename := range args {
			data, err := os.ReadFile(filename)
			if err != nil {
				myString := "ERROR: open " + filename + ": no such file or directory"
				for _, char := range myString {
					z01.PrintRune(char)
				}
				z01.PrintRune('\n')
				os.Exit(1)
			}
			for _, char := range data {
				z01.PrintRune(rune(char))
			}
		}
	} else {
		// no args, exiting?
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			for _, char := range "Error reading from stdin: " {
				z01.PrintRune(char)
			}
			for _, char := range err.Error() {
				z01.PrintRune(char)
			}
			z01.PrintRune('\n')
			os.Exit(1)
		}
	}
}
