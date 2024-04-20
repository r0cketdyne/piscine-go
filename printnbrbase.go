package piscine

import "github.com/01-edu/z01"

func ValidBase(Base []rune) bool {
	if len(Base) < 2 {
		return false
	}
	for i := range Base {
		if Base[i] == '-' || Base[i] == '+' {
			return false
		}
	}
	for i := range Base {
		for j := i + 1; j < len(Base); j++ {
			if Base[i] == Base[j] {
				return false
			}
		}
	}
	return true
}

func numberPrintingAsPerTheBase(nbr int, myBase []rune) []rune {
	IsNegative := false
	myNumber := make([]rune, 0)
	mod := nbr % len(myBase)
	nbr /= len(myBase)
	if nbr < 0 {
		IsNegative = true
		nbr *= -1
		mod *= -1
	}
	myNumber = append(myNumber, myBase[mod])
	for nbr > 0 {
		remainder := nbr % len(myBase)
		myNumber = append(myNumber, myBase[remainder])
		nbr /= len(myBase)
	}
	if IsNegative {
		myNumber = append(myNumber, '-')
	}
	return myNumber
}

func PrintRunes(runes []rune) {
	for r := range runes {
		z01.PrintRune(runes[len(runes)-r-1])
	}
}

func PrintNbrBase(nbr int, base string) {
	myBase := []rune(base)
	myNumber := make([]rune, 0)
	if ValidBase(myBase) {
		myNumber = numberPrintingAsPerTheBase(nbr, myBase)
	} else {
		myNumber = append(myNumber, 'V')
		myNumber = append(myNumber, 'N')
	}
	PrintRunes(myNumber)
}
