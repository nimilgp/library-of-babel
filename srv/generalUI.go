package main

import (
	"html/template"
	"log"
	"net/http"
)

func (app *application) getEmptySectionBody(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/section-body/empty-section-body.html")
	if err != nil {
		log.Print(err)
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err)
	}
}
