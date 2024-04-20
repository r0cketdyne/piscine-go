package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func findinsert(args []string) (int, string) {
	insert := -1
	addingstring := ""
	for index, s := range args {
		if len(s) > 9 {
			if s[0:9] == "--insert=" {
				addingstring = s[9:]
				insert = index
			}
		}
		if len(s) > 3 {
			if s[0:3] == "-i=" {
				addingstring = s[3:]
				insert = index
			}
		}
	}
	return insert, addingstring
}

func findorder(args []string) int {
	order := -1
	for index, s := range args {
		if len(s) == 7 {
			if s[0:7] == "--order" {
				order = index
			}
		}
		if len(s) == 2 {
			if s[0:2] == "-o" {
				order = index
			}
		}
	}
	return order
}

func main() {
	args := []string(os.Args)
	h := false
	if len(args) > 1 {
		args = args[1:]
		if args[0] == "-h" || args[0] == "--help" {
			h = true
		} else {
			insert, addingstring := findinsert(args)
			order := findorder(args)
			myString := make([]rune, 0)
			for index, s := range args {
				if insert != index && order != index {
					s2 := []rune(s)
					myString = append(myString, s2...)
				}
			}
			if insert > -1 {
				s := []rune(addingstring)
				myString = append(myString, s...)
			}
			if order > -1 {
				for i := range myString {
					for j := i + 1; j < len(myString); j++ {
						if myString[i] > myString[j] {
							myString[i], myString[j] = myString[j], myString[i]
						}
					}
				}
			}
			for _, char := range myString {
				z01.PrintRune(char)
			}
			z01.PrintRune('\n')
		}
	} else {
		h = true
	}
	if h {
		fmt.Println("--insert")
		fmt.Println("  -i")
		fmt.Println("	 This flag inserts the string into the string passed as argument.")
		fmt.Println("--order")
		fmt.Println("  -o")
		fmt.Println("	 This flag will behave like a boolean, if it is called it will order the argument.")
	}
}
