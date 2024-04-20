package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		for _, i := range os.Args {
			if i == "01" || i == "galaxy" || i == "galaxy 01" {
				fmt.Println("Alert!!!")
				return
			}
		}
	}
}
