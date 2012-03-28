package main

import (
	"fmt"
	"github.com/paulbellamy/mango"
)

type Env mango.Env

func (e Env) LoggedInAccount() interface{} {
	return map[string]interface{}(e)["accountid"]
}

func main() {
	e := Env(mango.Env{})

	fmt.Println(e.LoggedInAccount())
}
