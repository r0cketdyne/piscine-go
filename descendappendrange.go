package piscine

func DescendAppendRange(max, min int) []int {
	if max > min {
		rangeOfMaxToMin := []int{}
		for i := max; i > min; i-- {
			rangeOfMaxToMin = append(rangeOfMaxToMin, i)
		}
		return rangeOfMaxToMin
	}
	return []int{}
}
