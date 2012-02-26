package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []string{"1", "2"}
	b := []string{"3", "4", "5"}
	c := append(a, b...)
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(c)
}
