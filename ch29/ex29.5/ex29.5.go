package main

import (
	"fmt"
	"net/http"
)

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello World")
	})

	mux.HandleFunc("/bar", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello Bar")
	})

	return mux
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
