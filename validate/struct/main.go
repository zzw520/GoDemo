package main

import (
	"github.com/astaxie/beego/validation"
	"log"
)

// validation function follow with "valid" tag
// functions divide with ";"
// parameters in parentheses "()" and divide with ","
// Match function's pattern string must in "//"
type user struct {
	Id   int
	Name string `valid:"Required;Match(/^(test)?\\w*@;com$/)"`
	Age  int    `valid:"Required;Range(1, 140)"`
}

func main() {
	valid := validation.Validation{}
	// ignore empty field valid
	// see CanSkipFuncs
	// valid := validation.Validation{RequiredFirst:true}
	u := user{Name: "", Age: 0}
	b, err := valid.Valid(u)
	if err != nil {
		// handle error
	}
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}

	}
}
