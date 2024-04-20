package piscine

func Join(strs []string, sep string) string {
	myString := make([]rune, 0)
	Separator := []rune(sep)
	for i := range strs {
		element := []rune(strs[i])
		myString = append(myString, element...)
		if i != len(strs)-1 {
			myString = append(myString, Separator...)
		}
	}
	return string(myString)
}
