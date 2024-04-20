package piscine

func StringToIntSlice(str string) []int {
	bytes := []rune(str)
	myints := []int{}
	if str == "" {
		return nil
	}
	for _, i := range bytes {
		myints = append(myints, int(i))
	}
	return myints
}
