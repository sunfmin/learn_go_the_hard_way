package attaching_func_to_a_func

import (
	"fmt"
)

type IO func(input string) string

func (io IO) Print(input string) {
	o := io(input)

	fmt.Println("IN:  ", input)
	fmt.Println("OUT: ", o)
}
