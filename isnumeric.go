package piscine

func IsNumeric(s string) bool {
	myString := []rune(s)
	for index := range myString {
		if myString[index] < '0' || myString[index] > '9' {
			return false
		}
	}
	return true
}
