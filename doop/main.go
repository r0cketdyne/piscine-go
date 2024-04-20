package main

import "os"

func atoi(s string) (int, bool) {
	num := 0
	neg := false
	for i, c := range s {
		if i == 0 && c == '-' {
			neg = true
			continue
		}
		if c < '0' || c > '9' {
			return 0, false
		}
		if num > (1<<31-1)/10 || (num == (1<<31-1)/10) && int(c-'0') > (1<<31-1)%10 {
			return 0, false
		}
		num = num*10 + int(c-'0')
	}
	if neg {
		num = -num
	}
	return num, true
}

func itoa(num int) string {
	if num == 0 {
		return "0"
	}
	neg := false
	if num < 0 {
		neg = true
		num = -num
	}

	var result []byte
	for num > 0 {
		result = append([]byte{byte(num%10) + '0'}, result...)
		num /= 10
	}
	if neg {
		result = append([]byte{'-'}, result...)
	}
	return string(result)
}

func main() {
	// Check if correct number of arguments is provided
	if len(os.Args) != 4 {
		return
	}

	// Parse arguments
	arg1, err1 := atoi(os.Args[1])
	arg3, err3 := atoi(os.Args[3])
	operator := os.Args[2]

	// Check if arguments are valid integers
	if !err1 || !err3 {
		return
	}

	// Perform operation based on the operator
	var result int
	switch operator {
	case "+":
		result = arg1 + arg3
	case "-":
		result = arg1 - arg3
	case "*":
		result = arg1 * arg3
	case "/":
		if arg3 == 0 { // Check for division by zero
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		result = arg1 / arg3
	case "%":
		if arg3 == 0 { // Check for modulo by zero
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		result = arg1 % arg3
	default:
		return // Invalid operator
	}

	// Convert result to string
	resultStr := itoa(result)

	// Write the result to standard output
	os.Stdout.WriteString(resultStr + "\n")
}
