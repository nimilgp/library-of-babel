package main

import (
	"html/template"
	"log"
	"net/http"
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
	ts, err := template.ParseFiles("./ui/html/main-sections/home.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getBrowse(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/main-sections/browse.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getSearch(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/main-sections/search.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getNotice(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/main-sections/notice.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getSignIn(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/main-sections/sign-in.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}

func (app *application) getSignUp(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/main-sections/sign-up.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}
