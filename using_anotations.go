package main
import (
	"fmt"
	"reflect"
)

type User struct {
	Email string `column_name:"email" length:"100" null:"false" unique:"email_and_name_index"`
	Name string `column_name:"name" length:"200" null:"false" unique:"email_and_name_index"`
}

func main() {
	u := &User{}
	t1 := reflect.TypeOf(u)
	fmt.Println(t1.Kind())

	t := reflect.TypeOf(*u)
	fmt.Println(t)
	fmt.Println(t.Kind())
	fmt.Println(t.PkgPath())
	fmt.Println(t.NumMethod())
	fmt.Println(t.NumField())

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// fmt.Println(f)
		fmt.Printf("column_name => %v, length => %v, null => %v, unique => %v\n", f.Tag.Get("column_name"), f.Tag.Get("length"), f.Tag.Get("null"), f.Tag.Get("unique"))
	}
}
