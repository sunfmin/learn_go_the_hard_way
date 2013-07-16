package main

import (
	"bufio"
	"fmt"
	// "strconv"
	"os"
	// "strings"
)

func main() {
	// An artificial input source.
	f, _ := os.Open("/Users/sunfmin/Desktop/i-129.pdf")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		fmt.Println(len(data))
		l := 3
		if len(data) < l {
			l = len(data)
		}
		token = data[:l]
		advance = l
		// advance, token, err = bufio.ScanWords(data, atEOF)
		// if err == nil && token != nil {
		// _, err = strconv.ParseInt(string(token), 10, 32)
		// }
		return
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)
	// Validate the input
	for scanner.Scan() {
		fmt.Printf("\t =>%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}
