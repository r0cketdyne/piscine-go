package piscine

func IsLower(s string) bool {
	myString := []rune(s)
	for index := range myString {
		if myString[index] < 'a' || myString[index] > 'z' {
			return false
		}
	}
	return true
}
