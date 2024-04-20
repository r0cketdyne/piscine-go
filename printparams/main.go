package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		for index, arg := range os.Args {
			prName := make([]rune, 0)
			if index > 0 {
				prName = append(prName, []rune(arg)...)
				for rn := range prName {
					z01.PrintRune(prName[rn])
				}
				z01.PrintRune('\n')
			}
		}
	}
}
