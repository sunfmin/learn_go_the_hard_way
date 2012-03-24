package multiple_init_funcs

import (
	"fmt"
)

func init() {
	fmt.Println("first init in multiple_init_funcs/first.go")
}

func init() {
	fmt.Println("second init in multiple_init_funcs/first.go")
}
