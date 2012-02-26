package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	h := "Hello"
	l := "Hello World"
	n := "12345678901234567890123456789012345678901234567890123456789012345678901234567890"
	hAddress := &h
	hAddress2 := &hAddress
	var addressLiteral uintptr
	addressLiteral = 0x4212c0d0

	hAddressLiteral := unsafe.Pointer(addressLiteral)

	fmt.Println("type of pointer: ", reflect.TypeOf(hAddress))
	fmt.Println("size of pointer: ", unsafe.Sizeof(hAddress))
	fmt.Println("size of pointer of 'Hello World': ", unsafe.Sizeof(&l))
	fmt.Println("size of pointer of '12345678901234567890123456789012345678901234567890123456789012345678901234567890': ", unsafe.Sizeof(&n))

	fmt.Println("h address: ", hAddress)
	fmt.Println("h value: ", *hAddress)
	fmt.Println("h address, value: ", *&h)
	fmt.Println("h address variable's address: ", &hAddress)
	fmt.Println("h address variable's address: ", &hAddress2)
	fmt.Println("equal: ", *hAddress == h)
	fmt.Println("unsafe.Pointer: ", hAddressLiteral)
	/*	fmt.Println("equal: ", *hAddressLiteral)*/

	m := "Lv"
	m1 := &m
	m2 := &m1
	m3 := &m2
	m4 := &m3
	m5 := &m4
	fmt.Println("addresses: ", m1, m2, m3, m4, m5)

}
