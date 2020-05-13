package custom

import (
	"github.com/astaxie/beego/validation"
)

type user struct {
	Id   int
	Name string `valid:"Required;IsMe"`
	Age  int    `valid:"Required;Range(1, 140)"`
}

func IsMe(v *validation.Validation, obj interface{}, key string) {
	name, ok:= obj.(string)
	if !ok {
		// wrong use case?
		return
	}

	if name != "me" {
		// valid false
		v.SetError("Name", "is not me!")
	}
}

func main() {
	valid := validation.Validation{}
	if err := validation.AddCustomFunc("IsMe", IsMe); err != nil {
		// hadle error
	}
	u := user{Name: "test", Age: 40}
	b, err := valid.Valid(u)
	if err != nil {
		// handle error
	}
	if !b {
		// validation does not pass
		// blabla...
	}
}
