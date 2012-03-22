package main

import (
	"fmt"
)

func main() {
	h := []byte("Hello, I think you are right.你是对的！")

	fmt.Printf("% x\n", h)
	fmt.Printf("%x\n", h)

	var bs []byte
	fmt.Sscanf("48656c6c6f2c2049207468696e6b20796f7520617265207269676874", "%x", &bs)
	fmt.Println(bs)

	fmt.Println(string(bs))
	fmt.Sscanf("e4bda0e698afe5afb9e79a84efbc81", "%x", &bs)
	fmt.Println(bs)
	fmt.Println(string(bs))
}
