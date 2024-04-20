package piscine

func Compact(ptr *[]string) int {
	compacted := []string{}
	count := 0
	for _, str := range *ptr {
		if str != "" {
			compacted = append(compacted, str)
			count++
		}
	}
	*ptr = compacted
	return count
}
