package piscine

func Trim(s string) ([]string, int) {
	if len(s) > 0 {
		myStrings := []string{}
		counter := 0
		countsp := 0
		flag := false
		for i := range s {
			if rune(s[i]) == ' ' {
				if flag {
					countsp += 1
				}
				if counter != i {
					myStrings = append(myStrings, s[counter:i])
					counter = i + 1
				} else {
					counter += 1
				}
				flag = true
			} else {
				flag = false
			}
		}
		myStrings = append(myStrings, s[counter:])
		return myStrings, countsp
	}
	return nil, 0
}

func ShoppingSummaryCounter(str string) map[string]int {
	myMap := make(map[string]int)
	myStrs, countsp := Trim(str)
	if countsp > 0 {
		myMap[""] = countsp
	}
	if str == "" {
		myMap[""] = 1
		return myMap
	} else if str == " " {
		myMap[""] = 2
		return myMap
	}
	for _, i := range myStrs {
		_, f := myMap[i]
		if f {
			myMap[i] += 1
			continue
		}
		myMap[i] = 1
	}
	return myMap
}
