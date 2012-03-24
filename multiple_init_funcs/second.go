package multiple_init_funcs

import (
	"fmt"
)

func init() {
	fmt.Println("first init in multiple_init_funcs/second.go")
}

func init() {
	fmt.Println("second init in multiple_init_funcs/second.go")
}

func Hello() {
	fmt.Println("Hello")
}
