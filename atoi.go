package piscine

func Atoi(s string) int {
	num := 0
	sign := 1
	for i := range s {
		if i == 0 && (s[i] == '-' || s[i] == '+') {
			if s[i] == '-' {
				sign = -1
			}
			continue
		}
		digit := int(s[i] - '0')
		if digit < 0 || digit > 9 {
			return 0
		}
		num = num*10 + digit
	}
	return num * sign
}
