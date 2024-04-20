package piscine

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := range base {
		if base[i] == '-' || base[i] == '+' {
			return false
		}
	}
	for i := range base {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
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
