package piscine

func JumpOver(str string) string {
	myString := ""
	if len(str) > 2 {
		endingP := len(str) % 3
		for i := 0; i < len(str)-endingP; i++ {
			if i%3 == 2 {
				myString = myString + string(str[i])
			}
		}
	}
	myString = myString + "\n"
	return myString
}
