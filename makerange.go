package piscine

func MakeRange(min, max int) []int {
	if min < max {
		rangeOfMinToMax := make([]int, max-min)
		for i := min; i < max; i++ {
			rangeOfMinToMax[i-min] = i
		}
		return rangeOfMinToMax
	}
	return nil
}
