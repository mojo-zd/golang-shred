package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler(hello, WithBear, WithLogger))
	fmt.Println("start server ...")
	http.ListenAndServe(":8080", nil)
}

type HttpHandlerDecorator func(handler http.HandlerFunc) http.HandlerFunc

func Handler(h http.HandlerFunc, decorator ...HttpHandlerDecorator) http.HandlerFunc {
	for i := range decorator {
		d := decorator[len(decorator)-1-i]
		h = d(h)
	}
	return h
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello ...")
}

func WithBear(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("registry bear auth ...")
		handler(writer, request)
	}
}

func WithLogger(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("registry logger ...")
		handler(writer, request)
	}
}
