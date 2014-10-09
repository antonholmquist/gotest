package main

import (
	"fmt"
)

import "github.com/go-martini/martini"

func main() {
	fmt.Println("hello world")
	app := martini.Classic()
	app.Get("/", func() string {
		return "Hello world!"
	})

	for i := 0; i < 10; i++ {
		fmt.Println("loop: %d", i )
	}

	app.Run()
}
