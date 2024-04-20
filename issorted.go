package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	flag := 0
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			if f(a[i], a[j]) > 0 && (flag == 2 || flag == 0) {
				if flag == 0 {
					flag = 1
				} else {
					return false
				}
			}
			if f(a[i], a[j]) < 0 && (flag == 1 || flag == 0) {
				if flag == 0 {
					flag = 2
				} else {
					return false
				}
			}
		}
	}
	return true
}
