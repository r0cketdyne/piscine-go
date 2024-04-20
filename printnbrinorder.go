package piscine

import "github.com/01-edu/z01"

func fillArrayWithEvenAnAlmostOverFlowNumber(myNum []int, n int) []int {
	mod := n % 10
	div := n / 10
	if div < 0 {
		div *= -1
		mod *= -1
	}
	for div > 0 {
		myNum = append(myNum, div%10)
		div /= 10
	}
	myNum = append(myNum, mod%10)
	return myNum
}

func sortArray(myNum []int) {
	for i := range myNum {
		for j := i; j < len(myNum); j++ {
			if myNum[i] > myNum[j] {
				temp := myNum[i]
				myNum[i] = myNum[j]
				myNum[j] = temp
			}
		}
	}
}

func PrintingTheNumber(myNum []int) {
	for i := range myNum {
		z01.PrintRune(rune('0' + myNum[i]))
	}
}

func PrintNbrInOrder(n int) {
	myNum := make([]int, 0)
	myNum = fillArrayWithEvenAnAlmostOverFlowNumber(myNum, n)
	sortArray(myNum)
	PrintingTheNumber(myNum)
}
