package main

import (
	"github.com/astaxie/beego/validation"
	"log"
)

type User struct {
	Name string
	Age int
}

func main() {
	u := User{"man", -40}
	valid := validation.Validation{}
	valid.Required(u.Name, "name")
	valid.MaxSize(u.Name, 15, "nameMax")
	valid.Range(u.Age, 0, 140, "age")
	if valid.HasErrors() {
		// validation does not pass
		// print invalid message
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	// or use like this
	if v := valid.Max(u.Age, 140, "ageMax"); !v.Ok {
		log.Println(v.Error.Key, v.Error.Message)
	}
}
