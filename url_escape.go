package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Add("cid", "IE{7d1b54c0-c181-11e0-b928-806d6172696f}0107593355446140")
	fmt.Println(v.Encode())
}
