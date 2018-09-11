package parse

import (
	"gitlab.com/golang-commonmark/markdown"
	"io/ioutil"
	"net/http"
)

func Parse(response http.ResponseWriter, request *http.Request) {
	body, _ := ioutil.ReadAll(request.Body)

	md := markdown.New(markdown.HTML(true))
	parse := md.Parse(body)
	md.RenderTokens(response, parse)
}