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

func fetch() string {
	htmlReq := goreq.Request{
		Uri:          "https://partner.ikanobank.se/web/FAMILYuppdatera",
		MaxRedirects: 10,
	}

	htmlRes, err := htmlReq.Do()

	if err != nil {
		fmt.Println(err.Error())
    return err.Error()
	} else {

		responseString, _ := htmlRes.Body.ToString()

		re := regexp.MustCompile("<link href=\"/web/")

		var matches [][]int = re.FindAllStringIndex(responseString, -1)

        modifiedResponseString := responseString

        /*
        Modify relative url so that links will still be correct
        From: <link href="/web/(X(1)S(j0kdjk45m1cpmq45fdxok4n0))/web/site_files/templates/template.css"
        To: <link href="https://partner.ikanobank.se/web/site_files/templates/template.css"

        */

        for i := 0; i < len(matches); i++ {

            match := matches[i];
            from := match[0]
            to := match[1]

            oldString := responseString[from:to + 38]
            newString := "<link href=\"https://partner.ikanobank.se/web/"

            modifiedResponseString = strings.Replace(modifiedResponseString, oldString, newString, -1)
        }

		return modifiedResponseString

	}
}

func main() {

	fetch()

	app := martini.Classic()

	app.Get("/", func(res http.ResponseWriter, req *http.Request, params martini.Params) string {

		var responseString = fetch()
		return responseString

	})

  app.Get("/jsonp.js", func(res http.ResponseWriter, req *http.Request, params martini.Params) string {

    var responseString = fetch()

    jsonp := "var ikanoIkeaFamilyCallback = function(\"" + responseString + "\");"

    return jsonp

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
