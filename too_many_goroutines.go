package main

import (
	"fmt"
)

func main() {
	i := 0
	for {
		i = i + 1
		fmt.Println(i)

		go func() {
			for {
				fmt.Println(fmt.Sprintf("goroutine %d", i))
			}
		}()

	}
}
