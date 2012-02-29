package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
}

func (u *User) Hello() {
	fmt.Println("Hello", u.Name)
}

func main() {
	u := &User{"Felix"}
	v := reflect.ValueOf(u)
	fmt.Println(v)
	fmt.Println(v.MethodByName("Hello").Call([]reflect.Value{}))
}
