package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := point{}
	setPoint(&points)
	a := "x = "
	b := ", y = "
	for _, char := range a {
		z01.PrintRune(char)
	}
	PrintNbr(points.x)
	for _, char := range b {
		z01.PrintRune(char)
	}
	PrintNbr(points.y)
	z01.PrintRune('\n')
}

func PrintNbr(n int) {
	temp := n % 10
	n /= 10
	if n < 0 {
		n = -n
		temp = -temp
		z01.PrintRune(rune('-'))
	}
	if n > 0 {
		PrintNbr(n)
	}
	z01.PrintRune(rune(temp + '0'))
	return
}
