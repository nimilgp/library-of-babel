package main

import (
	"log"
	"net/http"
	"text/template"
)

func (app *application) getRoot(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/root.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getHome(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/sections/home.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getBrowse(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/sections/browse.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getSearch(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/sections/search.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getNotice(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/sections/notice.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getProfile(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/sections/profile.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}
