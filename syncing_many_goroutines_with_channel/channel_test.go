// http://rouninurashima.wordpress.com/2012/01/18/17/
package syncing_many_goroutines_with_channel

import (
	"fmt"
)

func ExampleRun() {
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
