package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Dog  Dog
}

type Dog struct {
	Name  string
	Breed string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name: "Petros Trak",
		Dog: Dog{
			Name:  "Jack",
			Breed: "Golden Retriever",
		},
	}

	// data.Name = "<script>alert('hi')</script>"
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
