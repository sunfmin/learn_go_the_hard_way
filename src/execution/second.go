package execution

import (
	"fmt"
)

func init() {
	fmt.Println("first init in execution/second.go")
}

func init() {
	fmt.Println("second init in execution/second.go")
}

func Hello() {
	fmt.Println("Hello")
}
