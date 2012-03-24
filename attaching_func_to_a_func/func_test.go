package attaching_func_to_a_func

import (
	"fmt"
	"reflect"
	"strings"
)

func ExampleRun() {
	m := IO(func(i string) string {
		r := strings.Replace(i, "lower", "upper", 1)
		return strings.ToUpper(r)
	})
	fmt.Println(reflect.TypeOf(m))
	m.Print("hi, this is lower case")
	// Output:
	// attaching_func_to_a_func.IO
	// IN:   hi, this is lower case
	// OUT:  HI, THIS IS UPPER CASE
}
