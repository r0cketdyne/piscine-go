package piscine

func SplitWhiteSpaces(s string) []string {
	if len(s) > 0 {
		myStrings := []string{}
		counter := 0
		for i := range s {
			if rune(s[i]) == '\n' || rune(s[i]) == ' ' || rune(s[i]) == '	' {
				if counter != i {
					myStrings = append(myStrings, s[counter:i])
					counter = i + 1
				} else {
					counter += 1
				}
			}
		}
		myStrings = append(myStrings, s[counter:])
		return myStrings
	}
	return nil
}
