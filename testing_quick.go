package main

import (
	"fmt"
	"testing/quick"
)

/*func TestOddMultipleOfThree(t *testing.T) {
	f := func(x int) bool {
		y := OddMultipleOfThree(x)
		return y%2 == 1 && y%3 == 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}*/

type Person struct {
	Name string
	Year int
	/*	Parent *Person*/
}

func main() {
	intf := func(x int) bool {
		fmt.Println(x)
		return true
	}

	quick.Check(intf, nil)

	personf := func(p *Person) bool {
		fmt.Printf("%+v\n", p)
		return true
	}
	quick.Check(personf, nil)
}
