package main

import (
	"fmt"
)

import "github.com/go-martini/martini"
import "github.com/franela/goreq"

func main() {
	fmt.Println("hello world")

	req := goreq.Request {
		Uri: "http://www.google.com",
	}

	

	res, err := req.Do()

	if err != nil {
		fmt.Println(err.Error())
	}

	


	fmt.Printf("%s\n", "\"string\"")


	string2, _ := res.Body.ToString()
	fmt.Printf("%s\n", string2)
	fmt.Println("test: ", string2)

	app := martini.Classic()
	app.Get("/", func() string {

		
		


		return string2
	})

	for i := 0; i < 10; i++ {
		fmt.Println("loop: %d", i )
	}

	app.Run()
}
