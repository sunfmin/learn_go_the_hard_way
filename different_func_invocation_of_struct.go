package main

import (
	"fmt"
)

type Man struct {
	Name string
}

func (h Man) Follow(who ...string) {
	fmt.Printf("%s just followed: \n", h.Name)
	for _, v := range who {
		fmt.Println(v)
	}
}

func main() {
	h := &Man{"Chris"}

	h.Follow("Aoi", "Box")

	Man.Follow(*h, "Sora", "Box")
}
