package main

import (
	handler "github.com/eucalytus/embed-swagger-ui"
	"github.com/eucalytus/embed-swagger-ui/template"
	"net/http"
)

func main() {
	http.Handle("/ui/", handler.ServeWithCustomIndexHtml("/ui", template.RendCustomIndexHtml("/ui/v2.1.json")))
	http.ListenAndServe(":80", nil)
}
