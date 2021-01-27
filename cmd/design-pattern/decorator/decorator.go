package decorator

import (
	"fmt"
	"net/http"
)

type HttpHandlerDecorator func(handler http.HandlerFunc) http.HandlerFunc

func Handler(h http.HandlerFunc, decorator ...HttpHandlerDecorator) http.HandlerFunc {
	for i := range decorator {
		d := decorator[len(decorator)-1-i]
		h = d(h)
	}
	return h
}

func Route() {
	http.HandleFunc("/v1", Handler(hello, WithBear))
}

func hello(w http.ResponseWriter, r *http.Request) {

}

func WithBear(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("with bear")
		handler(writer, request)
	}
}
