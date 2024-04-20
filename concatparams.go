package piscine

func ConcatParams(args []string) string {
	if len(args) > 0 {
		myString := ""
		for i := range args {
			myString = myString + args[i]
			if i+1 < len(args) {
				myString = myString + "\n"
			}
		}
		return myString
	}
	return ""
}
