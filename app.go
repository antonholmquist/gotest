package main

import (
	"fmt"
	"net/http"
)

import "github.com/go-martini/martini"
import "github.com/franela/goreq"

func main() {

	//originalURL = "http://www.google."

	fmt.Println("hello world")

	req := goreq.Request {
		Uri: "http://www.google.com",
		MaxRedirects: 10,
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

	app.Get("/", func(res http.ResponseWriter, req *http.Request, params martini.Params) string {

		htmlReq := goreq.Request {
			Uri: "https://partner.ikanobank.se/web/FAMILYuppdatera",
			MaxRedirects: 10,
		}

		htmlRes, err := htmlReq.Do()

		if err != nil {
			fmt.Println(err.Error())
		} else {
			responseString, _ := htmlRes.Body.ToString()

			return responseString
		}

		//https://partner.ikanobank.se/web/FAMILYuppdatera

		return "Hello " + params["_1"]

	})

	app.Get("/proxy/**", func(params martini.Params) string {

		req := goreq.Request {
			Uri: "http://www.google.com",
			MaxRedirects: 10,
		}

		req.Do()

		//https://partner.ikanobank.se/web/FAMILYuppdatera

		return "Hello " + params["_1"]

	})

	for i := 0; i < 10; i++ {
		fmt.Println("loop: %d", i )
	}

	app.Run()
}
