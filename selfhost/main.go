package main

import (
	"GoStart/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.SayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
