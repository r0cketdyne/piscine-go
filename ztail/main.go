package main

import (
	"fmt"
	"os"
)

func parseOffset(numBytes string, fileSize int64) (int64, error) {
	var offset int64
	var err error

	fmt.Sscanf(numBytes, "%d", &offset)

	return fileSize - offset, err
}

func main() {
	myError := false
	if len(os.Args) > 3 {
		args := os.Args[1:]
		option := args[0]
		if option != "-c" {
			myError = true
			fmt.Println("Error: Invalid arguments. style: go run -c number filename.")
			os.Exit(1)
		}
		args = args[1:]
		numBytes := args[0]
		args = args[1:]
		for _, i := range numBytes {
			if i < '0' || i > '9' {
				myError = true
				fmt.Println("Error: Invalid arguments. style: go run -c number filename.")
				os.Exit(1)
			}
		}
		for i, filename := range args {
			data, err := os.Open(filename)
			if err != nil {
				myError = true
				fmt.Printf("%v\n", err)
				if i < len(args)-1 {
					_, err = os.Open(args[i+1])
				}
				if i != len(args)-1 && err == nil {
					fmt.Println()
				}
				continue
			} else {
				fmt.Printf("==> %s <==\n", filename)
			}
			defer data.Close()

			stat, err := data.Stat()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error getting file information for %s: %v\n", filename, err)
				if i < len(args)-1 {
					_, err = os.Open(args[i+1])
				}
				if i != len(args)-1 && err == nil {
					fmt.Println()
				}
				continue
			}

			offset, err := parseOffset(numBytes, stat.Size())
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing offset for file %s: %v\n", filename, err)
				if i < len(args)-1 {
					_, err = os.Open(args[i+1])
				}
				if i != len(args)-1 && err == nil {
					fmt.Println()
				}
				continue
			}

			_, err = data.Seek(offset, 0)
			buffer := make([]byte, 4096)
			for {
				n, err := data.Read(buffer)
				if err != nil && err.Error() != "EOF" {
					fmt.Fprintf(os.Stderr, "Error reading from file %s: %v\n", filename, err)
					break
				}
				if n == 0 {
					break
				}
				os.Stdout.Write(buffer[:n])
			}
			if i < len(args)-1 {
				_, err = os.Open(args[i+1])
			}
			if i != len(args)-1 && err == nil {
				fmt.Println()
			}
		}
	}
	if myError {
		os.Exit(1)
	}
}
