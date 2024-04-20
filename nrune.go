package piscine

func NRune(s string, n int) rune {
	myRunes := []rune(s)
	RuneToPrint := rune('\x00')
	if n <= len(myRunes) && n > 0 {
		RuneToPrint = myRunes[n-1]
	}
	return RuneToPrint
}
