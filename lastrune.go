package piscine

func LastRune(s string) rune {
	myRunes := []rune(s)
	return myRunes[len(myRunes)-1]
}
