package main

import (
	"fmt"
	"net/http"
	//"text/scanner"
	//"io"
	"encoding/base64"
	//"regexp"
	"strings"
	"io/ioutil"
)

import "github.com/go-martini/martini"
import "github.com/franela/goreq"

func fetch() string {
	htmlReq := goreq.Request{
		Uri:          "https://partner.ikanobank.se/web/engines/page.aspx?structid=3870&AspxAutoDetectCookieSupport=1",
		MaxRedirects: 0,
	}

	htmlReq.AddHeader("Cookie", "AspxAutoDetectCookieSupport=1; path=/")

	htmlRes, err := htmlReq.Do()

	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	} else {

		responseString, _ := htmlRes.Body.ToString()

		modifiedResponseString := responseString

		// Make relative urls absolute
		modifiedResponseString = strings.Replace(modifiedResponseString, "\"/web/", "\"https://partner.ikanobank.se/web/", -1)

		// Special one
		modifiedResponseString = strings.Replace(modifiedResponseString, "BotDetectCaptcha.ashx?get=layoutStyleSheet", "https://partner.ikanobank.se/web/BotDetectCaptcha.ashx?get=layoutStyleSheet", 1)
		
		// Remove some standard strings
		modifiedResponseString = strings.Replace(modifiedResponseString, "<!DOCTYPE html>", "", 1)
		modifiedResponseString = strings.Replace(modifiedResponseString, "<html>", "", 1)
		modifiedResponseString = strings.Replace(modifiedResponseString, "<head>", "", 1)
		modifiedResponseString = strings.Replace(modifiedResponseString, "</head>", "", 1)
		modifiedResponseString = strings.Replace(modifiedResponseString, "<body>", "", 1)
		modifiedResponseString = strings.Replace(modifiedResponseString, "</body>", "", 1)
		modifiedResponseString = strings.Replace(modifiedResponseString, "</html>", "", 1)
		
		


		return modifiedResponseString

	}
}

func main() {

	fetch()

	app := martini.Classic()

	app.Get("/", func(res http.ResponseWriter, req *http.Request, params martini.Params) string {

		return "<!DOCTYPE html><html><body><div><script type=\"text/javascript\" src=\"http://ikano-ikea-family.herokuapp.com/jsonp.js\"></script></div></body></html>"

	})

	app.Get("/jsonp.js", func(res http.ResponseWriter, req *http.Request, params martini.Params) string {

		var responseString = fetch()

		responseStringBase64 := base64.StdEncoding.EncodeToString([]byte(responseString))


		bytes, err := ioutil.ReadFile("./src/github.com/antonholmquist/gotest/script.js")

		scriptString := string(bytes)

		if (err != nil) {
			fmt.Println("error: " + err.Error())
		}

		scriptString = strings.Replace(scriptString, "<base_64_content>", responseStringBase64, 1)

		return scriptString

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
