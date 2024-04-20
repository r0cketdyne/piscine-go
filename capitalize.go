package piscine

func Capitalize(s string) string {
	myString := []rune(s)
	flagCap := true
	flagDeCap := false
	for i := range myString {
		if (myString[i] < 'A' || myString[i] > 'Z') && (myString[i] < 'a' || myString[i] > 'z') && (myString[i] < '0' || myString[i] > '9') {
			flagCap = true
			flagDeCap = false
		}
		if myString[i] >= '0' && myString[i] <= '9' {
			flagCap = false
			flagDeCap = true
		}
		if myString[i] >= 'A' && myString[i] <= 'Z' {
			if flagDeCap {
				myString[i] += 32
			}
			flagCap = false
			flagDeCap = true
		}
		if myString[i] >= 'a' && myString[i] <= 'z' {
			if flagCap {
				myString[i] -= 32
				flagCap = false
			}
			flagDeCap = true
		}
	}
	return string(myString)
}
