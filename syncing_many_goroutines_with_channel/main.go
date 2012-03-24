// http://rouninurashima.wordpress.com/2012/01/18/17/
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				select {
				case c <- 1:
				case c <- 2:
				}
			}
		}()

		go func() {
			for {

				select {
				case c <- 3:
				case c <- 4:
				}
			}
		}()
	}

	for {
		fmt.Print(<-c)
	}

}
