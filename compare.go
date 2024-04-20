package piscine

func Compare(a, b string) int {
	RunesA := []rune(a)
	RunesB := []rune(b)
	flag := 0
	if len(RunesA) == len(RunesB) {
		for i := 0; i < len(RunesA); i++ {
			if flag != 0 {
				break
			}
			if RunesA[i] > RunesB[i] {
				flag = 1
			} else if RunesA[i] < RunesB[i] {
				flag = -1
			}
		}
	} else if len(RunesA) < len(RunesB) {
		for i := 0; i < len(RunesA); i++ {
			if flag != 0 {
				break
			}
			if RunesA[i] > RunesB[i] {
				flag = 1
			} else if RunesA[i] < RunesB[i] {
				flag = -1
			}
		}
		if flag == 0 {
			flag = -1
		}
	} else {
		for i := 0; i < len(RunesB); i++ {
			if flag != 0 {
				break
			}
			if RunesA[i] > RunesB[i] {
				flag = 1
			} else if RunesA[i] < RunesB[i] {
				flag = -1
			}
		}
		if flag == 0 {
			flag = 1
		}
	}
	return flag
}
