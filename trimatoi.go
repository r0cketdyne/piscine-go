package piscine

func trimming(myNum []rune) []rune {
	newNum := make([]rune, 0)
	for i := range myNum {
		if myNum[i] >= '0' && myNum[i] <= '9' {
			newNum = append(newNum, myNum[i])
		}
	}
	return newNum
}

func signOfNum(myNum []rune) rune {
	for i := range myNum {
		if myNum[i] >= '0' && myNum[i] <= '9' {
			break
		}
		if myNum[i] == '-' || myNum[i] == '+' {
			return myNum[i]
		}
	}
	return '+'
}

func PlaceInNum(n int) int {
	square := 1
	for i := 0; i < n; i++ {
		square *= 10
	}
	return square
}

func myChangingToInt(sign rune, myNum []rune) int {
	myInt := 0
	if len(myNum) != 0 {
		for i := range myNum {
			myInt += int(myNum[len(myNum)-i-1]-48) * PlaceInNum(i)
		}
		if sign == '-' {
			myInt *= -1
		}
	}
	return myInt
}

func TrimAtoi(s string) int {
	myNum := []rune(s)
	sign := signOfNum(myNum)
	myNum = trimming(myNum)
	myInt := myChangingToInt(sign, myNum)
	return myInt
}
