package piscine

func IsPrime2(n int) bool {
	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(num int) int {
	if num <= 1 {
		return 2
	}
	if num == 2 {
		return 2
	}
	if num%2 == 0 {
		num++
	}
	for !IsPrime2(num) {
		num += 2 // Only check odd numbers
	}

	return num
}
