package piscine

func Substring(myString, myStringToFind []rune, indexS, indexSTF int) int {
	if len(myStringToFind) == 0 {
		return 0
	}
	if myString[indexS] == myStringToFind[indexSTF] {
		if indexSTF == len(myStringToFind)-1 {
			return indexS - indexSTF
		}
		indexS++
		indexSTF++
		return Substring(myString, myStringToFind, indexS, indexSTF)
	}
	if indexS != len(myString)-1 {
		indexS++
		return Substring(myString, myStringToFind, indexS, indexSTF)
	}
	return -1
}

func Index(s string, toFind string) int {
	myString := []rune(s)
	myStringToFind := []rune(toFind)
	placeFound := -1
	if len(myString) > 0 {
		placeFound = Substring(myString, myStringToFind, 0, 0)
	}
	return placeFound
}
