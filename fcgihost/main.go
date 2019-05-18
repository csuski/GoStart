package main

import (
	"GoStart/handlers"
	"net/http"
	"net/http/fcgi"
)

func main() {
	http.HandleFunc("/", handlers.SayHello)
	if err := fcgi.Serve(nil, nil); err != nil {
		panic(err)
	}
}
