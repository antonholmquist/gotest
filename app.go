package main

import (
	"fmt"
	"net/http"
	//"text/scanner"
	//"io"
	//"strings"
	"regexp"
)

import "github.com/go-martini/martini"
import "github.com/franela/goreq" 



func run() {
	htmlReq := goreq.Request {
		Uri: "https://partner.ikanobank.se/web/FAMILYuppdatera",
		MaxRedirects: 10,
	}

	htmlRes, err := htmlReq.Do()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		responseString, _ := htmlRes.Body.ToString()

		re := regexp.MustCompile("<link href=\"")

		var matches [][]int = re.FindAllStringIndex(responseString, -1)

		fmt.Println(matches)

		/*
		var responseStringScanner scanner.Scanner

		var responseStringReader io.Reader = strings.NewReader(responseString)


		responseStringScanner.Init(responseStringReader)

		tok := responseStringScanner.Scan()
		for tok != scanner.EOF {
			// do something with tok


			tok = responseStringScanner.Scan()
			fmt.Println("test: ", tok)
		}
		*/

		//return responseString
	}
}

func main() {

	go run()

	app := martini.Classic()

	app.Get("/", func(res http.ResponseWriter, req *http.Request, params martini.Params) string {

		

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
