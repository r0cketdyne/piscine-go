package piscine

func Rot14(s string) string {
	bytes := []byte(s)
	for i, char := range s {
		if (char >= 'a' && char <= 'l') || (char >= 'A' && char <= 'L') {
			bytes[i] = byte(char + 14)
		}
		if (char >= 'm' && char <= 'z') || (char >= 'M' && char <= 'Z') {
			bytes[i] = byte(char - 12)
		}
	}
	return string(bytes)
}
