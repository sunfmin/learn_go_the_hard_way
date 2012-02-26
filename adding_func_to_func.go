package main

import (
	"fmt"
	"reflect"
)

type Multi5 func(input string) string

func (m5 Multi5) Print(input string) {
	fmt.Println(m5(input))
}

func multi5(input string) string {
	return input + input + input + input + input
}

func main() {
	m := Multi5(multi5)
	fmt.Println(reflect.TypeOf(m))
	m.Print("hi")
}
