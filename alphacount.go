package piscine

func AlphaCount(s string) int {
	myRunes := []rune(s)
	counter := 0
	for i := 0; i < len(myRunes); i++ {
		if (myRunes[i] >= 'a' && myRunes[i] <= 'z') || (myRunes[i] >= 'A' && myRunes[i] <= 'Z') {
			counter += 1
		}
	}
	return counter
}
