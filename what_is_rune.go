package main

import (
	"bytes"
	"fmt"
)

func main() {
	nameInBytes := bytes.NewBufferString("sunäあ孙♥")

	fmt.Println(nameInBytes.Bytes())
	fmt.Println(bytes.Runes(nameInBytes.Bytes()))
	for {
		r, size, err := nameInBytes.ReadRune()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(r), r, ", size: ", size)
	}

	fmt.Printf("~m~%d~m~~j~%s\n", 10, "sunäあ孙")
	fmt.Println("len of sunäあ孙 is", len("sunäあ孙"))
}
