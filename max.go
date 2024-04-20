package piscine

func Max(a []int) int {
	max := 0
	for _, i := range a {
		if max < i {
			max = i
		}
	}
	return max
}
