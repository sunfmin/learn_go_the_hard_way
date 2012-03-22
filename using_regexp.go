package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`\/l\/(\w+)`)

	str := "http://google.com/l/dd9c69bf0ad9d58ac343541f7037deb8b5cc38a7"
	rs1 := re.FindAllStringSubmatch(str, -1)
	fmt.Println(rs1)

	rs2 := re.FindAllString(str, -1)
	fmt.Println(rs2)

	rs3 := re.FindStringSubmatch(str)
	fmt.Println(rs3)
}
