package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<b>hellooooo!!</b>")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/hello", hello)

	http.ListenAndServe(":3333", mux)
}
