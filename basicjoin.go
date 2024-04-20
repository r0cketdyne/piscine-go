package piscine

func BasicJoin(elems []string) string {
	myString := make([]rune, 0)
	for i := range elems {
		element := []rune(elems[i])
		for j := range element {
			myString = append(myString, element[j])
		}
	}
	return string(myString)
}
