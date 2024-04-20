package piscine

func AppendRange(min, max int) []int {
	if min < max {
		rangeOfMinToMax := []int{}
		for i := min; i < max; i++ {
			rangeOfMinToMax = append(rangeOfMinToMax, i)
		}
		return rangeOfMinToMax
	}
	return nil
}
