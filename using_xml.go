package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type User struct {
	XMLName     xml.Name `xml:"user" json:"-"`
	Name        string   `xml:"name" json:"name"`
	Email       string   `xml:"email" json:"email"`
	YearOfBirth int      `xml:"year-of-birth,attr" json:"year_of_birth"`
	Sons        []string `xml:"sons>son" json:"sons"`
}

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   interface{}
}

func main() {
	u := &User{
		Name:        "Felix",
		Email:       "sunfmin@gmail.com",
		YearOfBirth: 1981,
		Sons:        []string{"Leon", "Lisa"},
	}

	us := []*User{u, u}
	uXml, _ := xml.MarshalIndent(u, "", "    ")
	fmt.Println(string(uXml))

	uJson, _ := json.MarshalIndent(u, "", "    ")
	fmt.Println(string(uJson))

	usXml, _ := xml.MarshalIndent(&Users{Users: us}, "", "    ")
	fmt.Println(string(usXml))

	usJson, _ := json.MarshalIndent(us, "", "    ")
	fmt.Println(string(usJson))
}
