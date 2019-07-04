package main_test

import (
	"fmt"
	"path"
)

func ExampleRun() {
	fmt.Println(path.Join("https://google.com/", "hello"))
	//Output: https:/google.com/hello
}

func ExampleBase() {
	fmt.Println(path.Base("https://google.com/"))
	//Output: google.com
}
