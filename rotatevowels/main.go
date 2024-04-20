package main

import (
	"os"

	"github.com/01-edu/z01"
)

var vowels = []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

func concatArgs(args []string) string {
	myString := ""
	for i := range args {
		if i+1 < len(args) {
			myString = myString + args[i] + " "
		} else {
			myString = myString + args[i]
		}
	}
	return myString
}

func VowelFound(r rune) bool {
	for i := range vowels {
		if vowels[i] == r {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]
		myString := concatArgs(args)
		myStringR := []rune(myString)
		vowelsFound := make([]rune, 0)
		for _, char := range myStringR {
			if VowelFound(char) {
				vowelsFound = append(vowelsFound, char)
			}
		}
		vC := len(vowelsFound) - 1
		for i, char := range myStringR {
			if VowelFound(char) {
				myStringR[i] = vowelsFound[vC]
				vC -= 1
			}
		}
		for _, char := range myStringR {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
