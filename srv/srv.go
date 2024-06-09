package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//static fileserver
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	//sections routes
	mux.HandleFunc("GET /{$}", getRoot)
	mux.HandleFunc("GET /home/{$}", getHome)
	mux.HandleFunc("GET /browse/{$}", getBrowse)
	mux.HandleFunc("GET /search/{$}", getSearch)
	mux.HandleFunc("GET /notice/{$}", getNotice)
	mux.HandleFunc("GET /profile/{$}", getProfile)

	http.ListenAndServe(":3333", mux)
}
