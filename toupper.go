package piscine

func ToUpper(s string) string {
	myString := []rune(s)
	for index := range myString {
		if myString[index] >= 'a' && myString[index] <= 'z' {
			myString[index] -= 32
		}
	}
	return string(myString)
}
