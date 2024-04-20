package piscine

func BasicAtoi(s string) int {
	num := 0
	for i := range s {
		digit := int(s[i] - '0')
		num = num*10 + digit
	}
	return num
}
