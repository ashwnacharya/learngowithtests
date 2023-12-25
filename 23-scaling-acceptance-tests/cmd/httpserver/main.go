package main

import (
	"net/http"
	"github.com/ashwnacharya/go-specs-greet/adapters/httpserver"
)

func main() {
	http.HandleFunc("/greet", httpserver.GreetHandler)
	http.HandleFunc("/curse", httpserver.CurseHandler)
	http.ListenAndServe(":8080", nil)
}
