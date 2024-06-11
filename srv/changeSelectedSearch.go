package main

import (
	"html/template"
	"log"
	"net/http"
)

func (app *application) getSearchTypeTitle(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/section-intro/search-sel-title.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getSearchTypeAuthor(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/section-intro/search-sel-author.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getSearchTypeISBN(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/section-intro/search-sel-isbn.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}
