package main

import (
	"fmt"
	"net/http"
	//"text/scanner"
	//"io"
	"strings"
	"regexp"
)

import "github.com/go-martini/martini"
import "github.com/franela/goreq"

func fetch(completion chan string) {
	htmlReq := goreq.Request{
		Uri:          "https://partner.ikanobank.se/web/FAMILYuppdatera",
		MaxRedirects: 10,
	}

	htmlRes, err := htmlReq.Do()

	if err != nil {
		fmt.Println(err.Error())
	} else {

		responseString, _ := htmlRes.Body.ToString()

		re := regexp.MustCompile("\"/web/")

		var matches [][]int = re.FindAllStringIndex(responseString, -1)

        modifiedResponseString := responseString

        for i := 0; i < len(matches); i++ {

            match := matches[i];
            from := match[0]
            to := match[1]

            oldString := responseString[from:to + 38]
            newString := "\"https://partner.ikanobank.se/web/"

            modifiedResponseString = strings.Replace(modifiedResponseString, oldString, newString, -1)

            fmt.Println(from)
            fmt.Println(to)
            fmt.Println(oldString)
            fmt.Println(newString)

        }

		fmt.Println(matches)

		completion <- modifiedResponseString

	}
}

func main() {

	go fetch(nil)

	app := martini.Classic()

	app.Get("/", func(res http.ResponseWriter, req *http.Request, params martini.Params) string {

		fmt.Println("Before routine")
		responseChannel := make(chan string)

		go fetch(responseChannel)

		var responseString string = <-responseChannel

		fmt.Println("After routine", responseString)

		return responseString

	})

	app.Get("/proxy/**", func(params martini.Params) string {

		req := goreq.Request{
			Uri:          "http://www.google.com",
			MaxRedirects: 10,
		}

		req.Do()

		//https://partner.ikanobank.se/web/FAMILYuppdatera

		return "Hello " + params["_1"]

	})

	for i := 0; i < 10; i++ {
		fmt.Println("loop: %d", i)
	}

	app.Run()
}
