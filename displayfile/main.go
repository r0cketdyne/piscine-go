package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]
		if len(args) == 1 {
			dat, _ := os.ReadFile(args[0])
			fmt.Print(string(dat))
		} else {
			fmt.Println("Too many arguments")
		}
	} else {
		// no args, exiting?
		fmt.Println("File name missing")
	}
}
