package main

import(
	"fmt"
	"reflect"
)

func main() {
	a := [...]string{"Felix", "Anatole", "Juice"}
	fmt.Println(len(a))
	fmt.Println(reflect.TypeOf(a))

	b := [...]string{"Felix", "Anatole", "Juice", "Bin"}
	fmt.Println(len(b))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(&b))
	Change(&b)
	fmt.Println(b[0])
	// Change(&a)  compile error

	s := b[:]
	Change2(s)
	fmt.Println(b[0])
	fmt.Println(reflect.TypeOf(s))


	Change2(a[:])
	fmt.Println(a[0])

}

func Change(arr *[4]string) {
	arr[0] = "Felix2"
}

func Change2(s []string) {
	s[0] = "Felix Change2"
}
