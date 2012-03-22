package main

import (
	"fmt"
)

type Values map[string][]string

type MyInt int64

func (v Values) Get(key string) string {
	m := map[string][]string(v)
	return m[key][0]
}

func (i MyInt) Tripple() int64 {
	return int64(i) * 3
}

type MyInterface interface {
	Tripple() int64
}

type Object interface {
}

func main() {

	a := map[string][]string{
		"key1": []string{"1", "2"},
		"key2": []string{"3", "4"},
	}

	b := Values(a)

	fmt.Println(b.Get("key1"))
	fmt.Println(b.Get("key2"))

	fmt.Println(MyInt(100).Tripple())

	var obj Object

	obj = b
	objB, yes := obj.(Values)
	fmt.Println(yes)
	fmt.Println(objB.Get("key1"))

	objInt, yes1 := obj.(MyInt)
	fmt.Println(yes1)
	fmt.Println(objInt)

	var m MyInterface
	m = MyInt(9)
	fmt.Println(m.Tripple())

	var myint1 MyInt
	myint1 = 102

	var myif MyInterface
	myif = myint1
	fmt.Println(myif.Tripple())

	// fmt.Println(obj.(type))

	switch i := obj.(type) {
	case MyInt:
		fmt.Println("x is nil", i)
	case Values:
		fmt.Println("x is Values", i) // i is an int
	default:
		fmt.Println("don't know the type", i)
	}

}
