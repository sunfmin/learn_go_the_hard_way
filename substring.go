package main

import (
	"fmt"
)

func main() {
	s := "helloworld"
	fmt.Println(s[0:2])
	fmt.Println(s[2:])
	fmt.Println(s[:6])
}
