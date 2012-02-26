// http://weekly.golang.org/doc/articles/error_handling.html

package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

func main() {
	f, err := os.Open("Notexists")
	fmt.Println(err)
	fmt.Println(f)
	fmt.Println(reflect.TypeOf(err))
	fmt.Println(err.Error())

	r, err1 := Num(1200).Divide(0)

	fmt.Println(r)
	fmt.Println(err1)
	fmt.Println(reflect.TypeOf(err1))
	fmt.Println(err1.Error())

	r1, err2 := Sqrt(-100)
	fmt.Println(err2)
	fmt.Println(r1)
	fmt.Println(reflect.TypeOf(err2))
}

type Num int

func (n Num) Divide(x int) (r int, err error) {
	if x == 0 {
		return 0, errors.New("can not divide 0.")
	}
	r = int(n) / x
	return
}

func Sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("root of negative number %g", n)
	}
	return 100, nil
}
