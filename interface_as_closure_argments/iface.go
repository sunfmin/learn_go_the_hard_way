package main

import (
	"fmt"
	"reflect"
)

type Input interface {
	Values() map[string]string
}

type in struct {
	values map[string]string
}

func (i *in) Values() map[string]string {
	return i.values
}

type Runable func(i Input)

func WithUser(Runable) Runable {
	return func(i Input) {
		m := i.Values()
		m["user"] = "felix"
		fmt.Println("pointer inside wrapper: ", reflect.ValueOf(i).Pointer())
	}
}

func Run1(i Input) {
	fmt.Println("pointer inside func: ", reflect.ValueOf(i).Pointer())
}

func main() {
	i := &in{map[string]string{}}
	fmt.Println("pointer outside: ", reflect.ValueOf(i).Pointer())
	WithUser(Run1)(i)
	fmt.Println(i.Values())
}
