package piscine

func isPerfectSquare(num int) bool {
	sqrt := 0
	for i := 1; i <= num/2+1; i++ {
		if i*i == num {
			sqrt = i
			break
		}
	}
	return sqrt*sqrt == num
}

func Sqrt(num int) int {
	if isPerfectSquare(num) {
		sqrt := 0
		for i := 1; i <= num/2+1; i++ {
			if i*i == num {
				sqrt = i
				break
			}
		}
		return sqrt
	}
	return 0
}
