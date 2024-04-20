package piscine

func Substring2(myString, myStringToFind []rune, indexS, indexSTF int) int {
	if len(myStringToFind) == 0 {
		return 0
	}
	if myString[indexS] == myStringToFind[indexSTF] {
		if indexSTF == len(myStringToFind)-1 {
			return indexS - indexSTF
		}
		indexS++
		indexSTF++
		return Substring2(myString, myStringToFind, indexS, indexSTF)
	}
	if indexS != len(myString)-1 {
		indexS++
		return Substring2(myString, myStringToFind, indexS, indexSTF)
	}
	return -1
}

func Index2(s string, toFind string) int {
	myString := []rune(s)
	myStringToFind := []rune(toFind)
	placeFound := -1
	if len(myString) > 0 {
		placeFound = Substring2(myString, myStringToFind, 0, 0)
	}
	return placeFound
}

func Split(s, sep string) []string {
	if len(s) > 0 {
		result := []string{}
		for len(s) > 0 {
			counter := Index2(s, sep)
			if counter != -1 {
				result = append(result, s[0:counter])
				s = s[len(sep)+counter:]
			} else {
				result = append(result, s[0:])
				s = ""
			}
		}
		return result
	}
	return nil
}
