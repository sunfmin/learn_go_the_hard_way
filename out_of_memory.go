package main

import (
	"fmt"
	"reflect"
	/*	"unsafe"*/
)

func main() {
	m := "Lv"

	/*	var mAddress *string*/

	m1 := &m
	mAddress := &m1
	for {
		mAddress1 := &mAddress
		fmt.Println("address's address type: ", reflect.TypeOf(mAddress1))
	}

}
