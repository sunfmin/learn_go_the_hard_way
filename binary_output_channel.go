// http://rouninurashima.wordpress.com/2012/01/18/17/
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		for {

			fmt.Print(<-c)
		}
	}()

	for {
		select {
		case c <- 1:
		case c <- 0:
		}
	}

}
