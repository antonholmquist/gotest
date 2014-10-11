package main

import (
	"fmt"
	"net/http"
	"text/scanner"
	"io"
	"strings"
)

import "github.com/go-martini/martini"
import "github.com/franela/goreq" 

func main() {

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

			var responseStringScanner scanner.Scanner

			var responseStringReader io.Reader = strings.NewReader(responseString)

			responseStringScanner.Init(responseStringReader)

			tok := responseStringScanner.Scan()
			for tok != scanner.EOF {
				// do something with tok


				tok = responseStringScanner.Scan()
				fmt.Println("test: ", tok)
			}

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
