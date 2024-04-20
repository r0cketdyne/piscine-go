package piscine

func BasicAtoi2(s string) int {
	num := 0
	for i := range s {
		digit := int(s[i] - '0')
		if digit < 0 || digit > 9 {
			return 0
		}
		num = num*10 + digit
	}
	return num
}
