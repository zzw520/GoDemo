package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"
)

type ecSolution struct{
	Action string
	SolutionName string
	TemplateId string
	Language string
	Currency string
}

func main()  {
	solution:=ecSolution{
		Action:       "2",
		SolutionName: "3",
		TemplateId:   "4",
		Language:     "5",
		Currency:     "6",
	}
	data := make(map[string]interface{})
	data["solution"]=solution
	data["UserId"]="1000"
	data["UserName"]="bob"
	files, err := template.ParseFiles("template/solution.gohtml","template/user.gohtml")
	if err != nil {
		log.Fatal(err)
		return
	}
	buf := new(bytes.Buffer)
	err = files.Execute(buf, data)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Print(buf.String())

}