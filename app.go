package main

import (
	"fmt"
)

import "github.com/go-martini/martini"

func main() {
	fmt.Println("hello world")
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
