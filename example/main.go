package main

import (
	handler "github.com/eucalytus/embed-swagger-ui"
	"net/http"
)

func main() {
	http.Handle("/ui/", http.StripPrefix("/ui", handler.SwaggerUIHandler))
	http.ListenAndServe(":80", nil)
}
