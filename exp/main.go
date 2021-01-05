package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name: "Petros Trak",
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
