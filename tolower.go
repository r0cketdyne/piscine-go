package piscine

func ToLower(s string) string {
	myString := []rune(s)
	for index := range myString {
		if myString[index] >= 'A' && myString[index] <= 'Z' {
			myString[index] += 32
		}
	}
	return string(myString)
}
