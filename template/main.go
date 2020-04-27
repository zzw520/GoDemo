package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type ecSolution struct{
	UserId string
	Action string
	SolutionName string
	TemplateId string
	Language string
	currency string
}

func main()  {
	solution:=ecSolution{
		UserId:       "1",
		Action:       "2",
		SolutionName: "3",
		TemplateId:   "4",
		Language:     "5",
		currency:     "6",
	}
	fmt.Printf("aaa:%s",filepath.Base("./"))
	s, _ := filepath.Abs("./")
	fmt.Printf("bbb:%s", s)
	file, err := filepath.Abs("./test.gohtml")
	if err!=nil {
		log.Fatal(err)
		return
	}
	files, err := template.ParseFiles(file)
	if err != nil {
		log.Fatal(err)
		return
	}
	files.Execute(os.Stdout,solution)

}