package piscine

func ValidBase2(Base []rune) bool {
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

func AtoiBase2(s string, base string) int {
	if !ValidBase([]rune(base)) {
		return 0
	}

	baseMap := make(map[rune]int)
	for i, char := range base {
		baseMap[char] = i
	}

	result := 0
	multiplier := 1
	for i := len(s) - 1; i >= 0; i-- {
		digit := baseMap[rune(s[i])]
		result += digit * multiplier
		multiplier *= len(base)
	}

	return result
}

func numberPrintingAsPerTheBase2(nbr int, myBase []rune) []rune {
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
	for i := range myNumber {
		for j := i + 1; j < len(myNumber); j++ {
			myNumber[i], myNumber[j] = myNumber[j], myNumber[i]
		}
	}
	return myNumber
}

func PrintNbrBase2(nbr int, base string) string {
	myBase := []rune(base)
	if ValidBase(myBase) {
		myNumber := numberPrintingAsPerTheBase2(nbr, myBase)
		return string(myNumber)
	}
	return ""
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	nbr = PrintNbrBase2(AtoiBase2(nbr, baseFrom), baseTo)
	return nbr
}
