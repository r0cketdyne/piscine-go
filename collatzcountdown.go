package piscine

func CollatzCountdown(start int) int {
	count := 0
	if start <= 0 {
		return -1
	}
	for start > 1 {
		for start%2 == 0 {
			start /= 2
			count++
		}
		if start != 1 {
			start = start*3 + 1
			count++
		}
	}
	return count
}
