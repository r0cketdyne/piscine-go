package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		for index := range os.Args {
			prName := make([]rune, 0)
			if index < len(os.Args)-1 {
				prName = append(prName, []rune(os.Args[len(os.Args)-index-1])...)
				for rn := range prName {
					z01.PrintRune(prName[rn])
				}
				z01.PrintRune('\n')
			}
		}
	}
}
