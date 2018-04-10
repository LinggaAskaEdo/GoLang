package main

import (
	"net/http"
	"fmt"
	"html"
	"log"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, %q", html.EscapeString(request.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}