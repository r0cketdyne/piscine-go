package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	myBytesBuffer := []byte(s)
	for i := 0; i < len(myBytesBuffer); i++ {
		z01.PrintRune(rune(myBytesBuffer[i]))
	}
}
