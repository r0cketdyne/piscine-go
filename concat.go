package piscine

func Concat(str1 string, str2 string) string {
	myString1 := []rune(str1)
	myString2 := []rune(str2)
	myWholeString := make([]rune, 0)
	for i := 0; i < len(myString1); i++ {
		myWholeString = append(myWholeString, myString1[i])
	}
	for i := 0; i < len(myString2); i++ {
		myWholeString = append(myWholeString, myString2[i])
	}
	return string(myWholeString)
}
