package httpserver

import (
	"fmt"
	"net/http"

	go_specs_greet "github.com/ashwnacharya/go-specs-greet/domain/interactions"
)

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, go_specs_greet.Greet(name))
}

func CurseHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, go_specs_greet.Curse(name))
}