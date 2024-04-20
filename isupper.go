package piscine

func IsUpper(s string) bool {
	myString := []rune(s)
	for index := range myString {
		if myString[index] < 'A' || myString[index] > 'Z' {
			return false
		}
	}
	return true
}
