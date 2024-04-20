package piscine

func Unmatch(a []int) int {
	unmatched := -1
	for i := range a {
		flag := true
		count := 0
		for j := range a {
			if a[i] == a[j] {
				flag = false
				count++
			}
		}
		if flag {
			unmatched = a[i]
			break
		}
		if !flag {
			if count%2 == 1 {
				unmatched = a[i]
				break
			}
		}
	}
	return unmatched
}
