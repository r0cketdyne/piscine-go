package piscine

func IsAlpha(s string) bool {
	myString := []rune(s)
	for index := range myString {
		if !((myString[index] >= 'a' && myString[index] <= 'z') || (myString[index] >= 'A' && myString[index] <= 'Z') || (myString[index] >= '0' && myString[index] <= '9')) {
			return false
		}
	}
	return true
}
