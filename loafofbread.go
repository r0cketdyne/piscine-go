package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	after := ""
	i := 0
	c := 0
	for i < len(str) {
		if c%6 != 5 {
			if str[i] != ' ' {
				after = after + string(str[i])
				c++
			}
		} else {
			after = after + " "
			c++
		}
		i++
	}
	i = len(after) - 1
	for i >= 0 {
		if after[i] == ' ' {
			after = after[0:i]
			i--
		}
		if after[i] != ' ' {
			break
		}
		i--
	}
	if len(after) == 0 {
		return "\n"
	}
	if len(after) < 5 {
		return "Invalid Output\n"
	}
	after = after + "\n"
	return after
}
