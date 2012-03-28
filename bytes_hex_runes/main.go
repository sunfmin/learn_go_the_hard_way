package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
)

func main() {
	h := []byte("Hello, I think you are right.你是对的！")

	fmt.Println(`Say you have a string:`, "\n\t", string(h))
	fmt.Println(`Whose byte array in decimal are:`, "\n\t", h)
	fmt.Println(`So in hex: `, "\n\t", fmt.Sprintf("% x", h))
	hex := fmt.Sprintf("%x", h)
	fmt.Println(`So you see this form a lot:`, "\n\t", hex)

	buffer := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.URLEncoding, buffer)
	encoder.Write(h)
	encoder.Close()

	fmt.Println(`And base64 of the same data:`, "\n\t", buffer.String())

	var bs []byte
	fmt.Sscanf(hex, "%x", &bs)

	fmt.Println(`And you can parse`, "\n\t", hex, "\n", "into bytes", "\n\t", bs)

	fmt.Println(`Whose runes are:`, "\n\t", bytes.Runes(bs))
	fmt.Println(`And each rune matching one charactor not matter its chinese or english:`, "\n\t", string(bs))

}
