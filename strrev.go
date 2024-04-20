package piscine

func StrRev(s string) string {
	myBuffer := []rune(s)
	myRevString := make([]rune, len(myBuffer))
	for i := 0; i < len(myBuffer); i++ {
		myRevString[len(myBuffer)-1-i] = myBuffer[i]
	}
	s = string(myRevString)
	return s
}
