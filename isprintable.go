package piscine

func IsPrintable(s string) bool {
	myString := []rune(s)
	for index := range myString {
		if myString[index] < 32 || myString[index] > 126 {
			return false
		}
	}
	return true
}
